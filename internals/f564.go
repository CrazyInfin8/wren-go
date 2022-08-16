package internals

import (
	"encoding/binary"
)

func f564(ctx *Context, l0 int32, l1 int64, l2 int32) int64 {
	var s0i32 int32
	_ = s0i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i64 = l1
	// get_local
	s2i32 = l2
	// call
	s0i64 = f563(ctx, s0i32, s1i64, s2i32)
	// return
	return s0i64
}
