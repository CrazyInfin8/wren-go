package internals

import (
	"encoding/binary"
	"math"
)

func f592(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
	_ = l9
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 float64
	_ = l18
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
	var s3i64 int64
	_ = s3i64
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
	var s5f64 float64
	_ = s5f64
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
		goto lbl1
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
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl0:
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// block
	// block
	// loop
lbl4:
	// block
	// get_local
	s0i32 = l5
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
		goto lbl5
	}
	// get_local
	s0i32 = l5
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
		goto lbl3
	}
	// const
	s0i64 = 0
	// set_local
	l7 = s0i64
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// br
	goto lbl2
	// end_block
lbl5:
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
		goto lbl6
	}
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
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
	goto lbl4
	// end_block
lbl6:
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l5 = s0i32
	// br
	goto lbl4
	// end
	// end_block
lbl3:
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
		goto lbl8
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
	goto lbl7
	// end_block
lbl8:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl7:
	// const
	s0i64 = 0
	// set_local
	l7 = s0i64
	// block
	// get_local
	s0i32 = l5
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
		goto lbl9
	}
	// const
	s0i32 = 1
	// set_local
	l8 = s0i32
	// br
	goto lbl2
	// end_block
lbl9:
	// loop
lbl10:
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
	goto lbl11
	// end_block
lbl12:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l5 = s0i32
	// end_block
lbl11:
	// get_local
	s0i64 = l7
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l7 = s0i64
	// get_local
	s0i32 = l5
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
		goto lbl10
	}
	// end
	// const
	s0i32 = 1
	// set_local
	l8 = s0i32
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// end_block
lbl2:
	// const
	s0i64 = 0
	// set_local
	l9 = s0i64
	// const
	s0f64 = 1
	// set_local
	l10 = s0f64
	// const
	s0f64 = 0
	// set_local
	l11 = s0f64
	// const
	s0i32 = 0
	// set_local
	l12 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// block
	// block
	// loop
lbl15:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l14 = s0i32
	// block
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l15 = s0i32
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
		goto lbl17
	}
	// block
	// get_local
	s0i32 = l14
	// const
	s1i32 = -97
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 6
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl18
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 46
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// end_block
lbl18:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 46
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
	// get_local
	s0i32 = l8
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// const
	s0i32 = 1
	// set_local
	l8 = s0i32
	// get_local
	s0i64 = l9
	// set_local
	l7 = s0i64
	// br
	goto lbl16
	// end_block
lbl17:
	// get_local
	s0i32 = l14
	// const
	s1i32 = -87
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l15
	// get_local
	s2i32 = l5
	// const
	s3i32 = 57
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
	l5 = s0i32
	// block
	// block
	// get_local
	s0i64 = l9
	// const
	s1i64 = 7
	// binary: i64.gt_s
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl20
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 4
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l12 = s0i32
	// br
	goto lbl19
	// end_block
lbl20:
	// block
	// get_local
	s0i64 = l9
	// const
	s1i64 = 13
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// get_local
	s0i32 = l5
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// get_local
	s1f64 = l10
	// const
	s2f64 = 0.0625
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l10 = s1f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l11
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l11 = s0f64
	// br
	goto lbl19
	// end_block
lbl21:
	// get_local
	s0f64 = l11
	// get_local
	s1f64 = l10
	// const
	s2f64 = 0.5
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l11
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2i32 = l5
	// unary: i32.eqz
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// get_local
	s3i32 = l13
	// const
	s4i32 = 0
	// binary: i32.ne
	if s3i32 != s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// tee_local
	l5 = s2i32
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// set_local
	l11 = s0f64
	// get_local
	s0i32 = l13
	// const
	s1i32 = 1
	// get_local
	s2i32 = l5
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l13 = s0i32
	// end_block
lbl19:
	// get_local
	s0i64 = l9
	// const
	s1i64 = 1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l9 = s0i64
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// end_block
lbl16:
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
		goto lbl22
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
lbl22:
	// get_local
	s0i32 = l0
	// call
	s0i32 = f590(ctx, s0i32)
	// set_local
	l5 = s0i32
	// br
	goto lbl15
	// end
	// end_block
lbl14:
	// const
	s0i32 = 46
	// set_local
	l5 = s0i32
	// end_block
lbl13:
	// block
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// block
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
		goto lbl26
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
	s0i32 = l4
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl25
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
	// get_local
	s0i32 = l8
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl24
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = -3
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl24
	// end_block
lbl26:
	// get_local
	s0i32 = l4
	// br_if
	if s0i32 != 0 {
		goto lbl24
	}
	// end_block
lbl25:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// end_block
lbl24:
	// get_local
	s0i32 = l3
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl23:
	// block
	// get_local
	s0i64 = l9
	// const
	s1i64 = 7
	// binary: i64.gt_s
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl27
	}
	// block
	// block
	// const
	s0i64 = 0
	// get_local
	s1i64 = l9
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// const
	s1i64 = 7
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// tee_local
	l16 = s0i64
	// unary: i64.eqz
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl29
	}
	// get_local
	s0i64 = l9
	// set_local
	l17 = s0i64
	// br
	goto lbl28
	// end_block
