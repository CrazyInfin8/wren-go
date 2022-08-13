package internals

import (
	"encoding/binary"
)

func f591(ctx *Context, l0 int32) int32 {
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
	var s3i32 int32
	_ = s3i32
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l1 = s0i32
	// set_global
	ctx.G0 = s0i32
	// const
	s0i32 = -1
	// set_local
	l2 = s0i32
	// block
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 15
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 1
	// get_local
	s3i32 = l0
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+28):]))
	// call_indirect
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].f == nil {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = table[s3i32].f.(func(*Context, int32, int32, int32) int32)(ctx, s0i32, s1i32, s2i32)
	// const
	s1i32 = 1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+15)])
	// set_local
	l2 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l2
	// return
	return s0i32
}
