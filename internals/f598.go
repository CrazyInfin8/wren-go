package internals

import (
	"encoding/binary"
)

func f598(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
	var l12 int32
	_ = l12
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
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l4 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 36
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// block
	// block
	// loop
lbl8:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl9
	// end_block
lbl10:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl9:
	// get_local
	s0i32 = l5
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 5
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl8
	case 1:
		goto lbl7
	case 2:
		goto lbl7
	case 3:
		goto lbl7
	case 4:
		goto lbl7
	case 5:
		goto lbl7
	case 6:
		goto lbl7
	case 7:
		goto lbl7
	case 8:
		goto lbl7
	case 9:
		goto lbl7
	case 10:
		goto lbl7
	case 11:
		goto lbl11
	case 12:
		goto lbl7
	case 13:
		goto lbl11
	default:
		goto lbl7
	}
	// end_block
lbl11:
	// end
	// const
	s0i32 = -1
	// const
	s1i32 = 0
	// get_local
	s2i32 = l5
	// const
	s3i32 = 45
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l6 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl6
	// end_block
lbl12:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// br
	goto lbl6
	// end_block
lbl7:
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// end_block
lbl6:
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -17
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 48
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl16
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl15
	// end_block
lbl16:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl15:
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 88
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl17
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl18
	// end_block
lbl19:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl18:
	// const
	s0i32 = 16
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 23377
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
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
		goto lbl3
	}
	// const
	s0i64 = 0
	// set_local
	l3 = s0i64
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// tee_local
	l5 = s1i32
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
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
		goto lbl20
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = -2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl21:
	// get_local
	s0i32 = l2
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// end_block
lbl20:
	// const
	s0i64 = 0
	// set_local
	l3 = s0i64
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f591(ctx, s0i32, s1i64)
	// br
	goto lbl0
	// end_block
lbl17:
	// get_local
	s0i32 = l1
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// const
	s0i32 = 8
	// set_local
	l1 = s0i32
	// br
	goto lbl3
	// end_block
lbl14:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 10
	// get_local
	s2i32 = l1
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// const
	s0i64 = 0
	// set_local
	l3 = s0i64
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl22
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// end_block
lbl22:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f591(ctx, s0i32, s1i64)
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42156):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl13:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 10
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// const
	s0i64 = 0
	// set_local
	l7 = s0i64
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// loop
lbl24:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 10
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l1 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl26
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl25
	// end_block
lbl26:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl25:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl27
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 429496729
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl24
	}
	// end_block
lbl27:
	// end
	// get_local
	s0i32 = l1
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l7 = s0i64
	// end_block
lbl23:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 9
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
	s0i64 = l7
	// const
	s1i64 = 10
	// binary: i64.mul
	s0i64 = s0i64 * s1i64
	// set_local
	l8 = s0i64
	// get_local
	s0i32 = l2
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l9 = s0i64
	// loop
lbl28:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl30
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl29
	// end_block
lbl30:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl29:
	// get_local
	s0i64 = l8
	// get_local
	s1i64 = l9
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l7 = s0i64
	// get_local
	s0i32 = l5
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 9
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
	s0i64 = l7
	// const
	s1i64 = 1844674407370955162
	// binary: i64.ge_u
	if uint64(s0i64) >= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i64 = l7
	// const
	s1i64 = 10
	// binary: i64.mul
	s0i64 = s0i64 * s1i64
	// tee_local
	l8 = s0i64
	// get_local
	s1i32 = l2
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// tee_local
	l9 = s1i64
	// const
	s2i64 = -1
	// binary: i64.xor
	s1i64 = s1i64 ^ s2i64
	// binary: i64.le_u
	if uint64(s0i64) <= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl28
	}
	// end
	// const
	s0i32 = 10
	// set_local
	l1 = s0i32
	// br
	goto lbl2
	// end_block
lbl5:
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42156):], uint32(s1i32))
	// const
	s0i64 = 0
	// set_local
	l3 = s0i64
	// br
	goto lbl0
	// end_block
