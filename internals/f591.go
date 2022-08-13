package internals

import (
	"encoding/binary"
)

func f591(ctx *Context, l0 int32, l1 int64) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+88):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// tee_local
	l2 = s2i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+96):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i64 = l1
	// unary: i64.eqz
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// unary: i64.extend_s/i32
	s0i64 = int64(s0i32)
	// get_local
	s1i64 = l1
	// binary: i64.le_s
	if s0i64 <= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l2
	// get_local
	s1i64 = l1
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+84):], uint32(s1i32))
}
