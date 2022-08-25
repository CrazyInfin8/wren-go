package internals

import (
	"encoding/binary"
	"math"
)

func f607(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var l15 float64
	_ = l15
	var l16 float64
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
	var l33 int32
	_ = l33
	var l34 int32
	_ = l34
	var l35 float64
	_ = l35
	var l36 float64
	_ = l36
	var l37 float64
	_ = l37
	var l38 float64
	_ = l38
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
	s1i32 = 560
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// set_global
	ctx.G0 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 24
	// binary: i32.div_s
	s0i32 = i32DivS(s0i32, s1i32)
	// tee_local
	l7 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l7
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
	// tee_local
	l8 = s0i32
	// const
	s1i32 = -24
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 23824
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l10 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l11 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
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
		goto lbl0
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l11
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l13 = s0i32
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
		goto lbl1
	}
	// get_local
	s0i32 = l7
	// const
	s1i32 = -2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 23848
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 320
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// loop
lbl2:
	// const
	s0f64 = 0
	// set_local
	l15 = s0f64
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// block
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l17 = s0i32
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
		goto lbl3
	}
	// get_local
	s0i32 = l7
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l16 = s0f64
	// end_block
lbl3:
	// get_local
	s0i32 = l2
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// block
	// get_local
	s0i32 = l17
	// const
	s1i32 = -1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l15 = s0f64
	// end_block
lbl4:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l15
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l6
	// const
	s2i32 = 2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l6 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// end
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l12 = s0i32
	// end_block
lbl1:
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
		goto lbl0
	}
	// block
	// block
	// get_local
	s0i32 = l12
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
		goto lbl6
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// br
	goto lbl5
	// end_block
lbl6:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 23840
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l16 = s0f64
	// end_block
lbl5:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 320
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// end_block
lbl0:
	// get_local
	s0i32 = l9
	// const
	s1i32 = -24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l18 = s0i32
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l10
	// const
	s1i32 = 0
	// get_local
	s2i32 = l10
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
	l13 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = -2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// loop
lbl7:
	// get_local
	s0i32 = l2
	// set_local
	l12 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
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
		goto lbl9
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// br
	goto lbl8
	// end_block
lbl9:
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// block
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
		goto lbl10
	}
	// get_local
	s0i32 = l14
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l6 = s0i32
	// loop
lbl11:
	// get_local
	s0i32 = l6
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1i32 = l6
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l2
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l16
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l7
	// const
	s2i32 = 2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l7 = s1i32
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
	// end
	// end_block
lbl10:
	// get_local
	s0i32 = l19
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
	s0i32 = l0
	// get_local
	s1i32 = l7
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l5
	// const
	s2i32 = 320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l12
	// get_local
	s3i32 = l11
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l7
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 3
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// end_block
lbl8:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l14
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l13
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// end
	// get_local
	s0i32 = l3
	// const
	s1i32 = -2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l10
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l21 = s0i32
	// const
	s0i32 = 47
	// get_local
	s1i32 = l9
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l22 = s0i32
	// const
	s0i32 = 48
	// get_local
	s1i32 = l9
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l10
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 480
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 320
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l9
	// const
	s1i32 = -25
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l10
	// set_local
	l12 = s0i32
	// block
	// loop
lbl13:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l2 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l16 = s0f64
	// block
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l19 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l30 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l12
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
		goto lbl16
	}
	// get_local
	s0i32 = l12
	// set_local
	l2 = s0i32
	// br
	goto lbl15
	// end_block
lbl16:
	// get_local
	s0i32 = l12
	// const
	s1i32 = -2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l28
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// loop
lbl17:
	// block
	// block
	// get_local
	s0f64 = l16
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l15 = s0f64
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
		goto lbl19
	}
	// get_local
	s0f64 = l15
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l14 = s0i32
	// br
	goto lbl18
	// end_block
lbl19:
	// const
	s0i32 = -2147483648
	// set_local
	l14 = s0i32
	// end_block
lbl18:
	// block
	// block
	// get_local
	s0i32 = l14
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// tee_local
	l15 = s0f64
	// const
	s1f64 = -1.6777216e+07
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
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
		goto lbl21
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l14 = s0i32
	// br
	goto lbl20
	// end_block
lbl21:
	// const
	s0i32 = -2147483648
	// set_local
	l14 = s0i32
	// end_block
