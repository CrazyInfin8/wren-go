package internals

import (
	"encoding/binary"
)

func f569(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+43188):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 43164
	// set_local
	l1 = s0i32
	// const
	s0i32 = 0
	// const
	s1i32 = 43164
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+43188):], uint32(s1i32))
	// end_block
lbl0:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l0
	// const
	s3i32 = 76
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 22688
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load16_u
	s0i32 = int32(binary.LittleEndian.Uint16(ctx.Mem[int(s0i32+0):]))
	// const
	s1i32 = 21128
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+20):]))
	// call
	s0i32 = f588(ctx, s0i32, s1i32)
	// return
	return s0i32
}
