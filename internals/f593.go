package internals

import (
	"encoding/binary"
)

func f593(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
	var l4 int64
	_ = l4
	var l5 int32
	_ = l5
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
	var s2i64 int64
	_ = s2i64
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+96):]))
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// tee_local
	l1 = s1i32
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+40):]))
	// tee_local
	l2 = s2i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l3 = s0i64
	// block
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// tee_local
	l4 = s0i64
	// unary: i64.eqz
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i64 = l3
	// get_local
	s1i64 = l4
	// binary: i64.ge_s
	if s0i64 >= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// end_block
lbl2:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f591(ctx, s0i32)
	// tee_local
	l2 = s0i32
	// const
	s1i32 = -1
	// binary: i32.gt_s
	if s0i32 > s1i32 {
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
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l2 = s0i32
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// const
	s1i64 = -1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+88):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+84):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l3
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l1
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// unary: i64.extend_s/i32
	s2i64 = int64(s2i32)
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+96):], uint64(s1i64))
	// const
	s0i32 = -1
	// return
	return s0i32
	// end_block
lbl0:
	// get_local
	s0i64 = l3
	// const
	s1i64 = 1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l3 = s0i64
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l5 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// tee_local
	l4 = s0i64
	// const
	s1i64 = 0
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i64 = l4
	// get_local
	s1i64 = l3
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// tee_local
	l4 = s0i64
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l1
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// binary: i64.ge_s
	if s0i64 >= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i64 = l4
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl3:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+84):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l3
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+40):]))
	// tee_local
	l5 = s2i32
	// get_local
	s3i32 = l1
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// unary: i64.extend_s/i32
	s2i64 = int64(s2i32)
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+96):], uint64(s1i64))
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// end_block
lbl4:
	// get_local
	s0i32 = l2
	// return
	return s0i32
}
