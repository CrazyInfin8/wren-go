package internals

import (
	"encoding/binary"
	"math"
)

func f569(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
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
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 int32
	_ = l30
	var l31 int32
	_ = l31
	var l32 int32
	_ = l32
	var l33 int64
	_ = l33
	var l34 int64
	_ = l34
	var l35 float64
	_ = l35
	var l36 int32
	_ = l36
	var l37 int32
	_ = l37
	var l38 int32
	_ = l38
	var l39 int32
	_ = l39
	var l40 int32
	_ = l40
	var l41 int32
	_ = l41
	var l42 float64
	_ = l42
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
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 880
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 68
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 55
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// const
	s0i32 = -2
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 68
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 11
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 8
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 9
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l11 = s0i32
	// const
	s0i32 = -10
	// get_local
	s1i32 = l5
	// const
	s2i32 = 68
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 68
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 10
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 404
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 4
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 400
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 56
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l17 = s0i32
	// const
	s0i32 = 0
	// set_local
	l18 = s0i32
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// block
	// block
	// block
	// loop
lbl3:
	// get_local
	s0i32 = l1
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 2147483647
	// get_local
	s2i32 = l19
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
		goto lbl2
	}
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l19
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l21
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// get_local
	s0i32 = l21
	// set_local
	l1 = s0i32
	// loop
lbl9:
	// block
	// block
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = 37
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
	// get_local
	s0i32 = l1
	// set_local
	l20 = s0i32
	// loop
lbl13:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 37
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
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 37
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// br
	goto lbl11
	// end
	// end_block
lbl12:
	// get_local
	s0i32 = l1
	// set_local
	l20 = s0i32
	// end_block
lbl11:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 2147483647
	// get_local
	s2i32 = l19
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l22 = s1i32
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// block
	// get_local
	s0i32 = l0
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl14:
	// get_local
	s0i32 = l20
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// const
	s0i32 = -1
	// set_local
	l23 = s0i32
	// block
	// get_local
	s0i32 = l1
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// tee_local
	l24 = s0i32
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l25 = s0i32
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
		goto lbl15
	}
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	// const
	s1i32 = 36
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl15
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+3)]))
	// set_local
	l24 = s0i32
	// const
	s0i32 = 1
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l23 = s0i32
	// end_block
lbl15:
	// const
	s0i32 = 0
	// set_local
	l26 = s0i32
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 31
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
	// const
	s0i32 = 1
	// get_local
	s1i32 = l1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 75913
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
		goto lbl16
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// const
	s0i32 = 0
	// set_local
	l26 = s0i32
	// loop
lbl17:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l26
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// tee_local
	l20 = s0i32
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	// tee_local
	l24 = s0i32
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl16
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// const
	s0i32 = 1
	// get_local
	s1i32 = l1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 75913
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl17
	}
	// end
	// end_block
lbl16:
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = 42
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl18
	}
	// block
	// block
	// get_local
	s0i32 = l20
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl20
	}
	// get_local
	s0i32 = l20
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	// const
	s1i32 = 36
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
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l20
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l20
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -384
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l27 = s0i32
	// const
	s0i32 = 1
	// set_local
	l18 = s0i32
	// br
	goto lbl19
	// end_block
lbl20:
	// get_local
	s0i32 = l18
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// const
	s0i32 = 0
	// set_local
	l18 = s0i32
	// const
	s0i32 = 0
	// set_local
	l27 = s0i32
	// br
	goto lbl7
	// end_block
lbl21:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l27 = s0i32
	// const
	s0i32 = 0
	// set_local
	l18 = s0i32
	// end_block
lbl19:
	// get_local
	s0i32 = l27
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
		goto lbl7
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l27
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l26
	// const
	s1i32 = 8192
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l26 = s0i32
	// br
	goto lbl7
	// end_block
lbl18:
	// const
	s0i32 = 0
	// set_local
	l27 = s0i32
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl22
	}
	// get_local
	s0i32 = l20
	// set_local
	l25 = s0i32
	// br
	goto lbl7
	// end_block
lbl22:
	// const
	s0i32 = 0
	// set_local
	l27 = s0i32
	// loop
lbl23:
	// block
	// get_local
	s0i32 = l27
	// const
	s1i32 = 214748364
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl24
	}
	// const
	s0i32 = -1
	// get_local
	s1i32 = l27
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// tee_local
	l25 = s1i32
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l1
	// const
	s3i32 = 2147483647
	// get_local
	s4i32 = l25
	// binary: i32.sub
	s3i32 = s3i32 - s4i32
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l27 = s0i32
	// get_local
	s0i32 = l20
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l25 = s0i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl23
	}
	// get_local
	s0i32 = l27
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
		goto lbl2
	}
	// br
	goto lbl7
	// end_block
lbl24:
	// get_local
	s0i32 = l20
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// set_local
	l1 = s0i32
	// const
	s0i32 = -1
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl23
	}
	// br
	goto lbl2
	// end
	// end_block
lbl10:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l20 = s0i32
	// br
	goto lbl9
	// end
	// end_block
lbl8:
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// block
	// get_local
	s0i32 = l18
	// br_if
	if s0i32 != 0 {
		goto lbl25
	}
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// br
	goto lbl0
	// end_block
lbl25:
	// block
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl27
	}
	// const
	s0i32 = 1
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl27:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl28
	}
	// const
	s0i32 = 2
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl28:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl29
	}
	// const
	s0i32 = 3
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl29:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl30
	}
	// const
	s0i32 = 4
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl30:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl31
	}
	// const
	s0i32 = 5
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl31:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 40
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl32
	}
	// const
	s0i32 = 6
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl32:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl33
	}
	// const
	s0i32 = 7
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl33:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 56
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// const
	s0i32 = 8
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl34:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 64
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl35
	}
	// const
	s0i32 = 9
	// set_local
	l1 = s0i32
	// br
	goto lbl26
	// end_block
lbl35:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 72
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// const
	s0i32 = 1
	// set_local
	l19 = s0i32
	// br
	goto lbl0
	// end_block
lbl26:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l1 = s0i32
	// loop
lbl36:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl36
	}
	// end
	// const
	s0i32 = 1
	// set_local
	l19 = s0i32
	// br
	goto lbl0
	// end_block
lbl7:
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// const
	s0i32 = -1
	// set_local
	l24 = s0i32
	// block
	// block
	// get_local
	s0i32 = l25
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
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
		goto lbl38
	}
	// get_local
	s0i32 = l25
	// set_local
	l1 = s0i32
	// const
	s0i32 = 0
	// set_local
	l28 = s0i32
	// br
	goto lbl37
	// end_block
lbl38:
	// block
	// get_local
	s0i32 = l25
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
	// tee_local
	l24 = s0i32
	// const
	s1i32 = 42
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl39
	}
	// block
	// block
	// get_local
	s0i32 = l25
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+2)]))
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
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
		goto lbl41
	}
	// get_local
	s0i32 = l25
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	// const
	s1i32 = 36
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
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l25
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l25
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+2)]))
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -384
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l24 = s0i32
	// br
	goto lbl40
	// end_block
