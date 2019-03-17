package simple

const restaurants = 65536
const bitmapLength = restaurants / 8 // 8192 bytes

var (
	nearMetro      = make([]byte, bitmapLength)
	privateParking = make([]byte, bitmapLength)
	terrace        = make([]byte, bitmapLength)
	reservations   = make([]byte, bitmapLength)
	veganFriendly  = make([]byte, bitmapLength)
	expensive      = make([]byte, bitmapLength)
)

func and(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & b[i]
	}
}

func or(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] | b[i]
	}
}

func not(a []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = ^a[i]
	}
}
