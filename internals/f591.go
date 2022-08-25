package internals

import (
	"encoding/binary"
	"math"
)

func f591(ctx *Context, l0 int32, l1 int32, l2 int32) float64 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float64
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int64
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int64
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int64
	_ = l22
	var l23 float64
	_ = l23
	var l24 float64
	_ = l24
	var l25 float64
	_ = l25
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 512
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// set_global
	ctx.G0 = s0i32
	// const
	s0i32 = -149
	// set_local
	l4 = s0i32
	// const
	s0i32 = 24
	// set_local
	l5 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// const
	s0f64 = 0
	// set_local
	l7 = s0f64
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// br_table
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl2
	case 2:
		goto lbl2
	default:
		goto lbl0
	}
	// end_block
lbl2:
	// const
	s0i32 = -1074
	// set_local
	l4 = s0i32
	// const
	s0i32 = 53
	// set_local
	l5 = s0i32
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l8 = s0i32
	// block
	// block
	// loop
lbl5:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl7
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl6
	// end_block
lbl7:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl6:
	// get_local
	s0i32 = l1
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
		goto lbl5
	}
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl4
	case 2:
		goto lbl4
	case 3:
		goto lbl4
	case 4:
		goto lbl4
	case 5:
		goto lbl4
	case 6:
		goto lbl4
	case 7:
		goto lbl4
	case 8:
		goto lbl4
	case 9:
		goto lbl4
	case 10:
		goto lbl4
	case 11:
		goto lbl8
	case 12:
		goto lbl4
	case 13:
		goto lbl8
	default:
		goto lbl4
	}
	// end_block
lbl8:
	// end
	// const
	s0i32 = -1
	// const
	s1i32 = 1
	// get_local
	s2i32 = l1
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
	l9 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl9
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl3
	// end_block
lbl9:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// br
	goto lbl3
	// end_block
lbl4:
	// const
	s0i32 = 1
	// set_local
	l9 = s0i32
	// end_block
lbl3:
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l10 = s0i32
	// const
	s1i32 = 73
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl14
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl13
	// end_block
lbl14:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl13:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 78
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl15
	// end_block
lbl16:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl15:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 70
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl18
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l10 = s0i32
	// br
	goto lbl17
	// end_block
lbl18:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l10 = s0i32
	// end_block
lbl17:
	// const
	s0i32 = 3
	// set_local
	l1 = s0i32
	// block
	// block
	// get_local
	s0i32 = l10
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l10 = s0i32
	// const
	s1i32 = 73
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl20
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl22
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l11 = s0i32
	// br
	goto lbl21
	// end_block
lbl22:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l11 = s0i32
	// end_block
lbl21:
	// const
	s0i32 = 4
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l11
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 78
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl25
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l11 = s0i32
	// br
	goto lbl24
	// end_block
lbl25:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l11 = s0i32
	// end_block
lbl24:
	// const
	s0i32 = 5
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l11
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 73
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl27
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l11 = s0i32
	// br
	goto lbl26
	// end_block
lbl27:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l11 = s0i32
	// end_block
lbl26:
	// const
	s0i32 = 6
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l11
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 84
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl29
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l11 = s0i32
	// br
	goto lbl28
	// end_block
lbl29:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l11 = s0i32
	// end_block
lbl28:
	// const
	s0i32 = 7
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l11
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 89
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
	// end_block
lbl23:
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
		goto lbl11
	}
	// end_block
lbl20:
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// tee_local
	l12 = s0i64
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
		goto lbl30
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl30:
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
		goto lbl19
	}
	// get_local
	s0i32 = l10
	// const
	s1i32 = 73
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// block
	// get_local
	s0i64 = l12
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
		goto lbl31
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl31:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -5
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// block
	// get_local
	s0i64 = l12
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
		goto lbl32
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl32:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -5
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// block
	// get_local
	s0i64 = l12
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
		goto lbl33
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl33:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -7
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -5
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// get_local
	s0i64 = l12
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
		goto lbl19
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl19:
	// get_local
	s0i32 = l9
	// unary: f32.convert_s/i32
	s0f32 = float32(s0i32)
	// const
	s1f32 = float32(math.Inf(0))
	// binary: f32.mul
	s0f32 = s0f32 * s1f32
	// unary: f64.promote/f32
	s0f64 = float64(s0f32)
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl12:
	// get_local
	s0i32 = l10
	// const
	s1i32 = 78
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl34
	// end_block
