package internals

import (
	"encoding/binary"
)

func f634(ctx *Context, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l4
	// get_local
	s2i64 = l1
	// binary: i64.mul
	s1i64 = s1i64 * s2i64
	// get_local
	s2i64 = l2
	// get_local
	s3i64 = l3
	// binary: i64.mul
	s2i64 = s2i64 * s3i64
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// get_local
	s2i64 = l3
	// const
	s3i64 = 32
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// tee_local
	l4 = s2i64
	// get_local
	s3i64 = l1
	// const
	s4i64 = 32
	// binary: i64.shr_u
	s3i64 = int64(uint64(s3i64) >> (uint64(s4i64) & 63))
	// tee_local
	l2 = s3i64
	// binary: i64.mul
	s2i64 = s2i64 * s3i64
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// get_local
	s2i64 = l3
	// const
	s3i64 = 4294967295
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// tee_local
	l3 = s2i64
	// get_local
	s3i64 = l1
	// const
	s4i64 = 4294967295
	// binary: i64.and
	s3i64 = s3i64 & s4i64
	// tee_local
	l1 = s3i64
	// binary: i64.mul
	s2i64 = s2i64 * s3i64
	// tee_local
	l5 = s2i64
	// const
	s3i64 = 32
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// get_local
	s3i64 = l3
	// get_local
	s4i64 = l2
	// binary: i64.mul
	s3i64 = s3i64 * s4i64
	// binary: i64.add
	s2i64 = s2i64 + s3i64
	// tee_local
	l3 = s2i64
	// const
	s3i64 = 32
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// get_local
	s2i64 = l3
	// const
	s3i64 = 4294967295
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// get_local
	s3i64 = l4
	// get_local
	s4i64 = l1
	// binary: i64.mul
	s3i64 = s3i64 * s4i64
	// binary: i64.add
	s2i64 = s2i64 + s3i64
	// tee_local
	l3 = s2i64
	// const
	s3i64 = 32
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l3
	// const
	s2i64 = 32
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// get_local
	s2i64 = l5
	// const
	s3i64 = 4294967295
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// binary: i64.or
	s1i64 = s1i64 | s2i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
}
