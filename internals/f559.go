package internals

import (
	"encoding/binary"
)

func f559(ctx *Context, l0 int32, l1 int64, l2 int32) int64 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l1
	// get_local
	s2i32 = l2
	// const
	s3i32 = 255
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// call
	s0i32 = f545(ctx, s0i32, s1i64, s2i32, s3i32)
	// tee_local
	l0 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 70
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l0
	// const
	s4i32 = 76
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// const
	s0i64 = -1
	// set_local
	l1 = s0i64
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l1 = s0i64
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i64 = l1
	// return
	return s0i64
}
