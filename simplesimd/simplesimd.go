//go:generate go run asm.go -out simd.s -stubs stub.go

package simplesimd

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
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	i := 0
	l := len(a)

loop:
	res[i] = a[i] & b[i]
	i++
	if i != l {
		goto loop
	}
}

func andnot(a []byte, b []byte, res []byte) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	i := 0
	l := len(a)

loop:
	res[i] = a[i] & ^b[i]
	i++
	if i != l {
		goto loop
	}
}

func or(a []byte, b []byte, res []byte) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	i := 0
	l := len(a)

loop:
	res[i] = a[i] | b[i]
	i++
	if i != l {
		goto loop
	}
}
