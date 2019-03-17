package simple

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

func TestSimpleBitmapIndex(t *testing.T) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	andnot(terrace, expensive, resBitmap)
	and(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func BenchmarkSimpleBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnot(terrace, expensive, resBitmap)
		and(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkSimpleBitmapIndexInlined(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotInlined(terrace, expensive, resBitmap)
		andInlined(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkSimpleBitmapIndexInlinedAndNoBoundsCheck(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		andnotInlinedAndNoBoundsCheck(terrace, expensive, resBitmap)
		andInlinedAndNoBoundsCheck(reservations, resBitmap, resBitmap)
	}
}