lbl41:
	// get_local
	s0i32 = l18
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l25
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl42
	}
	// const
	s0i32 = 0
	// set_local
	l24 = s0i32
	// br
	goto lbl40
	// end_block
lbl42:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l25 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l25
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l24 = s0i32
	// end_block
lbl40:
	// get_local
	s0i32 = l24
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// const
	s1i32 = 31
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l28 = s0i32
	// br
	goto lbl37
	// end_block
lbl39:
	// get_local
	s0i32 = l25
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l29 = s0i32
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
		goto lbl43
	}
	// const
	s0i32 = 1
	// set_local
	l28 = s0i32
	// const
	s0i32 = 0
	// set_local
	l24 = s0i32
	// br
	goto lbl37
	// end_block
lbl43:
	// const
	s0i32 = 0
	// set_local
	l25 = s0i32
	// loop
lbl44:
	// const
	s0i32 = -1
	// set_local
	l24 = s0i32
	// block
	// get_local
	s0i32 = l25
	// const
	s1i32 = 214748364
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// const
	s0i32 = -1
	// get_local
	s1i32 = l25
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// tee_local
	l25 = s1i32
	// get_local
	s2i32 = l29
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l29
	// const
	s3i32 = 2147483647
	// get_local
	s4i32 = l25
	// binary: i32.sub
	s3i32 = s3i32 - s4i32
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l24 = s0i32
	// end_block
lbl45:
	// const
	s0i32 = 1
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l24
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	// const
	s1i32 = -48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l29 = s0i32
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
		goto lbl44
	}
	// end
	// end_block
lbl37:
	// loop
lbl46:
	// get_local
	s0i32 = l20
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	// tee_local
	l20 = s0i32
	// const
	s1i32 = -123
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -58
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
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l25
	// const
	s2i32 = 58
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 22847
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// tee_local
	l20 = s0i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 8
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
	// end
	// block
	// block
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = 27
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl49
	}
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
		goto lbl6
	}
	// block
	// get_local
	s0i32 = l23
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
		goto lbl50
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l23
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l23
	// const
	s3i32 = 3
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// br
	goto lbl48
	// end_block
lbl50:
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl51
	}
	// const
	s0i32 = 0
	// set_local
	l19 = s0i32
	// br
	goto lbl0
	// end_block
lbl51:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 56
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l2
	// call
	f570(ctx, s0i32, s1i32, s2i32)
	// br
	goto lbl47
	// end_block
lbl49:
	// get_local
	s0i32 = l23
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
		goto lbl6
	}
	// end_block
lbl48:
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l0
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// end_block
lbl47:
	// get_local
	s0i32 = l26
	// const
	s1i32 = -65537
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l29 = s0i32
	// get_local
	s1i32 = l26
	// get_local
	s2i32 = l26
	// const
	s3i32 = 8192
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l30 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
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
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_s
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	// tee_local
	l20 = s0i32
	// const
	s1i32 = -33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 15
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 3
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
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l25
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l31 = s0i32
	// const
	s1i32 = -65
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl54
	case 1:
		goto lbl53
	case 2:
		goto lbl57
	case 3:
		goto lbl53
	case 4:
		goto lbl54
	case 5:
		goto lbl54
	case 6:
		goto lbl54
	case 7:
		goto lbl53
	case 8:
		goto lbl53
	case 9:
		goto lbl53
	case 10:
		goto lbl53
	case 11:
		goto lbl53
	case 12:
		goto lbl53
	case 13:
		goto lbl53
	case 14:
		goto lbl53
	case 15:
		goto lbl53
	case 16:
		goto lbl53
	case 17:
		goto lbl53
	case 18:
		goto lbl58
	case 19:
		goto lbl53
	case 20:
		goto lbl53
	case 21:
		goto lbl53
	case 22:
		goto lbl53
	case 23:
		goto lbl67
	case 24:
		goto lbl53
	case 25:
		goto lbl53
	case 26:
		goto lbl53
	case 27:
		goto lbl53
	case 28:
		goto lbl53
	case 29:
		goto lbl53
	case 30:
		goto lbl53
	case 31:
		goto lbl53
	case 32:
		goto lbl54
	case 33:
		goto lbl53
	case 34:
		goto lbl62
	case 35:
		goto lbl65
	case 36:
		goto lbl54
	case 37:
		goto lbl54
	case 38:
		goto lbl54
	case 39:
		goto lbl53
	case 40:
		goto lbl65
	case 41:
		goto lbl53
	case 42:
		goto lbl53
	case 43:
		goto lbl53
	case 44:
		goto lbl61
	case 45:
		goto lbl69
	case 46:
		goto lbl66
	case 47:
		goto lbl68
	case 48:
		goto lbl53
	case 49:
		goto lbl53
	case 50:
		goto lbl60
	case 51:
		goto lbl53
	case 52:
		goto lbl70
	case 53:
		goto lbl53
	case 54:
		goto lbl53
	case 55:
		goto lbl67
	default:
		goto lbl53
	}
	// end_block
lbl70:
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l33 = s0i64
	// br
	goto lbl64
	// end_block
lbl69:
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l25
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl77
	case 1:
		goto lbl76
	case 2:
		goto lbl75
	case 3:
		goto lbl74
	case 4:
		goto lbl73
	case 5:
		goto lbl3
	case 6:
		goto lbl72
	case 7:
		goto lbl71
	default:
		goto lbl3
	}
	// end_block
lbl77:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl76:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl75:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// br
	goto lbl3
	// end_block
lbl74:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// store: i32.store16
	binary.LittleEndian.PutUint16(ctx.Mem[int(s0i32+0):], uint16(s1i32))
	// br
	goto lbl3
	// end_block
lbl73:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// br
	goto lbl3
	// end_block
lbl72:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl71:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l19
	// unary: i64.extend_s/i32
	s1i64 = int64(s1i32)
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// br
	goto lbl3
	// end_block
lbl68:
	// get_local
	s0i32 = l24
	// const
	s1i32 = 8
	// get_local
	s2i32 = l24
	// const
	s3i32 = 8
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l24 = s0i32
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l30 = s0i32
	// const
	s0i32 = 120
	// set_local
	l31 = s0i32
	// end_block
lbl67:
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// block
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// tee_local
	l33 = s0i64
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
		goto lbl78
	}
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// br
	goto lbl63
	// end_block
lbl78:
	// get_local
	s0i32 = l31
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// loop
lbl79:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i64 = l33
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 23376
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// get_local
	s2i32 = l25
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i64 = l33
	// const
	s1i64 = 15
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l20 = s0i32
	// get_local
	s0i64 = l33
	// const
	s1i64 = 4
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l33 = s0i64
	// get_local
	s0i32 = l20
	// br_if
	if s0i32 != 0 {
		goto lbl79
	}
	// end
	// get_local
	s0i32 = l30
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
		goto lbl63
	}
	// get_local
	s0i32 = l31
	// const
	s1i32 = 4
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// const
	s1i32 = 1091
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l32 = s0i32
	// const
	s0i32 = 2
	// set_local
	l23 = s0i32
	// br
	goto lbl63
	// end_block