lbl29:
	// get_local
	s0i64 = l9
	// set_local
	l17 = s0i64
	// loop
lbl30:
	// get_local
	s0i64 = l17
	// const
	s1i64 = 1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l17 = s0i64
	// get_local
	s0i32 = l12
	// const
	s1i32 = 4
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l12 = s0i32
	// get_local
	s0i64 = l16
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l16 = s0i64
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
		goto lbl30
	}
	// end
	// end_block
lbl28:
	// get_local
	s0i64 = l9
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = 7
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl27
	}
	// get_local
	s0i64 = l17
	// const
	s1i64 = -8
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l17 = s0i64
	// loop
lbl31:
	// get_local
	s0i64 = l17
	// const
	s1i64 = 8
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l17 = s0i64
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
		goto lbl31
	}
	// end
	// const
	s0i32 = 0
	// set_local
	l12 = s0i32
	// end_block
lbl27:
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 80
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	s1i32 = l4
	// call
	s0i64 = f593(ctx, s0i32, s1i32)
	// tee_local
	l17 = s0i64
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
		goto lbl32
	}
	// block
	// get_local
	s0i32 = l4
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// get_local
	s0i32 = l0
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+88):]))
	// const
	s1i64 = -1
	// binary: i64.gt_s
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// br
	goto lbl33
	// end_block
lbl36:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// call
	f589(ctx, s0i32, s1i64)
	// const
	s0f64 = 0
	// return
	return s0f64
	// end_block
lbl35:
	// const
	s0i64 = 0
	// set_local
	l17 = s0i64
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
		goto lbl32
	}
	// end_block
lbl34:
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
lbl33:
	// const
	s0i64 = 0
	// set_local
	l17 = s0i64
	// end_block
lbl32:
	// block
	// get_local
	s0i32 = l12
	// br_if
	if s0i32 != 0 {
		goto lbl37
	}
	// get_local
	s0i32 = l3
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl37:
	// block
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l9
	// get_local
	s2i32 = l8
	// select
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	// const
	s1i64 = 2
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// get_local
	s1i64 = l17
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = -32
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l9 = s0i64
	// const
	s1i32 = 0
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
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
		goto lbl38
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l3
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
	// return
	return s0f64
	// end_block
lbl38:
	// block
	// get_local
	s0i64 = l9
	// get_local
	s1i32 = l2
	// const
	s2i32 = -106
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl39
	}
	// block
	// get_local
	s0i32 = l12
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
		goto lbl40
	}
	// loop
lbl41:
	// get_local
	s0f64 = l11
	// get_local
	s1f64 = l11
	// const
	s2f64 = -1
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l11
	// get_local
	s3f64 = l11
	// const
	s4f64 = 0.5
	// binary: f64.ge
	if s3f64 >= s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// tee_local
	l5 = s3i32
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l11 = s0f64
	// get_local
	s0i64 = l9
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l9 = s0i64
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// tee_local
	l12 = s0i32
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
		goto lbl41
	}
	// end
	// end_block
lbl40:
	// block
	// block
	// get_local
	s0i64 = l9
	// get_local
	s1i32 = l2
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// const
	s1i64 = 32
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l7 = s0i64
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l5
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
	s1i32 = l1
	// get_local
	s2i64 = l7
	// get_local
	s3i32 = l1
	// unary: i64.extend_u/i32
	s3i64 = int64(uint32(s3i32))
	// binary: i64.lt_s
	if s2i64 < s3i64 {
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
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 53
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl43
	}
	// get_local
	s0i32 = l3
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l10 = s0f64
	// const
	s0f64 = 0
	// set_local
	l18 = s0f64
	// br
	goto lbl42
	// end_block
lbl43:
	// const
	s0f64 = 1
	// const
	s1i32 = 84
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// get_local
	s1i32 = l3
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// tee_local
	l10 = s1f64
	// binary: f64.copysign
	s0f64 = math.Copysign(s0f64, s1f64)
	// set_local
	l18 = s0f64
	// end_block
lbl42:
	// block
	// get_local
	s0f64 = l10
	// const
	s1f64 = 0
	// get_local
	s2f64 = l11
	// get_local
	s3i32 = l5
	// const
	s4i32 = 32
	// binary: i32.lt_s
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// get_local
	s4f64 = l11
	// const
	s5f64 = 0
	// binary: f64.ne
	if s4f64 != s5f64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// get_local
	s4i32 = l12
	// const
	s5i32 = 1
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// unary: i32.eqz
	if s4i32 == 0 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// tee_local
	l5 = s3i32
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l10
	// get_local
	s2i32 = l12
	// get_local
	s3i32 = l5
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// unary: f64.convert_u/i32
	s2f64 = float64(uint32(s2i32))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l18
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l18
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l11 = s0f64
	// const
	s1f64 = 0
	// binary: f64.ne
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl44
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// end_block
lbl44:
	// get_local
	s0f64 = l11
	// get_local
	s1i64 = l9
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// return
	return s0f64
	// end_block
lbl39:
	// const
	s0i32 = 0
	// const
	s1i32 = 68
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// get_local
	s0i32 = l3
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
	// return
	return s0f64
}
