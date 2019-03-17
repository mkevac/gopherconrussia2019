package simple

import "math/rand"

func fill(r *rand.Rand, b []byte, probability float32) {
	for i := 0; i < len(b); i++ {
		for j := uint(0); j < 8; j++ {
			if r.Float32() < probability {
				b[i] |= 1 << j
			}
		}
	}
}

func indexes(a []byte) []int {
	var res []int
	for i := 0; i < len(a); i++ {
		for j := 7; j > 0; j-- {
			if a[i]&(1<<uint(j)) > 0 {
				res = append(res, (7-j)+(i*8))
			}
		}
	}
	return res
}

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

func andnot(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & ^b[i]
	}
}

func andnotInlined(a []byte, b []byte, res []byte) {
	i := 0
	l := len(a)

loop:
	res[i] = a[i] & ^b[i]
	i++
	if i != l {
		goto loop
	}
}