lbl35:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl34:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 65
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl37
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl36
	// end_block
lbl37:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl36:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 78
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl38
	// end_block
lbl39:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl38:
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 40
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl41
	}
	// const
	s0i32 = 1
	// set_local
	l11 = s0i32
	// const
	s0i32 = 1
	// set_local
	l10 = s0i32
	// br
	goto lbl40
	// end_block
lbl41:
	// const
	s0f64 = math.NaN()
	// set_local
	l7 = s0f64
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
		goto lbl0
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl40:
	// loop
lbl42:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl44
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl43
	// end_block
lbl44:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl43:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -65
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// block
	// block
	// get_local
	s0i32 = l1
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
		goto lbl46
	}
	// get_local
	s0i32 = l13
	// const
	s1i32 = 26
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl46
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -97
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 95
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl46
	}
	// get_local
	s0i32 = l13
	// const
	s1i32 = 26
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// end_block
lbl46:
	// get_local
	s0i32 = l11
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l10
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l10 = s0i32
	// br
	goto lbl42
	// end_block
lbl45:
	// end
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 41
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl47
	}
	// const
	s0f64 = math.NaN()
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl47:
	// block
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// tee_local
	l12 = s0i64
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
		goto lbl48
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl48:
	// block
	// block
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
		goto lbl50
	}
	// block
	// get_local
	s0i32 = l10
	// br_if
	if s0i32 != 0 {
		goto lbl51
	}
	// const
	s0f64 = math.NaN()
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl51:
	// get_local
	s0i32 = l10
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// block
	// get_local
	s0i32 = l10
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
		goto lbl52
	}
	// get_local
	s0i32 = l11
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l1 = s0i32
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// loop
lbl53:
	// block
	// get_local
	s0i64 = l12
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
		goto lbl54
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl54:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l0
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl53
	}
	// end
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l10 = s0i32
	// end_block
lbl52:
	// get_local
	s0i32 = l13
	// const
	s1i32 = 3
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl49
	}
	// const
	s0f64 = math.NaN()
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl50:
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// br
	goto lbl0
	// end_block
lbl49:
	// get_local
	s0i64 = l12
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l0 = s0i32
	// loop
lbl55:
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -3
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl56:
	// get_local
	s0i32 = l10
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l10 = s0i32
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl57
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl57:
	// get_local
	s0i32 = l10
	// br_if
	if s0i32 != 0 {
		goto lbl55
	}
	// end
	// const
	s0f64 = math.NaN()
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl11:
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
		goto lbl58
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl58:
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// br
	goto lbl0
	// end_block
lbl10:
	// block
	// get_local
	s0i32 = l1
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
		goto lbl59
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l10 = s0i32
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
		goto lbl61
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l10
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l10
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l10 = s0i32
	// br
	goto lbl60
	// end_block
lbl61:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l10 = s0i32
	// end_block
lbl60:
	// block
	// get_local
	s0i32 = l10
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
		goto lbl62
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l4
	// get_local
	s3i32 = l9
	// get_local
	s4i32 = l2
	// call
	s0f64 = f592(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl62:
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
		goto lbl59
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl59:
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l5
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l14 = s1i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l15 = s0i32
	// block
	// block
	// loop
lbl65:
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 48
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl66
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 46
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl64
	}
	// const
	s0i32 = 0
	// set_local
	l16 = s0i32
	// const
	s0i64 = 0
	// set_local
	l12 = s0i64
	// br
	goto lbl63
	// end_block
lbl66:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl67
	}
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl65
	// end_block
lbl67:
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// br
	goto lbl65
	// end
	// end_block
lbl64:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl69
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl68
	// end_block
lbl69:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl68:
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 48
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl70
	}
	// const
	s0i32 = 1
	// set_local
	l16 = s0i32
	// const
	s0i64 = 0
	// set_local
	l12 = s0i64
	// br
	goto lbl63
	// end_block
lbl70:
	// const
	s0i64 = 0
	// set_local
	l12 = s0i64
	// loop
lbl71:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl73
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl72
	// end_block
lbl73:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl72:
	// get_local
	s0i64 = l12
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l12 = s0i64
	// get_local
	s0i32 = l1
	// const
	s1i32 = 48
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl71
	}
	// end
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// const
	s0i32 = 1
	// set_local
	l16 = s0i32
	// end_block
