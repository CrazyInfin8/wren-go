package internals

import (
	"encoding/binary"
)

func f603(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+43160):], uint64(s1i64))
}
