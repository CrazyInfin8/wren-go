package internals

import (
	"encoding/binary"
)

func f596(ctx *Context, l0 int32, l1 int32, l2 int32) int64 {
	var l3 int32
	_ = l3
	var l4 int64
	_ = l4
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
	var s3i64 int64
	_ = s3i64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 112
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// const
	s3i64 = -9223372036854775808
	// call
	s0i64 = f595(ctx, s0i32, s1i32, s2i32, s3i64)
	// set_local
	l4 = s0i64
	// block
	// get_local
	s0i32 = l1
	// unary: i32.eqz
	if s0i32 == 0 {
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
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l3
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// get_local
	s3i32 = l3
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+96):]))
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l3
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+40):]))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i64 = l4
	// return
	return s0i64
}
