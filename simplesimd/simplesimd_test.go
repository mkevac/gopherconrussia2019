package simplesimd

import (
	"math/rand"
	"testing"
	"time"
)

const restaurants = 65536
const bitmapLength = restaurants / 8 // 8192 bytes

func initData() ([]byte, []byte, []byte, []byte, []byte, []byte) {
	var (
		nearMetro      = make([]byte, bitmapLength)
		privateParking = make([]byte, bitmapLength)
		terrace        = make([]byte, bitmapLength)
		reservations   = make([]byte, bitmapLength)
		veganFriendly  = make([]byte, bitmapLength)
		expensive      = make([]byte, bitmapLength)
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fill(r, nearMetro, 0.1)
	fill(r, privateParking, 0.01)
	fill(r, terrace, 0.05)
	fill(r, reservations, 0.95)
	fill(r, veganFriendly, 0.2)
	fill(r, expensive, 0.1)

	return nearMetro, privateParking, terrace, reservations, veganFriendly, expensive
}

func TestSIMDSimpleAnd(t *testing.T) {
	var (
		bitmapA = make([]byte, bitmapLength)
		bitmapB = make([]byte, bitmapLength)
	)

	for i := 0; i < len(bitmapA); i++ {
		bitmapA[i] = 3 << 1
		bitmapB[i] = 3
	}

	t.Run("scalar", func(t *testing.T) {

		var bitmapRes = make([]byte, bitmapLength)

		andScalar(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 2 {
				t.Fatalf("byte %d of result is %d (expected 2)", i, bitmapRes[i])
			}
		}
	})

	t.Run("scalar faster", func(t *testing.T) {

		var bitmapRes = make([]byte, bitmapLength)

		andScalarFaster(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 2 {
				t.Fatalf("byte %d of result is %d (expected 2)", i, bitmapRes[i])
			}
		}
	})

	t.Run("simd", func(t *testing.T) {

		var bitmapRes = make([]byte, bitmapLength)

		andSIMD(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 2 {
				t.Fatalf("byte %d of result is %d (expected 2)", i, bitmapRes[i])
			}
		}
	})
}

func TestSIMDSimpleOr(t *testing.T) {
	var (
		bitmapA = make([]byte, bitmapLength)
		bitmapB = make([]byte, bitmapLength)
	)

	for i := 0; i < len(bitmapA); i++ {
		bitmapA[i] = 3 << 1
		bitmapB[i] = 3
	}

	t.Run("scalar", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		orScalar(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 7 {
				t.Fatalf("byte %d of result is %d (expected 7)", i, bitmapRes[i])
			}
		}
	})

	t.Run("scalar faster", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		orScalarFaster(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 7 {
				t.Fatalf("byte %d of result is %d (expected 7)", i, bitmapRes[i])
			}
		}
	})

	t.Run("simd", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		orSIMD(bitmapA, bitmapB, bitmapRes)

		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != 7 {
				t.Fatalf("byte %d of result is %d (expected 7)", i, bitmapRes[i])
			}
		}
	})
}

func TestSIMDSimpleAndNot(t *testing.T) {
	var (
		bitmapA = make([]byte, bitmapLength)
		bitmapB = make([]byte, bitmapLength)
	)

	for i := 0; i < len(bitmapA); i++ {
		bitmapA[i] = 255
		bitmapB[i] = ^(byte(1) << 5)
	}

	t.Run("scalar", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		andnotScalar(bitmapA, bitmapB, bitmapRes)

		var expected byte = 1 << 5
		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != byte(expected) {
				t.Fatalf("byte %d of result is %d (expected %d)", i, bitmapRes[i], expected)
			}
		}
	})

	t.Run("scalar faster", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		andnotScalarFaster(bitmapA, bitmapB, bitmapRes)

		var expected byte = 1 << 5
		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != byte(expected) {
				t.Fatalf("byte %d of result is %d (expected %d)", i, bitmapRes[i], expected)
			}
		}
	})

	t.Run("simd", func(t *testing.T) {
		var bitmapRes = make([]byte, bitmapLength)
		andnotSIMD(bitmapA, bitmapB, bitmapRes)

		var expected byte = 1 << 5
		for i := 0; i < len(bitmapRes); i++ {
			if bitmapRes[i] != byte(expected) {
				t.Fatalf("byte %d of result is %d (expected %d)", i, bitmapRes[i], expected)
			}
		}
	})
}

func TestSIMDBitmapIndex(t *testing.T) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	andnotSIMD(terrace, expensive, resBitmap)
	andSIMD(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func TestScalarBitmapIndex(t *testing.T) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	andnotScalar(terrace, expensive, resBitmap)
	andScalar(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func TestScalar2BitmapIndex(t *testing.T) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	andnotScalarFaster(terrace, expensive, resBitmap)
	andScalarFaster(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func BenchmarkSimpleSIMDBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotSIMD(terrace, expensive, resBitmap)
		andSIMD(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkSimpleScalarFasterBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotScalarFaster(terrace, expensive, resBitmap)
		andScalarFaster(reservations, resBitmap, resBitmap)
	}
}


func BenchmarkSimpleScalarBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotScalar(terrace, expensive, resBitmap)
		andScalar(reservations, resBitmap, resBitmap)
	}
}