lbl66:
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// tee_local
	l33 = s0i64
	// unary: i64.eqz
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl80
	}
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// loop
lbl81:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i64 = l33
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// const
	s2i32 = 7
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i64 = l33
	// const
	s1i64 = 7
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l20 = s0i32
	// get_local
	s0i64 = l33
	// const
	s1i64 = 3
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l33 = s0i64
	// get_local
	s0i32 = l20
	// br_if
	if s0i32 != 0 {
		goto lbl81
	}
	// end
	// end_block
lbl80:
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l30
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
		goto lbl63
	}
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l17
	// get_local
	s2i32 = l21
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l20 = s1i32
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l24
	// get_local
	s3i32 = l20
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
	l24 = s0i32
	// br
	goto lbl63
	// end_block
lbl65:
	// block
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// tee_local
	l33 = s0i64
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
		goto lbl82
	}
	// get_local
	s0i32 = l5
	// const
	s1i64 = 0
	// get_local
	s2i64 = l33
	// binary: i64.sub
	s1i64 = s1i64 - s2i64
	// tee_local
	l33 = s1i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// const
	s0i32 = 1
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// br
	goto lbl64
	// end_block
lbl82:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 2048
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
	// const
	s0i32 = 1
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1092
	// set_local
	l32 = s0i32
	// br
	goto lbl64
	// end_block
lbl83:
	// const
	s0i32 = 1093
	// const
	s1i32 = 1091
	// get_local
	s2i32 = l30
	// const
	s3i32 = 1
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l23 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l32 = s0i32
	// end_block
lbl64:
	// block
	// block
	// get_local
	s0i64 = l33
	// const
	s1i64 = 4294967296
	// binary: i64.ge_u
	if uint64(s0i64) >= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl85
	}
	// get_local
	s0i64 = l33
	// set_local
	l34 = s0i64
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// br
	goto lbl84
	// end_block
lbl85:
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// loop
lbl86:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i64 = l33
	// get_local
	s2i64 = l33
	// const
	s3i64 = 10
	// binary: i64.div_u
	s2i64 = i64DivU(s2i64, s3i64)
	// tee_local
	l34 = s2i64
	// const
	s3i64 = 10
	// binary: i64.mul
	s2i64 = s2i64 * s3i64
	// binary: i64.sub
	s1i64 = s1i64 - s2i64
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i64 = l33
	// const
	s1i64 = 42949672959
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l20 = s0i32
	// get_local
	s0i64 = l34
	// set_local
	l33 = s0i64
	// get_local
	s0i32 = l20
	// br_if
	if s0i32 != 0 {
		goto lbl86
	}
	// end
	// end_block
lbl84:
	// get_local
	s0i64 = l34
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl63
	}
	// loop
lbl87:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l25 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl87
	}
	// end
	// end_block
lbl63:
	// block
	// get_local
	s0i32 = l28
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl88
	}
	// get_local
	s0i32 = l24
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
		goto lbl2
	}
	// end_block
lbl88:
	// get_local
	s0i32 = l30
	// const
	s1i32 = -65537
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l30
	// get_local
	s2i32 = l28
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l29 = s0i32
	// block
	// get_local
	s0i32 = l5
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// tee_local
	l33 = s0i64
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
		goto lbl89
	}
	// const
	s0i32 = 0
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l24
	// br_if
	if s0i32 != 0 {
		goto lbl89
	}
	// get_local
	s0i32 = l17
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l20 = s0i32
	// br
	goto lbl52
	// end_block
lbl89:
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l17
	// get_local
	s2i32 = l21
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// get_local
	s2i64 = l33
	// unary: i64.eqz
	if s2i64 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l20 = s1i32
	// get_local
	s2i32 = l24
	// get_local
	s3i32 = l20
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
	l26 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l20 = s0i32
	// br
	goto lbl52
	// end_block
lbl62:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+56):]))
	// store: i64.store8
	ctx.Mem[int(s0i32+55)] = uint8(s1i64)
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// const
	s0i32 = 1
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l7
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l20 = s0i32
	// br
	goto lbl52
	// end_block
lbl61:
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42068):]))
	// call
	s0i32 = f566(ctx, s0i32)
	// set_local
	l21 = s0i32
	// br
	goto lbl59
	// end_block
lbl60:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 6248
	// get_local
	s2i32 = l20
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l21 = s0i32
	// end_block
lbl59:
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l21
	// const
	s2i32 = 2147483647
	// get_local
	s3i32 = l24
	// get_local
	s4i32 = l24
	// const
	s5i32 = 0
	// binary: i32.lt_s
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s1i32 = f583(ctx, s1i32, s2i32)
	// tee_local
	l26 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l24
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
		goto lbl52
	}
	// get_local
	s0i32 = l20
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
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
	// br
	goto lbl2
	// end_block
lbl58:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l24
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// br
	goto lbl55
	// end_block
lbl57:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+56):]))
	// store: i64.store32
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i64))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+56):], uint32(s1i32))
	// const
	s0i32 = -1
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// end_block
lbl56:
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l21 = s0i32
	// block
	// loop
lbl91:
	// get_local
	s0i32 = l21
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l22 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl90
	}
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// call
	s0i32 = f598(ctx, s0i32, s1i32)
	// tee_local
	l22 = s0i32
	// const
	s1i32 = 0
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l26 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// get_local
	s0i32 = l22
	// get_local
	s1i32 = l24
	// get_local
	s2i32 = l20
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// get_local
	s0i32 = l21
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l20
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l20 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl91
	}
	// br
	goto lbl90
	// end_block
lbl92:
	// end
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// end_block
lbl90:
	// get_local
	s0i32 = l20
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
		goto lbl2
	}
	// end_block
lbl55:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l26 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l20
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l20
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l21 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l21
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl94
	}
	// loop
lbl95:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl96
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl96:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl95
	}
	// end
	// end_block
lbl94:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl93:
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
		goto lbl97
	}
	// const
	s0i32 = 0
	// set_local
	l21 = s0i32
	// loop
lbl98:
	// get_local
	s0i32 = l25
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l22 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl97
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// call
	s0i32 = f598(ctx, s0i32, s1i32)
	// tee_local
	l22 = s0i32
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl97
	}
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl99
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl99:
	// get_local
	s0i32 = l25
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l20
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl98
	}
	// end
	// end_block
lbl97:
	// block
	// get_local
	s0i32 = l26
	// const
	s1i32 = 8192
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l20
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l20
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l21 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l21
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl101
	}
	// loop
lbl102:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl103
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl103:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// end
	// end_block
lbl101:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl100:
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l20
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
	l20 = s0i32
	// br
	goto lbl3
	// end_block
lbl54:
	// block
	// get_local
	s0i32 = l28
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl104
	}
	// get_local
	s0i32 = l24
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
		goto lbl2
	}
	// end_block
