package internals

import (
	"encoding/binary"
)

func f55(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
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
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_global
	s0i32 = ctx.G0
	// set_local
	l1 = s0i32
	// const
	s0i32 = 16
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l3
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// call
	s0i32 = f38(ctx, s0i32)
	// set_local
	l5 = s0i32
	// const
	s0i32 = 61
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l9 = s0i32
	// const
	s0i32 = 1
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l10
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l11 = s0i32
	// block
	// get_local
	s0i32 = l11
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l12
	// call
	s0i32 = f486(ctx, s0i32)
	// set_local
	l13 = s0i32
	// const
	s0i32 = 7
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l13
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l14
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l15
	// get_local
	s1i32 = l16
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l17 = s0i32
	// const
	s0i32 = 1
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l18
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l19
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l20
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l21
	// call
	f25(ctx, s0i32)
	// end_block
lbl0:
	// const
	s0i32 = 16
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l22
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l23
	// set_global
	ctx.G0 = s0i32
	// return
	return
}