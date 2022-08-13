package internals

import (
	"encoding/binary"
)

func f590(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+60):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+60):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+24):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// const
	s2i32 = 0
	// get_local
	s3i32 = l0
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+32):]))
	// call_indirect
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].f == nil {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = table[s3i32].f.(func(*Context, int32, int32, int32) int32)(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl0:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 4
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = -1
	// return
	return s0i32
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+44):]))
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 27
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 31
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// return
	return s0i32
}