lbl20:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l14
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1f64 = l15
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l15 = s0f64
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
		goto lbl23
	}
	// get_local
	s0f64 = l15
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l14 = s0i32
	// br
	goto lbl22
	// end_block
lbl23:
	// const
	s0i32 = -2147483648
	// set_local
	l14 = s0i32
	// end_block
lbl22:
	// block
	// block
	// get_local
	s0i32 = l14
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// tee_local
	l15 = s0f64
	// const
	s1f64 = -1.6777216e+07
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
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
		goto lbl25
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l14 = s0i32
	// br
	goto lbl24
	// end_block
lbl25:
	// const
	s0i32 = -2147483648
	// set_local
	l14 = s0i32
	// end_block
lbl24:
	// get_local
	s0i32 = l6
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l14
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1f64 = l15
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l6
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l13
	// get_local
	s1i32 = l7
	// const
	s2i32 = 2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l7 = s1i32
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
	// end
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l7
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l2 = s0i32
	// end_block
lbl15:
	// get_local
	s0i32 = l30
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
	s0i32 = l7
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0f64 = l16
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l15 = s0f64
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
		goto lbl27
	}
	// get_local
	s0f64 = l15
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l7 = s0i32
	// br
	goto lbl26
	// end_block
lbl27:
	// const
	s0i32 = -2147483648
	// set_local
	l7 = s0i32
	// end_block
lbl26:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0i32 = l7
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// tee_local
	l15 = s0f64
	// const
	s1f64 = -1.6777216e+07
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
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
		goto lbl29
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l7 = s0i32
	// br
	goto lbl28
	// end_block
lbl29:
	// const
	s0i32 = -2147483648
	// set_local
	l7 = s0i32
	// end_block
lbl28:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1f64 = l15
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// end_block
lbl14:
	// block
	// block
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l18
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// tee_local
	l16 = s0f64
	// const
	s1f64 = 0.125
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// unary: f64.floor
	s0f64 = math.Floor(s0f64)
	// const
	s1f64 = -8
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
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
		goto lbl31
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l31 = s0i32
	// br
	goto lbl30
	// end_block
lbl31:
	// const
	s0i32 = -2147483648
	// set_local
	l31 = s0i32
	// end_block
lbl30:
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l31
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l16 = s0f64
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l18
	// const
	s1i32 = 1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l32 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 480
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l2 = s1i32
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l23
	// binary: i32.shr_s
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	// tee_local
	l2 = s2i32
	// get_local
	s3i32 = l23
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l6 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l22
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l31
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l31 = s0i32
	// br
	goto lbl35
	// end_block
lbl36:
	// get_local
	s0i32 = l18
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 480
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// const
	s1i32 = 23
	// binary: i32.shr_s
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	// set_local
	l33 = s0i32
	// end_block
lbl35:
	// get_local
	s0i32 = l33
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
		goto lbl32
	}
	// br
	goto lbl33
	// end_block
lbl34:
	// const
	s0i32 = 2
	// set_local
	l33 = s0i32
	// get_local
	s0f64 = l16
	// const
	s1f64 = 0.5
	// binary: f64.ge
	if s0f64 >= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl33
	}
	// const
	s0i32 = 0
	// set_local
	l33 = s0i32
	// br
	goto lbl32
	// end_block
lbl33:
	// block
	// block
	// get_local
	s0i32 = l19
	// unary: i32.eqz
	if s0i32 == 0 {
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
	// set_local
	l6 = s0i32
	// br
	goto lbl37
	// end_block
lbl38:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l34 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// block
	// get_local
	s0i32 = l12
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
		goto lbl39
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = -2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l30 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// loop
lbl40:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l7 = s0i32
	// const
	s0i32 = 16777215
	// set_local
	l14 = s0i32
	// block
	// block
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl42
	}
	// const
	s0i32 = 16777216
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl42
	}
	// const
	s0i32 = 1
	// set_local
	l14 = s0i32
	// br
	goto lbl41
	// end_block
lbl42:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l14
	// get_local
	s2i32 = l7
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l14 = s0i32
	// end_block
lbl41:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l19 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l7 = s0i32
	// const
	s0i32 = 16777215
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0i32 = l14
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl44
	}
	// const
	s0i32 = 16777216
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl44
	}
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// br
	goto lbl43
	// end_block
