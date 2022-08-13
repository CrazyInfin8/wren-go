package internals

import (
	"encoding/binary"
)

func f18(ctx *Context) int64 {
	var l0 int32
	_ = l0
	var l1 int64
	_ = l1
	var l2 int64
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	l0 = s0i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// const
	s0i32 = 1
	// const
	s1i64 = 0
	// get_local
	s2i32 = l0
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// call
	s0i32 = f544(ctx, s0i32, s1i64, s2i32)
	// const
	s0i32 = 0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+41616):]))
	// set_local
	l1 = s0i64
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l2 = s0i64
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i64 = l2
	// get_local
	s1i64 = l1
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// return
	return s0i64
}