lbl104:
	// get_local
	s0i32 = l5
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l35 = s0f64
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+108):], uint32(s1i32))
	// block
	// block
	// get_local
	s0f64 = l35
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
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
		goto lbl106
	}
	// get_local
	s0f64 = l35
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l35 = s0f64
	// const
	s0i32 = 1
	// set_local
	l36 = s0i32
	// const
	s0i32 = 0
	// set_local
	l37 = s0i32
	// const
	s0i32 = 1101
	// set_local
	l38 = s0i32
	// br
	goto lbl105
	// end_block
lbl106:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 2048
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
		goto lbl107
	}
	// const
	s0i32 = 1
	// set_local
	l36 = s0i32
	// const
	s0i32 = 0
	// set_local
	l37 = s0i32
	// const
	s0i32 = 1104
	// set_local
	l38 = s0i32
	// br
	goto lbl105
	// end_block
lbl107:
	// const
	s0i32 = 1107
	// const
	s1i32 = 1102
	// get_local
	s2i32 = l30
	// const
	s3i32 = 1
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l36 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l36
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l37 = s0i32
	// end_block
lbl105:
	// block
	// get_local
	s0f64 = l35
	// unary: f64.abs
	s0f64 = math.Abs(s0f64)
	// const
	s1f64 = math.Inf(0)
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl108
	}
	// get_local
	s0i32 = l36
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8192
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl109
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl109
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl110
	}
	// loop
lbl111:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl112
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl112:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl111
	}
	// end
	// end_block
lbl110:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl109
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl109:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl113
	}
	// get_local
	s0i32 = l38
	// get_local
	s1i32 = l36
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l20 = s0i32
	// end_block
lbl113:
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl114
	}
	// const
	s0i32 = 1496
	// const
	s1i32 = 1939
	// get_local
	s2i32 = l31
	// const
	s3i32 = 32
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l20 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// const
	s1i32 = 1591
	// const
	s2i32 = 1943
	// get_local
	s3i32 = l20
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// get_local
	s2f64 = l35
	// get_local
	s3f64 = l35
	// binary: f64.ne
	if s2f64 != s3f64 {
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
	// const
	s1i32 = 3
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl114:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 8192
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl115
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl115
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl116
	}
	// loop
lbl117:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl118
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl118:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl117
	}
	// end
	// end_block
lbl116:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl115
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl115:
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
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
	l20 = s0i32
	// br
	goto lbl3
	// end_block
lbl108:
	// block
	// block
	// block
	// get_local
	s0f64 = l35
	// get_local
	s1i32 = l5
	// const
	s2i32 = 108
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// call
	s0f64 = f617(ctx, s0f64, s1i32)
	// tee_local
	l35 = s0f64
	// get_local
	s1f64 = l35
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l35 = s0f64
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
		goto lbl121
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+108):]))
	// tee_local
	l20 = s1i32
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+108):], uint32(s1i32))
	// get_local
	s0i32 = l31
	// const
	s1i32 = 32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// tee_local
	l39 = s0i32
	// const
	s1i32 = 97
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// br
	goto lbl5
	// end_block
lbl121:
	// get_local
	s0i32 = l31
	// const
	s1i32 = 32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// tee_local
	l39 = s0i32
	// const
	s1i32 = 97
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
	// const
	s0i32 = 6
	// get_local
	s1i32 = l24
	// get_local
	s2i32 = l24
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
	l23 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+108):]))
	// set_local
	l25 = s0i32
	// br
	goto lbl119
	// end_block
lbl120:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l20
	// const
	s2i32 = -29
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l25 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+108):], uint32(s1i32))
	// const
	s0i32 = 6
	// get_local
	s1i32 = l24
	// get_local
	s2i32 = l24
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
	l23 = s0i32
	// get_local
	s0f64 = l35
	// const
	s1f64 = 2.68435456e+08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l35 = s0f64
	// end_block
lbl119:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l16
	// get_local
	s2i32 = l25
	// const
	s3i32 = 0
	// binary: i32.lt_s
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l40 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l32 = s0i32
	// set_local
	l21 = s0i32
	// loop
lbl122:
	// block
	// block
	// get_local
	s0f64 = l35
	// const
	s1f64 = 4.294967296e+09
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s1f64 = l35
	// const
	s2f64 = 0
	// binary: f64.ge
	if s1f64 >= s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
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
		goto lbl124
	}
	// get_local
	s0f64 = l35
	// unary: i32.trunc_u/f64
	s0i32 = int32(uint32(math.Trunc(s0f64)))
	// set_local
	l20 = s0i32
	// br
	goto lbl123
	// end_block
lbl124:
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// end_block
lbl123:
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l20
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l21
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0f64 = l35
	// get_local
	s1i32 = l20
	// unary: f64.convert_u/i32
	s1f64 = float64(uint32(s1i32))
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 1e+09
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l35 = s0f64
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
		goto lbl122
	}
	// end
	// block
	// block
	// get_local
	s0i32 = l25
	// const
	s1i32 = 1
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl126
	}
	// get_local
	s0i32 = l21
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l32
	// set_local
	l22 = s0i32
	// br
	goto lbl125
	// end_block
lbl126:
	// get_local
	s0i32 = l32
	// set_local
	l22 = s0i32
	// loop
lbl127:
	// get_local
	s0i32 = l25
	// const
	s1i32 = 29
	// get_local
	s2i32 = l25
	// const
	s3i32 = 29
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
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
	l25 = s0i32
	// block
	// get_local
	s0i32 = l21
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// get_local
	s1i32 = l22
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl128
	}
	// get_local
	s0i32 = l25
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// set_local
	l34 = s0i64
	// const
	s0i64 = 0
	// set_local
	l33 = s0i64
	// loop
lbl129:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l20
	// load: i64.load32_u
	s1i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i64 = l34
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// get_local
	s2i64 = l33
	// const
	s3i64 = 4294967295
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// tee_local
	l33 = s1i64
	// get_local
	s2i64 = l33
	// const
	s3i64 = 1000000000
	// binary: i64.div_u
	s2i64 = i64DivU(s2i64, s3i64)
	// tee_local
	l33 = s2i64
	// const
	s3i64 = 1000000000
	// binary: i64.mul
	s2i64 = s2i64 * s3i64
	// binary: i64.sub
	s1i64 = s1i64 - s2i64
	// store: i64.store32
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i64))
	// get_local
	s0i32 = l20
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// get_local
	s1i32 = l22
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl129
	}
	// end
	// get_local
	s0i64 = l33
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl128
	}
	// get_local
	s0i32 = l22
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l22 = s0i32
	// get_local
	s1i32 = l20
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl128:
	// block
	// loop
lbl131:
	// get_local
	s0i32 = l21
	// tee_local
	l20 = s0i32
	// get_local
	s1i32 = l22
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl130
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
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
		goto lbl131
	}
	// end
	// end_block
lbl130:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+108):]))
	// get_local
	s2i32 = l25
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l25 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+108):], uint32(s1i32))
	// get_local
	s0i32 = l20
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l25
	// const
	s1i32 = 0
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl127
	}
	// end
	// end_block