lbl44:
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l7
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// end_block
lbl43:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l30
	// get_local
	s1i32 = l13
	// const
	s2i32 = 2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l13 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl40
	}
	// end
	// end_block
lbl39:
	// get_local
	s0i32 = l34
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl37
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l13
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l14 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l2 = s0i32
	// const
	s0i32 = 16777215
	// set_local
	l7 = s0i32
	// block
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// const
	s0i32 = 16777216
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l2
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// br
	goto lbl37
	// end_block
lbl45:
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l7
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// end_block
lbl37:
	// block
	// get_local
	s0i32 = l32
	// br_if
	if s0i32 != 0 {
		goto lbl46
	}
	// const
	s0i32 = 8388607
	// set_local
	l2 = s0i32
	// block
	// block
	// get_local
	s0i32 = l29
	// br_table
	switch s0i32 {
	case 0:
		goto lbl47
	case 1:
		goto lbl48
	default:
		goto lbl46
	}
	// end_block
lbl48:
	// const
	s0i32 = 4194303
	// set_local
	l2 = s0i32
	// end_block
lbl47:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 480
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// get_local
	s1i32 = l7
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl46:
	// get_local
	s0i32 = l31
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l33
	// const
	s1i32 = 2
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl32
	}
	// const
	s0f64 = 1
	// get_local
	s1f64 = l16
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l16 = s0f64
	// const
	s0i32 = 2
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl32
	}
	// get_local
	s0f64 = l16
	// const
	s1f64 = 1
	// get_local
	s2i32 = l18
	// call
	s1f64 = f627(ctx, s1f64, s2i32)
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l16 = s0f64
	// end_block
lbl32:
	// block
	// get_local
	s0f64 = l16
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
		goto lbl49
	}
	// block
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l10
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl50
	}
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l10
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l14 = s0i32
	// block
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl51
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l12
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l14 = s0i32
	// loop
lbl52:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l2
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l2
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+0):]))
	// get_local
	s3i32 = l2
	// const
	s4i32 = 12
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+0):]))
	// get_local
	s4i32 = l6
	// binary: i32.or
	s3i32 = s3i32 | s4i32
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l14
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l13
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l13 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl52
	}
	// end
	// end_block
lbl51:
	// block
	// get_local
	s0i32 = l7
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl53
	}
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l14
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl54:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l6
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl54
	}
	// end
	// end_block
lbl53:
	// get_local
	s0i32 = l6
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
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l18
	// set_local
	l9 = s0i32
	// loop
lbl55:
	// get_local
	s0i32 = l12
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l9
	// const
	s1i32 = -24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl55
	}
	// br
	goto lbl12
	// end
	// end_block
lbl50:
	// get_local
	s0i32 = l24
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l14 = s0i32
	// loop
lbl56:
	// get_local
	s0i32 = l14
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// end
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l12
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// loop
lbl57:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 320
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// get_local
	s2i32 = l3
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l19 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l12 = s1i32
	// get_local
	s2i32 = l8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 23840
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// block
	// block
	// get_local
	s0i32 = l3
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
		goto lbl59
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// br
	goto lbl58
	// end_block
lbl59:
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// block
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
		goto lbl60
	}
	// get_local
	s0i32 = l13
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l6 = s0i32
	// loop
lbl61:
	// get_local
	s0i32 = l6
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1i32 = l6
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l2
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l16
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l7
	// const
	s2i32 = 2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l7 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl61
	}
	// end
	// end_block
lbl60:
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
		goto lbl58
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l7
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l5
	// const
	s2i32 = 320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// get_local
	s2i32 = l19
	// get_local
	s3i32 = l7
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 3
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// end_block
lbl58:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l13
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l14
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl57
	}
	// end
	// get_local
	s0i32 = l14
	// set_local
	l12 = s0i32
	// br
	goto lbl13
	// end_block
lbl49:
	// end
	// block
	// block
	// get_local
	s0f64 = l16
	// const
	s1i32 = 24
	// get_local
	s2i32 = l9
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// tee_local
	l16 = s0f64
	// const
	s1f64 = 1.6777216e+07
	// binary: f64.ge
	if s0f64 >= s1f64 {
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
		goto lbl63
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0f64 = l16
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l15 = s0f64
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
		goto lbl65
	}
	// get_local
	s0f64 = l15
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l2 = s0i32
	// br
	goto lbl64
	// end_block