lbl63:
	// const
	s0i32 = 0
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 46
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l10 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl80
	}
	// const
	s0i64 = 0
	// set_local
	l18 = s0i64
	// get_local
	s0i32 = l11
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
		goto lbl80
	}
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// br
	goto lbl79
	// end_block
lbl80:
	// const
	s0i64 = 0
	// set_local
	l18 = s0i64
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// const
	s0i32 = 0
	// set_local
	l17 = s0i32
	// loop
lbl81:
	// block
	// block
	// get_local
	s0i32 = l10
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
		goto lbl83
	}
	// block
	// get_local
	s0i32 = l16
	// br_if
	if s0i32 != 0 {
		goto lbl84
	}
	// get_local
	s0i64 = l18
	// set_local
	l12 = s0i64
	// const
	s0i32 = 1
	// set_local
	l16 = s0i32
	// br
	goto lbl82
	// end_block
lbl84:
	// get_local
	s0i32 = l13
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l10 = s0i32
	// br
	goto lbl78
	// end_block
lbl83:
	// get_local
	s0i64 = l18
	// const
	s1i64 = 1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l18 = s0i64
	// block
	// get_local
	s0i32 = l19
	// const
	s1i32 = 124
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl85
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 48
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l13 = s0i32
	// get_local
	s0i64 = l18
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l19
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l10 = s0i32
	// block
	// get_local
	s0i32 = l20
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl86
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l10
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// end_block
lbl86:
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l13
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// const
	s0i32 = 0
	// get_local
	s1i32 = l20
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l1 = s1i32
	// get_local
	s2i32 = l1
	// const
	s3i32 = 9
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l1 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// br
	goto lbl82
	// end_block
lbl85:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 48
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl82
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+496):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+496):], uint32(s1i32))
	// const
	s0i32 = 1116
	// set_local
	l17 = s0i32
	// end_block
lbl82:
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
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
		goto lbl88
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l1 = s0i32
	// br
	goto lbl87
	// end_block
lbl88:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l1 = s0i32
	// end_block
lbl87:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 46
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l10 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl81
	}
	// get_local
	s0i32 = l11
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
		goto lbl81
	}
	// end
	// end_block
lbl79:
	// get_local
	s0i64 = l12
	// get_local
	s1i64 = l18
	// get_local
	s2i32 = l16
	// select
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	// set_local
	l12 = s0i64
	// block
	// get_local
	s0i32 = l13
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl89
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 69
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl89
	}
	// block
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// call
	s0i64 = f593(ctx, s0i32, s1i32)
	// tee_local
	l22 = s0i64
	// const
	s1i64 = -9223372036854775808
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl90
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
		goto lbl75
	}
	// const
	s0i64 = 0
	// set_local
	l22 = s0i64
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
		goto lbl90
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl90:
	// get_local
	s0i32 = l13
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl76
	}
	// get_local
	s0i64 = l22
	// get_local
	s1i64 = l12
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l12 = s0i64
	// br
	goto lbl74
	// end_block
lbl89:
	// get_local
	s0i32 = l13
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 0
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl77
	}
	// end_block
lbl78:
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
		goto lbl77
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl77:
	// get_local
	s0i32 = l10
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl74
	}
	// end_block
lbl76:
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// const
	s0f64 = 0
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl75:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// const
	s0f64 = 0
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl74:
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl91
	}
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl91:
	// block
	// get_local
	s0i64 = l18
	// const
	s1i64 = 9
	// binary: i64.gt_s
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// get_local
	s0i64 = l12
	// get_local
	s1i64 = l18
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// unary: i32.eqz
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// get_local
	s1i32 = l0
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl92:
	// block
	// get_local
	s0i64 = l12
	// get_local
	s1i32 = l4
	// const
	s2i32 = -2
	// binary: i32.div_s
	s1i32 = i32DivS(s1i32, s2i32)
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.le_s
	if s0i64 <= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = 1.7976931348623157e+308
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 1.7976931348623157e+308
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl93:
	// block
	// get_local
	s0i64 = l12
	// get_local
	s1i32 = l4
	// const
	s2i32 = -106
	// binary: i32.add
	s1i32 = s1i32 + s2i32
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
		goto lbl94
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = 2.2250738585072014e-308
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 2.2250738585072014e-308
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl94:
	// block
	// get_local
	s0i32 = l20
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl95
	}
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = 8
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl96
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l19
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l10 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0i32
	// block
	// block
	// const
	s0i32 = 1
	// get_local
	s1i32 = l20
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l8 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl98
	}
	// get_local
	s0i32 = l20
	// set_local
	l1 = s0i32
	// br
	goto lbl97
	// end_block