lbl125:
	// block
	// get_local
	s0i32 = l25
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
		goto lbl132
	}
	// get_local
	s0i32 = l23
	// const
	s1i32 = 25
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 9
	// binary: i32.div_u
	s0i32 = i32DivU(s0i32, s1i32)
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l41 = s0i32
	// loop
lbl133:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l25
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 9
	// get_local
	s2i32 = l21
	// const
	s3i32 = 9
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
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
	l24 = s0i32
	// block
	// block
	// get_local
	s0i32 = l22
	// get_local
	s1i32 = l20
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl135
	}
	// get_local
	s0i32 = l22
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l21 = s0i32
	// br
	goto lbl134
	// end_block
lbl135:
	// const
	s0i32 = 1000000000
	// get_local
	s1i32 = l24
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l29 = s0i32
	// const
	s0i32 = -1
	// get_local
	s1i32 = l24
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l28 = s0i32
	// const
	s0i32 = 0
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l22
	// set_local
	l21 = s0i32
	// loop
lbl136:
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l21
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l26 = s1i32
	// get_local
	s2i32 = l24
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// get_local
	s2i32 = l25
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l28
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l29
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l21
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl136
	}
	// end
	// get_local
	s0i32 = l22
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l25
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl134
	}
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l25
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l20
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// end_block
lbl134:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l5
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+108):]))
	// get_local
	s2i32 = l24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l25 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+108):], uint32(s1i32))
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l21
	// unary: i32.eqz
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// const
	s3i32 = 2
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l22 = s1i32
	// get_local
	s2i32 = l39
	// const
	s3i32 = 102
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
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l41
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 2
	// binary: i32.shr_s
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	// get_local
	s3i32 = l41
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
	l20 = s0i32
	// get_local
	s0i32 = l25
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
		goto lbl133
	}
	// end
	// end_block
lbl132:
	// const
	s0i32 = 0
	// set_local
	l26 = s0i32
	// block
	// get_local
	s0i32 = l22
	// get_local
	s1i32 = l20
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl137
	}
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l22
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 2
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// const
	s1i32 = 9
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l22
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l25 = s0i32
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
		goto lbl137
	}
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// loop
lbl138:
	// get_local
	s0i32 = l26
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l21
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// tee_local
	l21 = s1i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl138
	}
	// end
	// end_block
lbl137:
	// block
	// get_local
	s0i32 = l23
	// const
	s1i32 = 0
	// get_local
	s2i32 = l26
	// get_local
	s3i32 = l39
	// const
	s4i32 = 102
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
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l23
	// const
	s2i32 = 0
	// binary: i32.ne
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// get_local
	s2i32 = l39
	// const
	s3i32 = 103
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l28 = s2i32
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 2
	// binary: i32.shr_s
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	// const
	s2i32 = 9
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// const
	s2i32 = -9
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl139
	}
	// get_local
	s0i32 = l21
	// const
	s1i32 = 9216
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l25 = s0i32
	// const
	s1i32 = 9
	// binary: i32.div_s
	s0i32 = i32DivS(s0i32, s1i32)
	// tee_local
	l24 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l15
	// get_local
	s2i32 = l14
	// get_local
	s3i32 = l40
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l40 = s0i32
	// const
	s1i32 = -4096
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l29 = s0i32
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l24
	// const
	s2i32 = 9
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l24 = s0i32
	// const
	s1i32 = 7
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
	// const
	s0i32 = 8
	// get_local
	s1i32 = l24
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l41 = s0i32
	// const
	s1i32 = 7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l25 = s0i32
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = -1
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
		goto lbl141
	}
	// get_local
	s0i32 = l41
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l24 = s0i32
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// loop
lbl142:
	// get_local
	s0i32 = l21
	// const
	s1i32 = 100000000
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l24
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l24 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl142
	}
	// end
	// end_block
lbl141:
	// get_local
	s0i32 = l25
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl140
	}
	// loop
lbl143:
	// get_local
	s0i32 = l21
	// const
	s1i32 = 10
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l25
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l25 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl143
	}
	// end
	// end_block
lbl140:
	// get_local
	s0i32 = l29
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l41 = s0i32
	// block
	// block
	// get_local
	s0i32 = l29
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l25 = s0i32
	// get_local
	s1i32 = l25
	// get_local
	s2i32 = l21
	// binary: i32.div_u
	s1i32 = i32DivU(s1i32, s2i32)
	// tee_local
	l39 = s1i32
	// get_local
	s2i32 = l21
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l24 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl145
	}
	// get_local
	s0i32 = l41
	// get_local
	s1i32 = l20
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl144
	}
	// end_block
lbl145:
	// block
	// block
	// get_local
	s0i32 = l39
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl147
	}
	// const
	s0f64 = 9.007199254740992e+15
	// set_local
	l35 = s0f64
	// get_local
	s0i32 = l21
	// const
	s1i32 = 1000000000
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl146
	}
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l22
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl146
	}
	// get_local
	s0i32 = l29
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
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
		goto lbl146
	}
	// end_block
lbl147:
	// const
	s0f64 = 9.007199254740994e+15
	// set_local
	l35 = s0f64
	// end_block
lbl146:
	// const
	s0f64 = 0.5
	// const
	s1f64 = 1
	// const
	s2f64 = 1.5
	// get_local
	s3i32 = l41
	// get_local
	s4i32 = l20
	// binary: i32.eq
	if s3i32 == s4i32 {
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
	// const
	s2f64 = 1.5
	// get_local
	s3i32 = l24
	// get_local
	s4i32 = l21
	// const
	s5i32 = 1
	// binary: i32.shr_u
	s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
	// tee_local
	l41 = s4i32
	// binary: i32.eq
	if s3i32 == s4i32 {
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
	s2i32 = l24
	// get_local
	s3i32 = l41
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
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
	l42 = s0f64
	// block
	// get_local
	s0i32 = l37
	// br_if
	if s0i32 != 0 {
		goto lbl148
	}
	// get_local
	s0i32 = l38
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 45
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl148
	}
	// get_local
	s0f64 = l42
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l42 = s0f64
	// get_local
	s0f64 = l35
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l35 = s0f64
	// end_block
lbl148:
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l25
	// get_local
	s2i32 = l24
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l25 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0f64 = l35
	// get_local
	s1f64 = l42
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l35
	// binary: f64.eq
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl144
	}
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l25
	// get_local
	s2i32 = l21
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l21 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l21
	// const
	s1i32 = 1000000000
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl149
	}
	// get_local
	s0i32 = l40
	// const
	s1i32 = -4100
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// loop
lbl150:
	// get_local
	s0i32 = l21
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l22
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl151
	}
	// get_local
	s0i32 = l22
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l22 = s0i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl151:
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l21
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l25 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l21
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l25
	// const
	s1i32 = 999999999
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl150
	}
	// end
	// get_local
	s0i32 = l21
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l29 = s0i32
	// end_block
