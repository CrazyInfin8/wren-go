package internals

import (
	"encoding/binary"
	"math"
)

func f612(ctx *Context, l0 float64, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
	var l9 int32
	_ = l9
	var l10 float64
	_ = l10
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
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
	s1i32 = 48
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l2 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// block
	// block
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l3 = s0i64
	// const
	s1i64 = 32
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 1074752122
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 1048575
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 598523
	// binary: i32.eq
	if s0i32 == s1i32 {
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
	s0i32 = l5
	// const
	s1i32 = 1073928572
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
	// block
	// get_local
	s0i64 = l3
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
		goto lbl5
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = -1.5707963267341256
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = -6.077100506506192e-11
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = -6.077100506506192e-11
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 1
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl5:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = 1.5707963267341256
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = 6.077100506506192e-11
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = 6.077100506506192e-11
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = -1
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl4:
	// block
	// get_local
	s0i64 = l3
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
		goto lbl6
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = -3.1415926534682512
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = -1.2154201013012384e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = -1.2154201013012384e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 2
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl6:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = 3.1415926534682512
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = 1.2154201013012384e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = 1.2154201013012384e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = -2
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1075594811
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1075183036
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
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1074977148
	// binary: i32.eq
	if s0i32 == s1i32 {
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
	s0i64 = l3
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
		goto lbl9
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = -4.712388980202377
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = -1.8231301519518578e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = -1.8231301519518578e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 3
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl9:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = 4.712388980202377
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = 1.8231301519518578e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = 1.8231301519518578e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = -3
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl8:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1075388923
	// binary: i32.eq
	if s0i32 == s1i32 {
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
	s0i64 = l3
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
		goto lbl10
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = -6.2831853069365025
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = -2.430840202602477e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = -2.430840202602477e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 4
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl10:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// const
	s2f64 = 6.2831853069365025
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// const
	s2f64 = 2.430840202602477e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = 2.430840202602477e-10
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = -4
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl7:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1094263290
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
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
	s0i32 = l1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// const
	s3f64 = 0.6366197723675814
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = 6.755399441055744e+15
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// const
	s3f64 = -6.755399441055744e+15
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// tee_local
	l6 = s2f64
	// const
	s3f64 = -1.5707963267341256
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l7 = s1f64
	// get_local
	s2f64 = l6
	// const
	s3f64 = 6.077100506506192e-11
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l8 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l0 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l5
	// const
	s1i32 = 20
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l9 = s0i32
	// get_local
	s1f64 = l0
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// const
	s2i64 = 52
	// binary: i64.shr_u
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// const
	s2i32 = 2047
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 17
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l5 = s0i32
	// block
	// block
	// get_local
	s0f64 = l6
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
		goto lbl12
	}
	// get_local
	s0f64 = l6
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l4 = s0i32
	// br
	goto lbl11
	// end_block
lbl12:
	// const
	s0i32 = -2147483648
	// set_local
	l4 = s0i32
	// end_block
lbl11:
	// block
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l7
	// get_local
	s2f64 = l6
	// const
	s3f64 = 6.077100506303966e-11
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l0 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l10 = s1f64
	// get_local
	s2f64 = l6
	// const
	s3f64 = 2.0222662487959506e-21
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l7
	// get_local
	s4f64 = l10
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// get_local
	s4f64 = l0
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// tee_local
	l8 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l0 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// block
	// get_local
	s0i32 = l9
	// get_local
	s1f64 = l0
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// const
	s2i64 = 52
	// binary: i64.shr_u
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// const
	s2i32 = 2047
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 50
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0f64 = l10
	// set_local
	l7 = s0f64
	// br
	goto lbl13
	// end_block
lbl14:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l10
	// get_local
	s2f64 = l6
	// const
	s3f64 = 2.0222662487111665e-21
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l0 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l7 = s1f64
	// get_local
	s2f64 = l6
	// const
	s3f64 = 8.4784276603689e-32
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l10
	// get_local
	s4f64 = l7
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// get_local
	s4f64 = l0
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// tee_local
	l8 = s2f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l0 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// end_block
lbl13:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l7
	// get_local
	s2f64 = l0
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l8
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// br
	goto lbl0
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 2146435072
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
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
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l0 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 0
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl15:
	// block
	// block
	// get_local
	s0i64 = l3
	// const
	s1i64 = 4503599627370495
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 4710765210229538816
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l0 = s0f64
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
		goto lbl17
	}
	// get_local
	s0f64 = l0
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l4 = s0i32
	// br
	goto lbl16
	// end_block
lbl17:
	// const
	s0i32 = -2147483648
	// set_local
	l4 = s0i32
	// end_block
lbl16:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], math.Float64bits(s1f64))
	// block
	// block
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l6
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 1.6777216e+07
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l0 = s0f64
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
	s0f64 = l0
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l4 = s0i32
	// br
	goto lbl18
	// end_block
lbl19:
	// const
	s0i32 = -2147483648
	// set_local
	l4 = s0i32
	// end_block
lbl18:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// unary: f64.convert_s/i32
	s1f64 = float64(s1i32)
	// tee_local
	l6 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+24):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l2
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// const
	s2f64 = 1.6777216e+07
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l0 = s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+32):], math.Float64bits(s1f64))
	// block
	// block
	// get_local
	s0f64 = l0
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
		goto lbl21
	}
	// const
	s0i32 = 2
	// set_local
	l9 = s0i32
	// br
	goto lbl20
	// end_block
lbl21:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 8
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l4 = s0i32
	// const
	s0i32 = 2
	// set_local
	l9 = s0i32
	// loop
lbl22:
	// get_local
	s0i32 = l9
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l4
	// const
	s1i32 = -8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0f64 = l0
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
		goto lbl22
	}
	// end
	// end_block
lbl20:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l5
	// const
	s3i32 = 20
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = -1046
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l9
	// const
	s4i32 = 1
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 1
	// call
	s0i32 = f611(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l2
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0f64
	// block
	// get_local
	s0i64 = l3
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
		goto lbl23
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// unary: f64.neg
	s1f64 = -s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// unary: f64.neg
	s1f64 = -s1f64
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl23:
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// end_block
lbl0:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l4
	// return
	return s0i32
}