lbl65:
	// const
	s0i32 = -2147483648
	// set_local
	l2 = s0i32
	// end_block
lbl64:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0i32 = l2
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// const
	s1f64 = -1.6777216e+07
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l16 = s0f64
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
		goto lbl67
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l7 = s0i32
	// br
	goto lbl66
	// end_block
lbl67:
	// const
	s0i32 = -2147483648
	// set_local
	l7 = s0i32
	// end_block
lbl66:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l12 = s0i32
	// br
	goto lbl62
	// end_block
lbl63:
	// block
	// block
	// get_local
	s0f64 = l16
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
		goto lbl69
	}
	// get_local
	s0f64 = l16
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l2 = s0i32
	// br
	goto lbl68
	// end_block
lbl69:
	// const
	s0i32 = -2147483648
	// set_local
	l2 = s0i32
	// end_block
lbl68:
	// get_local
	s0i32 = l18
	// set_local
	l9 = s0i32
	// end_block
lbl62:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl12:
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
		goto lbl70
	}
	// const
	s0f64 = 1
	// get_local
	s1i32 = l9
	// call
	s0f64 = f627(ctx, s0f64, s1i32)
	// set_local
	l16 = s0f64
	// block
	// block
	// get_local
	s0i32 = l12
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
		goto lbl72
	}
	// get_local
	s0i32 = l12
	// set_local
	l2 = s0i32
	// br
	goto lbl71
	// end_block
lbl72:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// get_local
	s2i32 = l5
	// const
	s3i32 = 480
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l12
	// const
	s4i32 = 2
	// binary: i32.shl
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+0):]))
	// unary: f64.convert_s/i32
	s2f64 = float64(s2i32)
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l12
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0f64 = l16
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l16 = s0f64
	// end_block
lbl71:
	// block
	// get_local
	s0i32 = l12
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl73
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 480
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// loop
lbl74:
	// get_local
	s0i32 = l6
	// get_local
	s1f64 = l16
	// const
	s2f64 = 5.960464477539063e-08
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l15 = s1f64
	// get_local
	s2i32 = l2
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+0):]))
	// unary: f64.convert_s/i32
	s2f64 = float64(s2i32)
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l6
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// get_local
	s2i32 = l2
	// const
	s3i32 = 4
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+0):]))
	// unary: f64.convert_s/i32
	s2f64 = float64(s2i32)
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0f64 = l15
	// const
	s1f64 = 5.960464477539063e-08
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l7
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl74
	}
	// end
	// end_block
lbl73:
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
		goto lbl70
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l2 = s0i32
	// loop
lbl75:
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l2
	// tee_local
	l13 = s1i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l14 = s0i32
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// block
	// loop
lbl77:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 26608
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l17
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l10
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl76
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l14
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl77
	}
	// end
	// end_block
lbl76:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l14
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l17
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l13
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l13
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
		goto lbl75
	}
	// end
	// end_block
lbl70:
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l4
	// br_table
	switch s0i32 {
	case 0:
		goto lbl83
	case 1:
		goto lbl82
	case 2:
		goto lbl82
	case 3:
		goto lbl84
	default:
		goto lbl78
	}
	// end_block
lbl84:
	// const
	s0f64 = 0
	// set_local
	l35 = s0f64
	// get_local
	s0i32 = l12
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
		goto lbl79
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l15 = s0f64
	// block
	// block
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl86
	}
	// get_local
	s0f64 = l15
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l12
	// set_local
	l2 = s0i32
	// br
	goto lbl85
	// end_block
lbl86:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l17
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l6
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l36 = s1f64
	// get_local
	s2f64 = l15
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l16 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// get_local
	s1f64 = l15
	// get_local
	s2f64 = l36
	// get_local
	s3f64 = l16
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l17
	// set_local
	l2 = s0i32
	// end_block
lbl85:
	// block
	// get_local
	s0i32 = l17
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl87
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 160
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl88:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l37 = s1f64
	// get_local
	s2i32 = l2
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// tee_local
	l7 = s2i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// tee_local
	l38 = s2f64
	// get_local
	s3f64 = l16
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// tee_local
	l15 = s2f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l36 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// get_local
	s2f64 = l38
	// get_local
	s3f64 = l15
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l7
	// get_local
	s1f64 = l15
	// get_local
	s2f64 = l37
	// get_local
	s3f64 = l36
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0f64 = l36
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l6
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 1
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl88
	}
	// end
	// end_block
