// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func genScalarFasterOp(op string) {
	switch op {
	case "and":
		TEXT("andScalarFaster", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "or":
		TEXT("orScalarFaster", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "andnot":
		TEXT("andnotScalarFaster", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	}

	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
	res := Mem{Base: Load(Param("res").Base(), GP64())}
	n := Load(Param("a").Len(), GP64())

	ir := GP64()
	XORQ(ir, ir)

	ar := GP64()
	br := GP64()

	Label("loop")
	CMPQ(n, ir)
	JE(LabelRef("done"))

	MOVQ(a.Idx(ir, 1), ar)
	MOVQ(b.Idx(ir, 1), br)

	switch op {
	case "and":
		ANDQ(ar, br)
	case "or":
		ORQ(ar, br)
	case "andnot":
		ANDNQ(ar, br, br)
	}

	MOVQ(br, res.Idx(ir, 1))

	ADDQ(Imm(8), ir)

	JMP(LabelRef("loop"))

	Label("done")
	RET()
}

func genScalarOp(op string) {
	switch op {
	case "and":
		TEXT("andScalar", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "or":
		TEXT("orScalar", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	case "andnot":
		TEXT("andnotScalar", NOSPLIT, "func(a []byte, b []byte, res []byte)")
	}

	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
	res := Mem{Base: Load(Param("res").Base(), GP64())}
	n := Load(Param("a").Len(), GP64())

	s := GP64()
	XORQ(s, s)

	Label("loop")
	Comment("Loop until zero bytes remain.")
	CMPQ(n, Imm(0))
	JE(LabelRef("done"))

	MOVQ(Mem{Base: b.Base}, s)

	switch op {
	case "and":
		ANDQ(Mem{Base: a.Base}, s)
	case "or":
		ORQ(Mem{Base: a.Base}, s)
	case "andnot":
		ANDNQ(Mem{Base: a.Base}, s, s)
	}

	MOVQ(s, Mem{Base: res.Base})

	ADDQ(U32(8), a.Base)
	ADDQ(U32(8), b.Base)
	ADDQ(U32(8), res.Base)
	SUBQ(U32(8), n)
	JMP(LabelRef("loop"))

	Label("done")
	RET()
}

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

	genScalarOp("and")
	genScalarOp("or")
	genScalarOp("andnot")

	genScalarFasterOp("and")
	genScalarFasterOp("or")
	genScalarFasterOp("andnot")

	Generate()
}