lbl98:
	// get_local
	s0i32 = l20
	// set_local
	l1 = s0i32
	// loop
lbl99:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 10
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l8
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l8 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl99
	}
	// end
	// end_block
lbl97:
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 7
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// loop
lbl101:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 100000000
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl101
	}
	// end
	// end_block
lbl100:
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl96:
	// get_local
	s0i32 = l19
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// end_block
lbl95:
	// get_local
	s0i64 = l12
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l16 = s0i32
	// block
	// get_local
	s0i32 = l17
	// const
	s1i32 = 9
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l16
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// get_local
	s0i32 = l16
	// const
	s1i32 = 17
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// block
	// get_local
	s0i32 = l16
	// const
	s1i32 = 9
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl103
	}
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl103:
	// block
	// get_local
	s0i32 = l16
	// const
	s1i32 = 8
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl104
	}
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1i32 = 8
	// get_local
	s2i32 = l16
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 23392
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl104:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0i32
	// block
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l16
	// const
	s2i32 = -3
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 27
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 30
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl105
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// end_block
lbl105:
	// get_local
	s0i32 = l9
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// get_local
	s1i32 = l0
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1i32 = l16
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 23352
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl0
	// end_block
lbl102:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l19
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// loop
lbl106:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l11 = s0i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// end
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// block
	// block
	// get_local
	s0i32 = l16
	// const
	s1i32 = 9
	// binary: i32.rem_s
	s0i32 = i32RemS(s0i32, s1i32)
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl108
	}
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// br
	goto lbl107
	// end_block
lbl108:
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l16
	// const
	s3i32 = 0
	// binary: i32.lt_s
	if s2i32 < s3i32 {
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
	l21 = s0i32
	// block
	// block
	// get_local
	s0i32 = l1
	// br_if
	if s0i32 != 0 {
		goto lbl110
	}
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// br
	goto lbl109
	// end_block
lbl110:
	// const
	s0i32 = 1000000000
	// const
	s1i32 = 8
	// get_local
	s2i32 = l21
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 23392
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// binary: i32.div_s
	s0i32 = i32DivS(s0i32, s1i32)
	// set_local
	l17 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l3
	// set_local
	l0 = s0i32
	// const
	s0i32 = 0
	// set_local
	l10 = s0i32
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// loop
lbl111:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l20 = s1i32
	// get_local
	s2i32 = l6
	// binary: i32.div_u
	s1i32 = i32DivU(s1i32, s2i32)
	// tee_local
	l2 = s1i32
	// get_local
	s2i32 = l13
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l13 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l10
	// get_local
	s3i32 = l8
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// get_local
	s3i32 = l13
	// unary: i32.eqz
	if s3i32 == 0 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l13 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l16
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l16
	// get_local
	s2i32 = l13
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l6
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l17
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l10
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l10 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl111
	}
	// end
	// get_local
	s0i32 = l13
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl109
	}
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l13
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// end_block
lbl109:
	// get_local
	s0i32 = l16
	// get_local
	s1i32 = l21
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l16 = s0i32
	// end_block
lbl107:
	// loop
lbl112:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// block
	// loop
lbl114:
	// block
	// get_local
	s0i32 = l16
	// const
	s1i32 = 18
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl115
	}
	// get_local
	s0i32 = l16
	// const
	s1i32 = 18
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl113
	}
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// const
	s1i32 = 9007198
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl113
	}
	// end_block
lbl115:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 127
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// const
	s0i32 = 0
	// set_local
	l10 = s0i32
	// loop
lbl116:
	// block
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l11
	// const
	s2i32 = 127
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l0 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l11 = s0i32
	// load: i64.load32_u
	s0i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// const
	s1i64 = 29
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// get_local
	s1i32 = l10
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l12 = s0i64
	// const
	s1i64 = 1000000001
	// binary: i64.ge_u
	if uint64(s0i64) >= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl118
	}
	// const
	s0i32 = 0
	// set_local
	l10 = s0i32
	// br
	goto lbl117
	// end_block
