package biggerbatch

import "math/rand"

func fill(r *rand.Rand, b []uint64, probability float32) {
	for i := 0; i < len(b); i++ {
		for j := uint64(0); j < 64; j++ {
			if r.Float32() < probability {
				b[i] |= 1 << j
			}
		}
	}
}

func indexes(a []uint64) []int {
	var res []int
	for i := 0; i < len(a); i++ {
		for j := 63; j > 0; j-- {
			if a[i]&(1<<uint64(j)) > 0 {
				res = append(res, (63-j)+(i*64))
			}
		}
	}
	return res
}

func and(a []uint64, b []uint64, res []uint64) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & b[i]
	}
}

func andInlined(a []uint64, b []uint64, res []uint64) {
	i := 0

loop:
	if i < len(a) {
		res[i] = a[i] & b[i]
		i++
		goto loop
	}
}

func andNoBoundsCheck(a []uint64, b []uint64, res []uint64) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	for i := 0; i < len(a); i++ {
		res[i] = a[i] & b[i]
	}
}

func andInlinedAndNoBoundsCheck(a []uint64, b []uint64, res []uint64) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	i := 0

loop:
	if i < len(a) {
		res[i] = a[i] & b[i]
		i++
		goto loop
	}
}

func andnot(a []uint64, b []uint64, res []uint64) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & ^b[i]
	}
}

func andnotInlined(a []uint64, b []uint64, res []uint64) {
	i := 0

loop:
	if i < len(a) {
		res[i] = a[i] & ^b[i]
		i++
		goto loop
	}
}

func andnotNoBoundsCheck(a []uint64, b []uint64, res []uint64) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	for i := 0; i < len(a); i++ {
		res[i] = a[i] & ^b[i]
	}
}

func andnotInlinedAndNoBoundsCheck(a []uint64, b []uint64, res []uint64) {
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

	i := 0

loop:
	if i < len(a) {
		res[i] = a[i] & ^b[i]
		i++
		goto loop
	}
}
