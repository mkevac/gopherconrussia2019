// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

var unroll = 8

func genSIMDop(op string) {

	switch op {
	case "and":
		TEXT("andSIMD", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "or":
		TEXT("orSIMD", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "andnot":
		TEXT("andnotSIMD", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	}

	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
	res := Mem{Base: Load(Param("res").Base(), GP64())}
	n := Load(Param("a").Len(), GP64())

	// Allocate accumulation registers.
	acc := make([]VecVirtual, unroll)
	for i := 0; i < unroll; i++ {
		acc[i] = YMM()
	}

	// Zero initialization.
	for i := 0; i < unroll; i++ {
		VXORPS(acc[i], acc[i], acc[i])
	}

	blocksize := 32 * unroll

	Label("blockloop")
	CMPQ(n, U32(blocksize))
	JL(LabelRef("tail"))

	// Load b.
	bs := make([]VecVirtual, unroll)
	for i := 0; i < unroll; i++ {
		bs[i] = YMM()
	}

	for i := 0; i < unroll; i++ {
		VMOVUPS(b.Offset(32*i), bs[i])
	}

	// The actual operation.
	for i := 0; i < unroll; i++ {
		switch op {
		case "and":
			VPAND(a.Offset(32*i), bs[i], bs[i])
		case "or":
			VPOR(a.Offset(32*i), bs[i], bs[i])
		case "andnot":
			VPANDN(a.Offset(32*i), bs[i], bs[i])
		}
	}

	for i := 0; i < unroll; i++ {
		VMOVUPS(bs[i], res.Offset(32*i))
	}

	ADDQ(U32(blocksize), a.Base)
	ADDQ(U32(blocksize), b.Base)
	ADDQ(U32(blocksize), res.Base)
	SUBQ(U32(blocksize), n)
	JMP(LabelRef("blockloop"))

	Label("tail")
	RET()
}

func main() {
	genSIMDop("and")
	genSIMDop("or")
	genSIMDop("andnot")
	Generate()
}