lbl118:
	// get_local
	s0i64 = l12
	// get_local
	s1i64 = l12
	// const
	s2i64 = 1000000000
	// binary: i64.div_u
	s1i64 = i64DivU(s1i64, s2i64)
	// tee_local
	l18 = s1i64
	// const
	s2i64 = 1000000000
	// binary: i64.mul
	s1i64 = s1i64 * s2i64
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// set_local
	l12 = s0i64
	// get_local
	s0i64 = l18
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l10 = s0i32
	// end_block
lbl117:
	// get_local
	s0i32 = l11
	// get_local
	s1i64 = l12
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// tee_local
	l13 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l1
	// get_local
	s3i32 = l0
	// get_local
	s4i32 = l13
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// get_local
	s3i32 = l0
	// get_local
	s4i32 = l8
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l1
	// const
	s4i32 = -1
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 127
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// binary: i32.ne
	if s2i32 != s3i32 {
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
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l8
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl116
	}
	// end
	// get_local
	s0i32 = l19
	// const
	s1i32 = -29
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l10
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl114
	}
	// end
	// block
	// get_local
	s0i32 = l8
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl119
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// const
	s2i32 = 126
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 127
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l3
	// get_local
	s3i32 = l1
	// const
	s4i32 = -1
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 127
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// tee_local
	l0 = s3i32
	// const
	s4i32 = 2
	// binary: i32.shl
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+0):]))
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// set_local
	l1 = s0i32
	// end_block
lbl119:
	// get_local
	s0i32 = l16
	// const
	s1i32 = 9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl112
	// end_block
lbl113:
	// end
	// block
	// loop
lbl121:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 127
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl122:
	// const
	s0i32 = 9
	// const
	s1i32 = 1
	// get_local
	s2i32 = l16
	// const
	s3i32 = 27
	// binary: i32.gt_s
	if s2i32 > s3i32 {
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
	l10 = s0i32
	// block
	// loop
lbl124:
	// block
	// block
	// get_local
	s0i32 = l8
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl126
	}
	// get_local
	s0i32 = l16
	// const
	s1i32 = 18
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl123
	}
	// br
	goto lbl125
	// end_block
lbl126:
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 9007199
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl127
	}
	// get_local
	s0i32 = l8
	// const
	s1i32 = 9007199
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl125
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l11 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl127
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l11
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// const
	s1i32 = 254740991
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl125
	}
	// get_local
	s0i32 = l16
	// const
	s1i32 = 18
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl125
	}
	// const
	s0i32 = 9007199
	// set_local
	l8 = s0i32
	// br
	goto lbl120
	// end_block
lbl127:
	// get_local
	s0i32 = l16
	// const
	s1i32 = 18
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// end_block
lbl125:
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l10
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l1
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl124
	}
	// end
	// const
	s0i32 = 1000000000
	// get_local
	s1i32 = l10
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l6 = s0i32
	// const
	s0i32 = -1
	// get_local
	s1i32 = l10
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l20 = s0i32
	// const
	s0i32 = 0
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l8 = s0i32
	// loop
lbl128:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l13 = s0i32
	// get_local
	s1i32 = l13
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l13 = s1i32
	// get_local
	s2i32 = l10
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// get_local
	s2i32 = l11
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l11 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l8
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// get_local
	s3i32 = l11
	// unary: i32.eqz
	if s3i32 == 0 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l11 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l16
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l16
	// get_local
	s2i32 = l11
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l13
	// get_local
	s1i32 = l20
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l6
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl128
	}
	// end
	// get_local
	s0i32 = l11
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl122
	}
	// block
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l8
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl129
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l17
	// set_local
	l1 = s0i32
	// br
	goto lbl121
	// end_block
lbl129:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl122
	// end_block
lbl123:
	// end
	// end
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l10 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l10
	// set_local
	l1 = s0i32
	// end_block