lbl149:
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l22
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 2
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// const
	s1i32 = 9
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l22
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l25 = s0i32
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
		goto lbl144
	}
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// loop
lbl152:
	// get_local
	s0i32 = l26
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l21
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// tee_local
	l21 = s1i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl152
	}
	// end
	// end_block
lbl144:
	// get_local
	s0i32 = l29
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l21
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l20 = s0i32
	// end_block
lbl139:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l21 = s0i32
	// block
	// loop
lbl154:
	// get_local
	s0i32 = l21
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l20
	// tee_local
	l29 = s0i32
	// get_local
	s1i32 = l22
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l24 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl153
	}
	// get_local
	s0i32 = l25
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l29
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
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
		goto lbl154
	}
	// end
	// end_block
lbl153:
	// block
	// block
	// get_local
	s0i32 = l28
	// br_if
	if s0i32 != 0 {
		goto lbl156
	}
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l39 = s0i32
	// br
	goto lbl155
	// end_block
lbl156:
	// get_local
	s0i32 = l26
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// const
	s1i32 = -1
	// get_local
	s2i32 = l23
	// const
	s3i32 = 1
	// get_local
	s4i32 = l23
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// tee_local
	l20 = s2i32
	// get_local
	s3i32 = l26
	// binary: i32.gt_s
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// get_local
	s3i32 = l26
	// const
	s4i32 = -5
	// binary: i32.gt_s
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l21 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// get_local
	s1i32 = l20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l23 = s0i32
	// const
	s0i32 = -1
	// const
	s1i32 = -2
	// get_local
	s2i32 = l21
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// get_local
	s1i32 = l31
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l39 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl155
	}
	// const
	s0i32 = -9
	// set_local
	l20 = s0i32
	// block
	// get_local
	s0i32 = l24
	// br_if
	if s0i32 != 0 {
		goto lbl157
	}
	// get_local
	s0i32 = l29
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l24 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl157
	}
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l24
	// const
	s1i32 = 10
	// binary: i32.rem_u
	s0i32 = i32RemU(s0i32, s1i32)
	// br_if
	if s0i32 != 0 {
		goto lbl157
	}
	// const
	s0i32 = 10
	// set_local
	l21 = s0i32
	// const
	s0i32 = 0
	// set_local
	l20 = s0i32
	// loop
lbl158:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l21
	// const
	s2i32 = 10
	// binary: i32.mul
	s1i32 = s1i32 * s2i32
	// tee_local
	l21 = s1i32
	// binary: i32.rem_u
	s0i32 = i32RemU(s0i32, s1i32)
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl158
	}
	// end
	// end_block
lbl157:
	// get_local
	s0i32 = l25
	// const
	s1i32 = 2
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// const
	s1i32 = 9
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l31
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
		goto lbl159
	}
	// const
	s0i32 = 0
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l20
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l20 = s1i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l20
	// const
	s4i32 = 0
	// binary: i32.gt_s
	if s3i32 > s4i32 {
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
	// tee_local
	l20 = s1i32
	// get_local
	s2i32 = l23
	// get_local
	s3i32 = l20
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
	l23 = s0i32
	// br
	goto lbl155
	// end_block
lbl159:
	// const
	s0i32 = 0
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l26
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l20
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l20 = s1i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l20
	// const
	s4i32 = 0
	// binary: i32.gt_s
	if s3i32 > s4i32 {
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
	// tee_local
	l20 = s1i32
	// get_local
	s2i32 = l23
	// get_local
	s3i32 = l20
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
	l23 = s0i32
	// end_block
lbl155:
	// get_local
	s0i32 = l23
	// const
	s1i32 = 2147483645
	// const
	s2i32 = 2147483646
	// get_local
	s3i32 = l23
	// get_local
	s4i32 = l39
	// binary: i32.or
	s3i32 = s3i32 | s4i32
	// tee_local
	l37 = s3i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l37
	// const
	s2i32 = 0
	// binary: i32.ne
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l41 = s0i32
	// block
	// block
	// get_local
	s0i32 = l31
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
	// tee_local
	l40 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl161
	}
	// get_local
	s0i32 = l26
	// const
	s1i32 = 2147483647
	// get_local
	s2i32 = l41
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
		goto lbl2
	}
	// get_local
	s0i32 = l26
	// const
	s1i32 = 0
	// get_local
	s2i32 = l26
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
	// set_local
	l20 = s0i32
	// br
	goto lbl160
	// end_block
lbl161:
	// get_local
	s0i32 = l6
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l6
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l26
	// const
	s2i32 = 31
	// binary: i32.shr_s
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	// tee_local
	l20 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl162
	}
	// loop
lbl163:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l24 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l25
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l24
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l28
	// br_if
	if s0i32 != 0 {
		goto lbl163
	}
	// end
	// end_block
lbl162:
	// block
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l25
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 1
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl164
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l13
	// get_local
	s2i32 = l25
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l12
	// get_local
	s3i32 = l25
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl164:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l28 = s0i32
	// get_local
	s1i32 = l31
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 45
	// const
	s2i32 = 43
	// get_local
	s3i32 = l26
	// const
	s4i32 = 0
	// binary: i32.lt_s
	if s3i32 < s4i32 {
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
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l28
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 2147483647
	// get_local
	s2i32 = l41
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
		goto lbl2
	}
	// end_block
lbl160:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l41
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// get_local
	s1i32 = l36
	// const
	s2i32 = 2147483647
	// binary: i32.xor
	s1i32 = s1i32 ^ s2i32
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l36
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l41 = s0i32
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l30 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl165
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l41
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl165
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l41
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l21 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l21
	// br_if
	if s0i32 != 0 {
		goto lbl166
	}
	// loop
lbl167:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl168
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl168:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl167
	}
	// end
	// end_block
lbl166:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl165
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl165:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl169
	}
	// get_local
	s0i32 = l38
	// get_local
	s1i32 = l36
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl169:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 65536
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl170
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l41
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl170
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l41
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l21 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l21
	// br_if
	if s0i32 != 0 {
		goto lbl171
	}
	// loop
lbl172:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl173
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl173:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl172
	}
	// end
	// end_block
lbl171:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl170
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl170:
	// block
	// block
	// get_local
	s0i32 = l40
	// br_if
	if s0i32 != 0 {
		goto lbl175
	}
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l22
	// get_local
	s3i32 = l32
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l28 = s0i32
	// set_local
	l24 = s0i32
	// loop
lbl176:
	// block
	// block
	// get_local
	s0i32 = l24
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l20 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl178
	}
	// get_local
	s0i32 = l11
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l11
	// set_local
	l21 = s0i32
	// br
	goto lbl177
	// end_block
lbl178:
	// get_local
	s0i32 = l11
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l11
	// set_local
	l21 = s0i32
	// loop
lbl179:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l25 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l22
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl179
	}
	// end
	// end_block
lbl177:
	// block
	// block
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l28
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl181
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl180
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l22
	// get_local
	s3i32 = l5
	// const
	s4i32 = 80
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// br
	goto lbl180
	// end_block
