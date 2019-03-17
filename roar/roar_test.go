package roar

import (
	"math/rand"
	"testing"
	"time"

	"github.com/RoaringBitmap/gocroaring"
	"github.com/RoaringBitmap/roaring"
)

const restaurants = 65536

func fill(b *roaring.Bitmap, probability float32) {
	rand.Seed(time.Now().UnixNano())

	for i := uint32(0); i < restaurants; i++ {
		if rand.Float32() < probability {
			b.Add(i)
		}
	}
}

func fill2(b *gocroaring.Bitmap, probability float32) {
	rand.Seed(time.Now().UnixNano())

	for i := uint32(0); i < restaurants; i++ {
		if rand.Float32() < probability {
			b.Add(i)
		}
	}
}

func TestRoaringBitmapIndex(t *testing.T) {
	nearMetro := roaring.New()
	privateParking := roaring.New()
	terrace := roaring.New()
	reservations := roaring.New()
	veganFriendly := roaring.New()
	expensive := roaring.New()

	fill(nearMetro, 0.1)
	fill(privateParking, 0.01)
	fill(terrace, 0.05)
	fill(reservations, 0.95)
	fill(veganFriendly, 0.2)
	fill(expensive, 0.1)

	res := roaring.And(reservations, roaring.AndNot(terrace, expensive))
	t.Log(res.GetCardinality())
}

func TestCRoaringBitmapIndex(t *testing.T) {
	nearMetro := gocroaring.New()
	privateParking := gocroaring.New()
	terrace := gocroaring.New()
	reservations := gocroaring.New()
	veganFriendly := gocroaring.New()
	expensive := gocroaring.New()

	fill2(nearMetro, 0.1)
	fill2(privateParking, 0.01)
	fill2(terrace, 0.05)
	fill2(reservations, 0.95)
	fill2(veganFriendly, 0.2)
	fill2(expensive, 0.1)

	res := gocroaring.And(reservations, gocroaring.AndNot(terrace, expensive))
	t.Log(res.Cardinality())
}

func BenchmarkRoaringBitmapIndex(b *testing.B) {
	nearMetro := roaring.New()
	privateParking := roaring.New()
	terrace := roaring.New()
	reservations := roaring.New()
	veganFriendly := roaring.New()
	expensive := roaring.New()

	fill(nearMetro, 0.1)
	fill(privateParking, 0.01)
	fill(terrace, 0.05)
	fill(reservations, 0.95)
	fill(veganFriendly, 0.2)
	fill(expensive, 0.1)

	var res *roaring.Bitmap

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = roaring.And(reservations, roaring.AndNot(terrace, expensive))
	}

	_ = res
}

func BenchmarkCRoaringBitmapIndex(b *testing.B) {
	nearMetro := gocroaring.New()
	privateParking := gocroaring.New()
	terrace := gocroaring.New()
	reservations := gocroaring.New()
	veganFriendly := gocroaring.New()
	expensive := gocroaring.New()

	fill2(nearMetro, 0.1)
	fill2(privateParking, 0.01)
	fill2(terrace, 0.05)
	fill2(reservations, 0.95)
	fill2(veganFriendly, 0.2)
	fill2(expensive, 0.1)

	var res *gocroaring.Bitmap

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = gocroaring.And(reservations, gocroaring.AndNot(terrace, expensive))
	}

	_ = res
}
