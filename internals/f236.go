package internals

import (
	"encoding/binary"
	"math"
)

func f236(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
	var l9 int32
	_ = l9
	var l10 float64
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 float64
	_ = l20
	var l21 int64
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 16
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l6 = s0i64
	// get_local
	s0i64 = l6
	// call
	s0f64 = f98(ctx, s0i64)
	// set_local
	l7 = s0f64
	// const
	s0f64 = 4.294967296e+09
	// set_local
	l8 = s0f64
	// get_local
	s0f64 = l7
	// get_local
	s1f64 = l8
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l9 = s0i32
	// const
	s0f64 = 0
	// set_local
	l10 = s0f64
	// get_local
	s0f64 = l7
	// get_local
	s1f64 = l10
	// binary: f64.ge
	if s0f64 >= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l11
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l12
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l13 = s0i32
	// block
	// block
	// get_local
	s0i32 = l13
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0f64 = l7
	// unary: i32.trunc_u/f64
	s0i32 = int32(uint32(math.Trunc(s0f64)))
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l14
	// set_local
	l15 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// const
	s0i32 = 0
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// set_local
	l15 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l15
	// set_local
	l17 = s0i32
	// const
	s0i32 = -1
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l18
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l19
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l20 = s0f64
	// get_local
	s0f64 = l20
	// call
	s0i64 = f326(ctx, s0f64)
	// set_local
	l21 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l22
	// get_local
	s1i64 = l21
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// const
	s0i32 = 1
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l24
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l25 = s0i32
	// const
	s0i32 = 16
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l26
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l27
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l25
	// return
	return s0i32
}
