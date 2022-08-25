package internals

import (
	"encoding/binary"
)

func f577(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
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
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 3
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
	s0i32 = l2
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
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
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
	s0i32 = l3
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
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+1)])
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// get_local
	s0i32 = l2
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
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
	s0i32 = l3
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
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+2)])
	// store: i32.store8
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	// get_local
	s0i32 = l2
	// const
	s1i32 = -3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
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
	s0i32 = l3
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
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+3)])
	// store: i32.store8
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l2
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l1
	// set_local
	l5 = s0i32
	// end_block
lbl0:
	// block
	// block
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// block
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 16
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 16
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l1
	// set_local
	l3 = s0i32
	// end_block
lbl7:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// loop
lbl8:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 15
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// end
	// end_block
lbl6:
	// get_local
	s0i32 = l3
	// set_local
	l1 = s0i32
	// end_block
lbl5:
	// block
	// get_local
	s0i32 = l1
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
		goto lbl9
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l5
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// end_block
lbl9:
	// block
	// get_local
	s0i32 = l1
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
		goto lbl10
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// end_block
lbl10:
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2
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
		goto lbl11
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load16_u
	s1i32 = int32(binary.LittleEndian.Uint16(ctx.Mem[int(s1i32+0):]))
	// store: i32.store16
	binary.LittleEndian.PutUint16(ctx.Mem[int(s0i32+0):], uint16(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl11:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// br
	goto lbl2
	// end_block
lbl4:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 32
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl15
	case 1:
		goto lbl14
	case 2:
		goto lbl13
	default:
		goto lbl12
	}
	// end_block
lbl15:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l6
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// store: i32.store8
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l6
	// const
	s2i32 = 8
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// get_local
	s0i32 = l3
	// const
	s1i32 = -3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// loop
lbl16:
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 8
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 24
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 8
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 24
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 12
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 8
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 24
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 8
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 24
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 16
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl16
	}
	// end
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// br
	goto lbl12
	// end_block
lbl14:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// store: i32.store16
	binary.LittleEndian.PutUint16(ctx.Mem[int(s0i32+0):], uint16(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// loop
lbl17:
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 16
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 16
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 12
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 16
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 16
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 17
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl17
	}
	// end
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// br
	goto lbl12
	// end_block
lbl13:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l3
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// loop
lbl18:
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 24
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 24
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 12
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 24
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l6
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 24
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l8
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 18
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl18
	}
	// end
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl12:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 16
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
		goto lbl19
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+1):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+1):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+5):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+5):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load16_u
	s1i32 = int32(binary.LittleEndian.Uint16(ctx.Mem[int(s1i32+13):]))
	// store: i32.store16
	binary.LittleEndian.PutUint16(ctx.Mem[int(s0i32+13):], uint16(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+15)])
	// store: i32.store8
	ctx.Mem[int(s0i32+15)] = uint8(s1i32)
	// get_local
	s0i32 = l4
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl19:
	// block
	// get_local
	s0i32 = l3
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
		goto lbl20
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl20:
	// block
	// get_local
	s0i32 = l3
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
		goto lbl21
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl21:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2
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
		goto lbl22
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load16_u
	s1i32 = int32(binary.LittleEndian.Uint16(ctx.Mem[int(s1i32+0):]))
	// store: i32.store16
	binary.LittleEndian.PutUint16(ctx.Mem[int(s0i32+0):], uint16(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl22:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
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
		goto lbl2
	}
	// end_block
lbl3:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// end_block
lbl2:
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
