package simple

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpleBitmapIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	fill(nearMetro, 0.1)
	fill(privateParking, 0.01)
	fill(terrace, 0.05)
	fill(reservations, 0.95)
	fill(veganFriendly, 0.2)
	fill(expensive, 0.1)

	resBitmap := make([]byte, bitmapLength)

	not(expensive, resBitmap)
	and(terrace, resBitmap, resBitmap)
	and(reservations, resBitmap, resBitmap)

	resRestaurants := indexes(resBitmap)

	t.Log(len(resRestaurants))
}

func BenchmarkSimpleBitmapIndex(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	fill(nearMetro, 0.1)
	fill(privateParking, 0.01)
	fill(terrace, 0.05)
	fill(reservations, 0.95)
	fill(veganFriendly, 0.2)
	fill(expensive, 0.1)

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		not(expensive, resBitmap)
		and(terrace, resBitmap, resBitmap)
		and(reservations, resBitmap, resBitmap)
	}
}

func BenchmarkSimpleBitmapIndexInlined(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	fill(nearMetro, 0.1)
	fill(privateParking, 0.01)
	fill(terrace, 0.05)
	fill(reservations, 0.95)
	fill(veganFriendly, 0.2)
	fill(expensive, 0.1)

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		notInlined(expensive, resBitmap)
		andInlined(terrace, resBitmap, resBitmap)
		andInlined(reservations, resBitmap, resBitmap)
	}
}
