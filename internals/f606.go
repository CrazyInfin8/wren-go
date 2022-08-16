package internals

import (
	"math"
)

func f606(ctx *Context, l0 float64) float64 {
	var l1 int64
	_ = l1
	var l2 int32
	_ = l2
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
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
	var s6f64 float64
	_ = s6f64
	var s7f64 float64
	_ = s7f64
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l1 = s0i64
	// const
	s1i64 = 32
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1072693248
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1072693248
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i64 = l1
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1.5707963267948966
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1f64 = 7.52316384526264e-37
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl1:
	// const
	s0f64 = 0
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1071644671
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
	s0i32 = l2
	// const
	s1i32 = -1048576
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 1044381696
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l3 = s1f64
	// get_local
	s2f64 = l3
	// get_local
	s3f64 = l3
	// get_local
	s4f64 = l3
	// get_local
	s5f64 = l3
	// get_local
	s6f64 = l3
	// const
	s7f64 = 3.479331075960212e-05
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7f64 = 0.0007915349942898145
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = -0.04005553450067941
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = 0.20121253213486293
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = -0.3255658186224009
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = 0.16666666666666666
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l3
	// get_local
	s3f64 = l3
	// get_local
	s4f64 = l3
	// get_local
	s5f64 = l3
	// const
	s6f64 = 0.07703815055590194
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = -0.6882839716054533
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = 2.0209457602335057
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = -2.403394911734414
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = 1
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.div
	s1f64 = s1f64 / s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl3:
	// const
	s0f64 = 1
	// get_local
	s1f64 = l0
	// unary: f64.abs
	s1f64 = math.Abs(s1f64)
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 0.5
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l0 = s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l0
	// get_local
	s5f64 = l0
	// const
	s6f64 = 3.479331075960212e-05
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = 0.0007915349942898145
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = -0.04005553450067941
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = 0.20121253213486293
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = -0.3255658186224009
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 0.16666666666666666
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l0
	// const
	s5f64 = 0.07703815055590194
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = -0.6882839716054533
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = 2.0209457602335057
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = -2.403394911734414
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 1
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l4 = s0f64
	// get_local
	s0f64 = l0
	// unary: f64.sqrt
	s0f64 = math.Sqrt(s0f64)
	// set_local
	l3 = s0f64
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1072640819
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
	// const
	s0f64 = 1.5707963267948966
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l4
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l3
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// get_local
	s2f64 = l0
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// const
	s2f64 = -6.123233995736766e-17
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l0 = s0f64
	// br
	goto lbl4
	// end_block
lbl5:
	// const
	s0f64 = 0.7853981633974483
	// get_local
	s1f64 = l3
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// const
	s2i64 = -4294967296
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// tee_local
	l5 = s1f64
	// get_local
	s2f64 = l5
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l3
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l4
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 6.123233995736766e-17
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l5
	// get_local
	s5f64 = l5
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// get_local
	s4f64 = l3
	// get_local
	s5f64 = l5
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.div
	s3f64 = s3f64 / s4f64
	// tee_local
	l0 = s3f64
	// get_local
	s4f64 = l0
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 0.7853981633974483
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl4:
	// get_local
	s0f64 = l0
	// unary: f64.neg
	s0f64 = -s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2i64 = l1
	// const
	s3i64 = 0
	// binary: i64.lt_s
	if s2i64 < s3i64 {
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
	l0 = s0f64
	// end_block
lbl2:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