lbl120:
	// get_local
	s0i32 = l8
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l7 = s0f64
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl130
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl130:
	// get_local
	s0f64 = l7
	// const
	s1f64 = 1e+09
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l8
	// const
	s3i32 = 2
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l9
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// tee_local
	l23 = s1f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l24 = s0f64
	// const
	s0f64 = 0
	// set_local
	l7 = s0f64
	// block
	// block
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l4
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 53
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l8
	// const
	s3i32 = 0
	// binary: i32.gt_s
	if s2i32 > s3i32 {
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
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l8
	// get_local
	s3i32 = l5
	// binary: i32.lt_s
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l11 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 52
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl132
	}
	// const
	s0f64 = 0
	// set_local
	l25 = s0f64
	// br
	goto lbl131
	// end_block
lbl132:
	// const
	s0f64 = 1
	// const
	s1i32 = 105
	// get_local
	s2i32 = l8
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// get_local
	s1f64 = l24
	// binary: f64.copysign
	s0f64 = math.Copysign(s0f64, s1f64)
	// tee_local
	l25 = s0f64
	// get_local
	s1f64 = l24
	// get_local
	s2f64 = l24
	// const
	s3f64 = 1
	// const
	s4i32 = 53
	// get_local
	s5i32 = l8
	// binary: i32.sub
	s4i32 = s4i32 - s5i32
	// call
	s3f64 = f627(ctx, s3f64, s4i32)
	// call
	s2f64 = f616(ctx, s2f64, s3f64)
	// tee_local
	l7 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l24 = s0f64
	// end_block
lbl131:
	// get_local
	s0i32 = l19
	// const
	s1i32 = 53
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l10 = s0i32
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l13 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl133
	}
	// block
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l13
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l13 = s0i32
	// const
	s1i32 = 499999999
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl135
	}
	// block
	// get_local
	s0i32 = l13
	// br_if
	if s0i32 != 0 {
		goto lbl136
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl134
	}
	// end_block
lbl136:
	// get_local
	s0f64 = l23
	// const
	s1f64 = 0.25
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l7
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl134
	// end_block
lbl135:
	// block
	// get_local
	s0i32 = l13
	// const
	s1i32 = 500000000
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl137
	}
	// get_local
	s0f64 = l23
	// const
	s1f64 = 0.75
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l7
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl134
	// end_block
lbl137:
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl138
	}
	// get_local
	s0f64 = l23
	// const
	s1f64 = 0.5
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l7
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l7 = s0f64
	// br
	goto lbl134
	// end_block
lbl138:
	// get_local
	s0f64 = l23
	// const
	s1f64 = 0.75
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l7
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l7 = s0f64
	// end_block
lbl134:
	// get_local
	s0f64 = l7
	// get_local
	s1f64 = l7
	// get_local
	s2f64 = l7
	// const
	s3f64 = 1
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// get_local
	s3f64 = l7
	// const
	s4f64 = 1
	// call
	s3f64 = f616(ctx, s3f64, s4f64)
	// const
	s4f64 = 0
	// binary: f64.ne
	if s3f64 != s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// get_local
	s2i32 = l8
	// const
	s3i32 = 51
	// binary: i32.gt_s
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// set_local
	l7 = s0f64
	// end_block
lbl133:
	// get_local
	s0f64 = l24
	// get_local
	s1f64 = l7
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l25
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l24 = s0f64
	// block
	// block
	// get_local
	s0i32 = l10
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = -2
	// get_local
	s2i32 = l14
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl140
	}
	// get_local
	s0i32 = l19
	// set_local
	l0 = s0i32
	// br
	goto lbl139
	// end_block
lbl140:
	// get_local
	s0f64 = l24
	// const
	s1f64 = 0.5
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l24
	// get_local
	s2f64 = l24
	// unary: f64.abs
	s2f64 = math.Abs(s2f64)
	// const
	s3f64 = 9.007199254740992e+15
	// binary: f64.ge
	if s2f64 >= s3f64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l1 = s2i32
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// set_local
	l24 = s0f64
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l19
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 50
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l15
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl141
	}
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -53
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l19
	// binary: i32.ne
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l11
	// get_local
	s2i32 = l1
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// const
	s1i32 = 1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl139
	}
	// get_local
	s0f64 = l7
	// const
	s1f64 = 0
	// binary: f64.eq
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl139
	}
	// end_block
lbl141:
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// end_block
lbl139:
	// get_local
	s0f64 = l24
	// get_local
	s1i32 = l0
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// set_local
	l7 = s0f64
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 512
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0f64 = l7
	// return
	return s0f64
}
