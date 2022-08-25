package internals

import (
	"encoding/binary"
)

func f558(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 204
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 64
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// call
	s0i32 = f557(ctx, s0i32)
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+64):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	s0i32 = f556(ctx, s0i32, s1i32, s2i32)
	// return
	return s0i32
}
