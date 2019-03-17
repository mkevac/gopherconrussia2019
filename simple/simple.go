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

func andInlined(a []byte, b []byte, res []byte) {
	i := 0
	l := len(a)

loop:
	res[i] = a[i] & b[i]
	i++
	if i != l {
		goto loop
	}
}

func or(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] | b[i]
	}
}

func orInlined(a []byte, b []byte, res []byte) {
	i := 0
	l := len(a)

loop:
	res[i] = a[i] | b[i]
	i++
	if i != l {
		goto loop
	}
}

func not(a []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = ^a[i]
	}
}

func notInlined(a []byte, res []byte) {
	i := 0
	l := len(a)

loop:
	res[i] = ^a[i]
	i++
	if i != l {
		goto loop
	}
}