lbl181:
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l11
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl180
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 48
	// store: i32.store8
	ctx.Mem[int(s0i32+88)] = uint8(s1i32)
	// get_local
	s0i32 = l10
	// set_local
	l21 = s0i32
	// end_block
lbl180:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl182
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l11
	// get_local
	s2i32 = l21
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl182:
	// get_local
	s0i32 = l24
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l24 = s0i32
	// get_local
	s1i32 = l32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl176
	}
	// end
	// block
	// get_local
	s0i32 = l37
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl183
	}
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl183
	}
	// const
	s0i32 = 6127
	// const
	s1i32 = 1
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl183:
	// block
	// block
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l29
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl185
	}
	// get_local
	s0i32 = l23
	// set_local
	l20 = s0i32
	// br
	goto lbl184
	// end_block
lbl185:
	// block
	// get_local
	s0i32 = l23
	// const
	s1i32 = 1
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl186
	}
	// get_local
	s0i32 = l23
	// set_local
	l20 = s0i32
	// br
	goto lbl184
	// end_block
lbl186:
	// loop
lbl187:
	// block
	// block
	// block
	// get_local
	s0i32 = l24
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l20 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl190
	}
	// get_local
	s0i32 = l11
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l11
	// set_local
	l22 = s0i32
	// br
	goto lbl189
	// end_block
lbl190:
	// get_local
	s0i32 = l11
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l11
	// set_local
	l21 = s0i32
	// loop
lbl191:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l25 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l22
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl191
	}
	// end
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl188
	}
	// end_block
lbl189:
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l22
	// get_local
	s3i32 = l5
	// const
	s4i32 = 80
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl188:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl192
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l23
	// const
	s2i32 = 9
	// get_local
	s3i32 = l23
	// const
	s4i32 = 9
	// binary: i32.lt_s
	if s3i32 < s4i32 {
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
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl192:
	// get_local
	s0i32 = l23
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l24
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l24 = s0i32
	// get_local
	s1i32 = l29
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl184
	}
	// get_local
	s0i32 = l23
	// const
	s1i32 = 9
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l21
	// br_if
	if s0i32 != 0 {
		goto lbl187
	}
	// end
	// end_block
lbl184:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 48
	// get_local
	s2i32 = l20
	// const
	s3i32 = 9
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 9
	// const
	s4i32 = 0
	// call
	f571(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// br
	goto lbl174
	// end_block
lbl175:
	// block
	// get_local
	s0i32 = l23
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
		goto lbl193
	}
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l22
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l29
	// get_local
	s3i32 = l22
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
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
	l29 = s0i32
	// get_local
	s0i32 = l22
	// set_local
	l24 = s0i32
	// loop
lbl194:
	// block
	// block
	// get_local
	s0i32 = l24
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l20 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl196
	}
	// const
	s0i32 = 0
	// set_local
	l21 = s0i32
	// loop
lbl197:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l25 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l25
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl197
	}
	// end
	// get_local
	s0i32 = l21
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl196
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// br
	goto lbl195
	// end_block
lbl196:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 48
	// store: i32.store8
	ctx.Mem[int(s0i32+88)] = uint8(s1i32)
	// get_local
	s0i32 = l10
	// set_local
	l20 = s0i32
	// end_block
lbl195:
	// block
	// block
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l22
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl199
	}
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl198
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l5
	// const
	s4i32 = 80
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// br
	goto lbl198
	// end_block
lbl199:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl200
	}
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl200:
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// block
	// get_local
	s0i32 = l39
	// br_if
	if s0i32 != 0 {
		goto lbl201
	}
	// get_local
	s0i32 = l23
	// const
	s1i32 = 1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl198
	}
	// end_block
lbl201:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl198
	}
	// const
	s0i32 = 6127
	// const
	s1i32 = 1
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl198:
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l20
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl202
	}
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l23
	// get_local
	s3i32 = l23
	// get_local
	s4i32 = l21
	// binary: i32.gt_s
	if s3i32 > s4i32 {
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
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl202:
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l21
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l24
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l24 = s0i32
	// get_local
	s1i32 = l29
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl193
	}
	// get_local
	s0i32 = l23
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
		goto lbl194
	}
	// end
	// end_block
lbl193:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 48
	// get_local
	s2i32 = l23
	// const
	s3i32 = 18
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 18
	// const
	s4i32 = 0
	// call
	f571(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl174
	}
	// get_local
	s0i32 = l28
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l28
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl174:
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8192
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl203
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l41
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl203
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l41
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l21 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l21
	// br_if
	if s0i32 != 0 {
		goto lbl204
	}
	// loop
lbl205:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl206
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl206:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl205
	}
	// end
	// end_block
lbl204:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl203
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl203:
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l41
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l41
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
	l20 = s0i32
	// br
	goto lbl4
	// end_block
lbl53:
	// const
	s0i32 = 0
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1091
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l30
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l24
	// set_local
	l26 = s0i32
	// end_block
lbl52:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l24 = s0i32
	// get_local
	s1i32 = l26
	// get_local
	s2i32 = l26
	// get_local
	s3i32 = l24
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
	// tee_local
	l28 = s0i32
	// const
	s1i32 = 2147483647
	// get_local
	s2i32 = l23
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
		goto lbl2
	}
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l25 = s0i32
	// get_local
	s1i32 = l27
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l25
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
	// tee_local
	l20 = s0i32
	// get_local
	s1i32 = l22
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// block
	// get_local
	s0i32 = l29
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l29 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl207
	}
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l27
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl207
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l25
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l22 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l22
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l30 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l30
	// br_if
	if s0i32 != 0 {
		goto lbl208
	}
	// loop
lbl209:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl210
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl210:
	// get_local
	s0i32 = l22
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l22 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl209
	}
	// end
	// end_block
lbl208:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl207
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl207:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl211
	}
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l23
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl211:
	// block
	// get_local
	s0i32 = l29
	// const
	s1i32 = 65536
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl212
	}
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l27
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl212
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l25
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l22 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l22
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l23 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l23
	// br_if
	if s0i32 != 0 {
		goto lbl213
	}
	// loop
lbl214:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl215
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl215:
	// get_local
	s0i32 = l22
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l22 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl214
	}
	// end
	// end_block
lbl213:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl212
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl212:
	// block
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l26
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl216
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l28
	// get_local
	s3i32 = l24
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l22 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l22
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l26 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l26
	// br_if
	if s0i32 != 0 {
		goto lbl217
	}
	// loop
lbl218:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl219
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl219:
	// get_local
	s0i32 = l22
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l22 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl218
	}
	// end
	// end_block
lbl217:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl216
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl216:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl220
	}
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l24
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl220:
	// get_local
	s0i32 = l29
	// const
	s1i32 = 8192
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
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l27
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l20
	// get_local
	s3i32 = l25
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l21 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l21
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl221
	}
	// loop
lbl222:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl223
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl223:
	// get_local
	s0i32 = l21
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl222
	}
	// end
	// end_block