lbl87:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl79
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l17
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l16 = s0f64
	// loop
lbl89:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l36 = s1f64
	// get_local
	s2f64 = l16
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l15 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1f64 = l16
	// get_local
	s2f64 = l36
	// get_local
	s3f64 = l15
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0f64 = l15
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l6
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 2
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl89
	}
	// end
	// get_local
	s0i32 = l12
	// const
	s1i32 = 2
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl79
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l17
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l6 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl81
	}
	// const
	s0f64 = 0
	// set_local
	l35 = s0f64
	// br
	goto lbl80
	// end_block
lbl83:
	// block
	// block
	// get_local
	s0i32 = l12
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
		goto lbl91
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// br
	goto lbl90
	// end_block
lbl91:
	// block
	// block
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l12
	// set_local
	l6 = s0i32
	// br
	goto lbl92
	// end_block
lbl93:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l12
	// set_local
	l6 = s0i32
	// loop
lbl94:
	// get_local
	s0i32 = l6
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl94
	}
	// end
	// end_block
lbl92:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl90
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 160
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl95:
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// const
	s2i32 = 24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl95
	}
	// end
	// end_block
lbl90:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l16
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2f64 = l16
	// get_local
	s3i32 = l33
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// br
	goto lbl78
	// end_block
lbl82:
	// block
	// block
	// get_local
	s0i32 = l12
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
		goto lbl97
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// br
	goto lbl96
	// end_block
lbl97:
	// block
	// block
	// get_local
	s0i32 = l12
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl99
	}
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l12
	// set_local
	l6 = s0i32
	// br
	goto lbl98
	// end_block
lbl99:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0f64 = 0
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l12
	// set_local
	l6 = s0i32
	// loop
lbl100:
	// get_local
	s0i32 = l6
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// end
	// end_block
lbl98:
	// get_local
	s0i32 = l12
	// const
	s1i32 = 3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl96
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 160
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl101:
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// const
	s2i32 = 24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l7
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l7 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl101
	}
	// end
	// end_block
lbl96:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l16
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2f64 = l16
	// get_local
	s3i32 = l33
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l5
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+160):]))
	// get_local
	s1f64 = l16
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l16 = s0f64
	// const
	s0i32 = 1
	// set_local
	l2 = s0i32
	// block
	// get_local
	s0i32 = l12
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
		goto lbl102
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l6 = s0i32
	// block
	// get_local
	s0i32 = l12
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl103
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = -4
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// loop
lbl104:
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// const
	s2i32 = -24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = -16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = -8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l7
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l7 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl104
	}
	// end
	// get_local
	s0i32 = l7
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// end_block
lbl103:
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl105:
	// get_local
	s0f64 = l16
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l16 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl105
	}
	// end
	// end_block
lbl102:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l16
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2f64 = l16
	// get_local
	s3i32 = l33
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// br
	goto lbl78
	// end_block
lbl81:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l12
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0f64 = 0
	// set_local
	l35 = s0f64
	// loop
lbl106:
	// get_local
	s0i32 = l12
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l12 = s0i32
	// get_local
	s0f64 = l35
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l35 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// end
	// end_block
lbl80:
	// get_local
	s0i32 = l7
	// const
	s1i32 = 3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl79
	}
	// get_local
	s0i32 = l12
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l12
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// const
	s2i32 = 160
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl107:
	// get_local
	s0f64 = l35
	// get_local
	s1i32 = l2
	// const
	s2i32 = 24
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 16
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l35 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = -32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l6
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 5
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl107
	}
	// end
	// end_block
lbl79:
	// get_local
	s0i32 = l5
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+160):]))
	// set_local
	l16 = s0f64
	// block
	// get_local
	s0i32 = l33
	// br_if
	if s0i32 != 0 {
		goto lbl108
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l16
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l35
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+168):]))
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// br
	goto lbl78
	// end_block
lbl108:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l16
	// unary: f64.neg
	s1f64 = -s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l35
	// unary: f64.neg
	s1f64 = -s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+168):]))
	// unary: f64.neg
	s1f64 = -s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// end_block
lbl78:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 560
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l31
	// const
	s1i32 = 7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// return
	return s0i32
}
