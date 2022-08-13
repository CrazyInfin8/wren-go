package internals

import (
	"encoding/binary"
)

func f556(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// block
	// get_local
	s0i32 = l0
	// call
	s0i32 = f545(ctx, s0i32)
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 0
	// return
	return s0i32
	// end_block
lbl0:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42124):], uint32(s1i32))
	// const
	s0i32 = -1
	// return
	return s0i32
}
