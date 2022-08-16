package internals

import (
	"encoding/binary"
)

func f597(ctx *Context, l0 int32, l1 int32) int64 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
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
	var s2i64 int64
	_ = s2i64
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
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
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l3 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f594(ctx, s0i32)
	// set_local
	l3 = s0i32
	// end_block
lbl0:
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = -43
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl6
	case 1:
		goto lbl5
	case 2:
		goto lbl6
	default:
		goto lbl5
	}
	// end_block
lbl6:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
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
		goto lbl8
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l2 = s0i32
	// br
	goto lbl7
	// end_block
lbl8:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f594(ctx, s0i32)
	// set_local
	l2 = s0i32
	// end_block
lbl7:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 45
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -58
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
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
		goto lbl4
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = -11
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
		goto lbl3
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
	// br
	goto lbl3
	// end_block
lbl5:
	// get_local
	s0i32 = l3
	// const
	s1i32 = -58
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// const
	s0i32 = 0
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// set_local
	l2 = s0i32
	// end_block
lbl4:
	// get_local
	s0i32 = l5
	// const
	s1i32 = -10
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
	l6 = s0i64
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
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
		goto lbl9
	}
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// loop
lbl10:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
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
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l2 = s0i32
	// br
	goto lbl11
	// end_block
lbl12:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f594(ctx, s0i32)
	// set_local
	l2 = s0i32
	// end_block
lbl11:
	// get_local
	s0i32 = l3
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
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
		goto lbl13
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 214748364
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// end_block
lbl13:
	// end
	// get_local
	s0i32 = l3
	// unary: i64.extend_s/i32
	s0i64 = int64(s0i32)
	// set_local
	l6 = s0i64
	// end_block
lbl9:
	// block
	// get_local
	s0i32 = l5
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
		goto lbl14
	}
	// loop
lbl15:
	// get_local
	s0i32 = l2
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// get_local
	s1i64 = l6
	// const
	s2i64 = 10
	// binary: i64.mul
	s1i64 = s1i64 * s2i64
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l6 = s0i64
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
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
		goto lbl17
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l2 = s0i32
	// br
	goto lbl16
	// end_block
lbl17:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f594(ctx, s0i32)
	// set_local
	l2 = s0i32
	// end_block
lbl16:
	// get_local
	s0i64 = l6
	// const
	s1i64 = -48
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l6 = s0i64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
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
		goto lbl14
	}
	// get_local
	s0i64 = l6
	// const
	s1i64 = 92233720368547758
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl15
	}
	// end
	// end_block
lbl14:
	// block
	// get_local
	s0i32 = l5
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
		goto lbl18
	}
	// loop
lbl19:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
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
		goto lbl21
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l2 = s0i32
	// br
	goto lbl20
	// end_block
lbl21:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f594(ctx, s0i32)
	// set_local
	l2 = s0i32
	// end_block
lbl20:
	// get_local
	s0i32 = l2
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 10
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// end
	// end_block
lbl18:
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
	// const
	s0i64 = 0
	// get_local
	s1i64 = l6
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// get_local
	s1i64 = l6
	// get_local
	s2i32 = l4
	// select
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	// set_local
	l6 = s0i64
	// br
	goto lbl2
	// end_block
lbl3:
	// const
	s0i64 = -9223372036854775808
	// set_local
	l6 = s0i64
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
		goto lbl2
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
	// const
	s0i64 = -9223372036854775808
	// return
	return s0i64
	// end_block
lbl2:
	// get_local
	s0i64 = l6
	// return
	return s0i64
}
