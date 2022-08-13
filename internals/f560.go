package internals

import (
	"encoding/binary"
)

func f560(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l1 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// call
	s0i32 = f546(ctx, s0i32, s1i32)
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 59
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	// const
	s1i32 = 2
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+16)])
	// const
	s1i32 = 36
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 1
	// set_local
	l2 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42124):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l2
	// return
	return s0i32
}
