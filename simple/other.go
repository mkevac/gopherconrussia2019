package simple

import "math/rand"

func fill(b []byte, probability float32) {
	for i := 0; i < len(b); i++ {
		for j := uint(0); j < 8; j++ {
			if rand.Float32() < probability {
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
