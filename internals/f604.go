package internals

import (
	"encoding/binary"
)

func f604(ctx *Context) int32 {
	var l0 int64
	_ = l0
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+43160):]))
	// const
	s2i64 = 6364136223846793005
	// binary: i64.mul
	s1i64 = s1i64 * s2i64
	// const
	s2i64 = 1
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// tee_local
	l0 = s1i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+43160):], uint64(s1i64))
	// get_local
	s0i64 = l0
	// const
	s1i64 = 33
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// return
	return s0i32
}