lbl4:
	// const
	s0i32 = 10
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 9
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// br
	goto lbl1
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l1
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
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
		goto lbl31
	}
	// const
	s0i64 = 0
	// set_local
	l7 = s0i64
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l10 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl32
	}
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// loop
lbl33:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l2 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl35
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl34
	// end_block
lbl35:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl34:
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l10 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 119304647
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl33
	}
	// end_block
lbl36:
	// end
	// get_local
	s0i32 = l2
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l7 = s0i64
	// end_block
lbl32:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l10
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l1
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l8 = s0i64
	// loop
lbl37:
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l8
	// binary: i64.mul
	s0i64 = s0i64 * s1i64
	// tee_local
	l9 = s0i64
	// get_local
	s1i32 = l10
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// const
	s2i64 = 255
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// tee_local
	l11 = s1i64
	// const
	s2i64 = -1
	// binary: i64.xor
	s1i64 = s1i64 ^ s2i64
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl39
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl38
	// end_block
lbl39:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl38:
	// get_local
	s0i64 = l9
	// get_local
	s1i64 = l11
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l7 = s0i64
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l10 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l8
	// const
	s2i64 = 0
	// get_local
	s3i64 = l7
	// const
	s4i64 = 0
	// call
	f634(ctx, s0i32, s1i64, s2i64, s3i64, s4i64)
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// const
	s1i64 = 0
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// br
	goto lbl37
	// end
	// end_block
lbl31:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 23
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// const
	s1i32 = 5
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// const
	s1i32 = 7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 23633
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	// set_local
	l12 = s0i32
	// const
	s0i64 = 0
	// set_local
	l7 = s0i64
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l2 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl40
	}
	// const
	s0i32 = 0
	// set_local
	l10 = s0i32
	// loop
lbl41:
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l12
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l10 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl43
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl42
	// end_block
lbl43:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl42:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l10
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l10 = s0i32
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l2 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl44
	}
	// get_local
	s0i32 = l10
	// const
	s1i32 = 134217728
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl41
	}
	// end_block
lbl44:
	// end
	// get_local
	s0i32 = l10
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l7 = s0i64
	// end_block
lbl40:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// const
	s0i64 = -1
	// get_local
	s1i32 = l12
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// tee_local
	l9 = s1i64
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// tee_local
	l11 = s0i64
	// get_local
	s1i64 = l7
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// loop
lbl45:
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l9
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l7 = s0i64
	// get_local
	s0i32 = l2
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// const
	s1i64 = 255
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l8 = s0i64
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl47
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl46
	// end_block
lbl47:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl46:
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l8
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// set_local
	l7 = s0i64
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// tee_local
	l2 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l11
	// binary: i64.le_u
	if uint64(s0i64) <= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// end
	// end_block
lbl2:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// loop
lbl48:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+84):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl50
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l5 = s0i32
	// br
	goto lbl49
	// end_block
lbl50:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f592(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl49:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 23377
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl48
	}
	// end
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42156):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// const
	s1i32 = 0
	// get_local
	s2i64 = l3
	// const
	s3i64 = 1
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// unary: i64.eqz
	if s2i64 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l6 = s0i32
	// get_local
	s0i64 = l3
	// set_local
	l7 = s0i64
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl51
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// end_block
lbl51:
	// block
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l3
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl52
	}
	// block
	// get_local
	s0i64 = l3
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl53
	}
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl53
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42156):], uint32(s1i32))
	// get_local
	s0i64 = l3
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l3 = s0i64
	// br
	goto lbl0
	// end_block
lbl53:
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l3
	// binary: i64.le_u
	if uint64(s0i64) <= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl52
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42156):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl52:
	// get_local
	s0i64 = l7
	// get_local
	s1i32 = l6
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// tee_local
	l3 = s1i64
	// binary: i64.xor
	s0i64 = s0i64 ^ s1i64
	// get_local
	s1i64 = l3
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// set_local
	l3 = s0i64
	// end_block
lbl0:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i64 = l3
	// return
	return s0i64
}
