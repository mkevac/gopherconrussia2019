package biggerbatch

import (
	"math/rand"
	"testing"
	"time"
)

const restaurants = 65536
const bitmapLength = restaurants / (8 * 8) // 8192 bytes (1024 elements)

func initData() ([]uint64, []uint64, []uint64, []uint64, []uint64, []uint64) {
	var (
		nearMetro      = make([]uint64, bitmapLength)
		privateParking = make([]uint64, bitmapLength)
		terrace        = make([]uint64, bitmapLength)
		reservations   = make([]uint64, bitmapLength)
		veganFriendly  = make([]uint64, bitmapLength)
		expensive      = make([]uint64, bitmapLength)
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

func TestBiggerBatchBitmapIndex(t *testing.T) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	andnot(terrace, expensive, resBitmap)
	and(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func BenchmarkBiggerBatchBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnot(terrace, expensive, resBitmap)
		and(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkBiggerBatchBitmapIndexInlined(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotInlined(terrace, expensive, resBitmap)
		andInlined(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkBiggerBatchBitmapIndexNoBoundsCheck(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotNoBoundsCheck(terrace, expensive, resBitmap)
		andNoBoundsCheck(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkBiggerBatchBitmapIndexInlinedAndNoBoundsCheck(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotInlinedAndNoBoundsCheck(terrace, expensive, resBitmap)
		andInlinedAndNoBoundsCheck(reservations, resBitmap, resBitmap)
	}
}