lbl221:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 112
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// br
	goto lbl3
	// end_block
lbl6:
	// const
	s0i32 = 0
	// const
	s1i32 = 28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// br
	goto lbl1
	// end_block
lbl5:
	// get_local
	s0i32 = l38
	// get_local
	s1i32 = l31
	// const
	s2i32 = 26
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 31
	// binary: i32.shr_s
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	// const
	s2i32 = 9
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l23 = s0i32
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = 11
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl224
	}
	// block
	// block
	// const
	s0i32 = 12
	// get_local
	s1i32 = l24
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l21 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl226
	}
	// const
	s0f64 = 16
	// set_local
	l42 = s0f64
	// br
	goto lbl225
	// end_block
lbl226:
	// get_local
	s0i32 = l24
	// const
	s1i32 = -12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// const
	s0f64 = 16
	// set_local
	l42 = s0f64
	// loop
lbl227:
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0f64 = l42
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l42 = s0f64
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl227
	}
	// end
	// const
	s0i32 = 0
	// get_local
	s1i32 = l20
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l20 = s0i32
	// end_block
lbl225:
	// block
	// get_local
	s0i32 = l24
	// const
	s1i32 = -5
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
		goto lbl228
	}
	// loop
lbl229:
	// get_local
	s0f64 = l42
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l42 = s0f64
	// get_local
	s0i32 = l20
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl229
	}
	// end
	// end_block
lbl228:
	// block
	// get_local
	s0i32 = l23
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 45
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl230
	}
	// get_local
	s0f64 = l42
	// get_local
	s1f64 = l35
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2f64 = l42
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l35 = s0f64
	// br
	goto lbl224
	// end_block
lbl230:
	// get_local
	s0f64 = l35
	// get_local
	s1f64 = l42
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l42
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l35 = s0f64
	// end_block
lbl224:
	// block
	// block
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+108):]))
	// tee_local
	l26 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl232
	}
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l26
	// const
	s2i32 = 31
	// binary: i32.shr_s
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	// tee_local
	l20 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l20 = s0i32
	// const
	s0i32 = 0
	// set_local
	l21 = s0i32
	// loop
lbl233:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 68
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 11
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l20
	// const
	s3i32 = 10
	// binary: i32.div_u
	s2i32 = i32DivU(s2i32, s3i32)
	// tee_local
	l22 = s2i32
	// const
	s3i32 = 10
	// binary: i32.mul
	s2i32 = s2i32 * s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 48
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l21
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = 9
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l22
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l25
	// br_if
	if s0i32 != 0 {
		goto lbl233
	}
	// end
	// get_local
	s0i32 = l21
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl232
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 68
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// br
	goto lbl231
	// end_block
lbl232:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 48
	// store: i32.store8
	ctx.Mem[int(s0i32+79)] = uint8(s1i32)
	// get_local
	s0i32 = l9
	// set_local
	l20 = s0i32
	// end_block
lbl231:
	// get_local
	s0i32 = l36
	// const
	s1i32 = 2
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l31
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l20
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l29 = s0i32
	// get_local
	s1i32 = l31
	// const
	s2i32 = 15
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l20
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 45
	// const
	s2i32 = 43
	// get_local
	s3i32 = l26
	// const
	s4i32 = 0
	// binary: i32.lt_s
	if s3i32 < s4i32 {
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
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l30
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// loop
lbl234:
	// get_local
	s0i32 = l21
	// set_local
	l20 = s0i32
	// block
	// block
	// get_local
	s0f64 = l35
	// unary: f64.abs
	s0f64 = math.Abs(s0f64)
	// const
	s1f64 = 2.147483648e+09
	// binary: f64.lt
	if s0f64 < s1f64 {
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
		goto lbl236
	}
	// get_local
	s0f64 = l35
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l21 = s0i32
	// br
	goto lbl235
	// end_block
lbl236:
	// const
	s0i32 = -2147483648
	// set_local
	l21 = s0i32
	// end_block
lbl235:
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// const
	s2i32 = 23376
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// get_local
	s2i32 = l22
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0f64 = l35
	// get_local
	s1i32 = l21
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l35 = s0f64
	// block
	// get_local
	s0i32 = l20
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l21 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 80
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
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
		goto lbl237
	}
	// block
	// get_local
	s0i32 = l25
	// br_if
	if s0i32 != 0 {
		goto lbl238
	}
	// get_local
	s0i32 = l24
	// const
	s1i32 = 0
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl238
	}
	// get_local
	s0f64 = l35
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
		goto lbl237
	}
	// end_block
lbl238:
	// get_local
	s0i32 = l20
	// const
	s1i32 = 46
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// get_local
	s0i32 = l20
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// end_block
lbl237:
	// get_local
	s0f64 = l35
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
		goto lbl234
	}
	// end
	// const
	s0i32 = 2147483645
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l29
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l32 = s1i32
	// get_local
	s2i32 = l28
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l20 = s1i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l24
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l24
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l5
	// const
	s3i32 = 80
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l22 = s1i32
	// get_local
	s2i32 = l8
	// get_local
	s3i32 = l21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l24
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
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l24
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l26 = s0i32
	// get_local
	s1i32 = l20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l21 = s0i32
	// block
	// get_local
	s0i32 = l30
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l25 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl239
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl239
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l24 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l24
	// br_if
	if s0i32 != 0 {
		goto lbl240
	}
	// loop
lbl241:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl242
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl242:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl241
	}
	// end
	// end_block
lbl240:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl239
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl239:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl243
	}
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l28
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl243:
	// block
	// get_local
	s0i32 = l25
	// const
	s1i32 = 65536
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl244
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl244
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l24 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l24
	// br_if
	if s0i32 != 0 {
		goto lbl245
	}
	// loop
lbl246:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl247
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl247:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl246
	}
	// end
	// end_block
lbl245:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl244
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl244:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl248
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl248:
	// block
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l22
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl249
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 48
	// get_local
	s2i32 = l20
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl250
	}
	// loop
lbl251:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl252
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl252:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl251
	}
	// end
	// end_block
lbl250:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl249
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl249:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl253
	}
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l32
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl253:
	// block
	// get_local
	s0i32 = l25
	// const
	s1i32 = 8192
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl254
	}
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl254
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l20 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l20
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l22 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// block
	// get_local
	s0i32 = l22
	// br_if
	if s0i32 != 0 {
		goto lbl255
	}
	// loop
lbl256:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl257
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl257:
	// get_local
	s0i32 = l20
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l20 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl256
	}
	// end
	// end_block
lbl255:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl254
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l20
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl254:
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l27
	// get_local
	s3i32 = l21
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
	l20 = s0i32
	// end_block
lbl4:
	// get_local
	s0i32 = l20
	// const
	s1i32 = 0
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// end
	// end_block
lbl2:
	// const
	s0i32 = 0
	// const
	s1i32 = 61
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// end_block
lbl1:
	// const
	s0i32 = -1
	// set_local
	l19 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 880
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l19
	// return
	return s0i32
}
