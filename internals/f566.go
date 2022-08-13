package internals

import (
	"encoding/binary"
)

func f566(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 8
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
		goto lbl0
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
lbl0:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+4):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+44):]))
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// const
	s0i32 = 0
	// return
	return s0i32
}
