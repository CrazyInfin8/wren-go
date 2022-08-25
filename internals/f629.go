package internals

import (
	"math"
)

func f629(ctx *Context, l0 float64, l1 float64, l2 int32) float64 {
	var l3 int64
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
	var l9 float64
	_ = l9
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
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
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
	var s8f64 float64
	_ = s8f64
	var s9f64 float64
	_ = s9f64
	var s10f64 float64
	_ = s10f64
	var s11f64 float64
	_ = s11f64
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
	s1i32 = 2147483640
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 1072010280
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// tee_local
	l5 = s0i32
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
	// br
	goto lbl0
	// end_block
lbl1:
	// const
	s0f64 = 0.7853981633974483
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// unary: f64.neg
	s2f64 = -s2f64
	// get_local
	s3i64 = l3
	// const
	s4i64 = -1
	// binary: i64.gt_s
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// tee_local
	l6 = s3i32
	// select
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// const
	s1f64 = 3.061616997868383e-17
	// get_local
	s2f64 = l1
	// get_local
	s3f64 = l1
	// unary: f64.neg
	s3f64 = -s3f64
	// get_local
	s4i32 = l6
	// select
	if s4i32 != 0 {
		// s2f64 = s2f64
	} else {
		s2f64 = s3f64
	}
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l4
	// const
	s1i32 = 31
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l4 = s0i32
	// const
	s0f64 = 0
	// set_local
	l1 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l0
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l7 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l8 = s1f64
	// const
	s2f64 = 0.3333333333333341
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l7
	// get_local
	s3f64 = l8
	// get_local
	s4f64 = l7
	// get_local
	s5f64 = l7
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// tee_local
	l9 = s4f64
	// get_local
	s5f64 = l9
	// get_local
	s6f64 = l9
	// get_local
	s7f64 = l9
	// get_local
	s8f64 = l9
	// const
	s9f64 = -1.8558637485527546e-05
	// binary: f64.mul
	s8f64 = s8f64 * s9f64
	// const
	s9f64 = 7.817944429395571e-05
	// binary: f64.add
	s8f64 = s8f64 + s9f64
	// binary: f64.mul
	s7f64 = s7f64 * s8f64
	// const
	s8f64 = 0.0005880412408202641
	// binary: f64.add
	s7f64 = s7f64 + s8f64
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7f64 = 0.0035920791075913124
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = 0.021869488294859542
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = 0.13333333333320124
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// get_local
	s5f64 = l7
	// get_local
	s6f64 = l9
	// get_local
	s7f64 = l9
	// get_local
	s8f64 = l9
	// get_local
	s9f64 = l9
	// get_local
	s10f64 = l9
	// const
	s11f64 = 2.590730518636337e-05
	// binary: f64.mul
	s10f64 = s10f64 * s11f64
	// const
	s11f64 = 7.140724913826082e-05
	// binary: f64.add
	s10f64 = s10f64 + s11f64
	// binary: f64.mul
	s9f64 = s9f64 * s10f64
	// const
	s10f64 = 0.0002464631348184699
	// binary: f64.add
	s9f64 = s9f64 + s10f64
	// binary: f64.mul
	s8f64 = s8f64 * s9f64
	// const
	s9f64 = 0.0014562094543252903
	// binary: f64.add
	s8f64 = s8f64 + s9f64
	// binary: f64.mul
	s7f64 = s7f64 * s8f64
	// const
	s8f64 = 0.0088632398235993
	// binary: f64.add
	s7f64 = s7f64 + s8f64
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7f64 = 0.05396825397622605
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// get_local
	s4f64 = l1
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l1
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l7 = s1f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l9 = s0f64
	// block
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// const
	s0i32 = 1
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// tee_local
	l1 = s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l7
	// get_local
	s3f64 = l9
	// get_local
	s4f64 = l9
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// get_local
	s4f64 = l9
	// get_local
	s5f64 = l1
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.div
	s3f64 = s3f64 / s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l9 = s1f64
	// get_local
	s2f64 = l9
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l9 = s0f64
	// unary: f64.neg
	s0f64 = -s0f64
	// get_local
	s1f64 = l9
	// get_local
	s2i32 = l4
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// return
	return s0f64
	// end_block
lbl2:
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
		goto lbl3
	}
	// const
	s0f64 = -1
	// get_local
	s1f64 = l9
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// tee_local
	l1 = s0f64
	// get_local
	s1f64 = l1
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// const
	s2i64 = -4294967296
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// tee_local
	l1 = s1f64
	// get_local
	s2f64 = l7
	// get_local
	s3f64 = l9
	// unary: i64.reinterpret/f64
	s3i64 = int64(math.Float64bits(s3f64))
	// const
	s4i64 = -4294967296
	// binary: i64.and
	s3i64 = s3i64 & s4i64
	// unary: f64.reinterpret/i64
	s3f64 = math.Float64frombits(uint64(s3i64))
	// tee_local
	l9 = s3f64
	// get_local
	s4f64 = l0
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l1
	// get_local
	s3f64 = l9
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = 1
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l9 = s0f64
	// end_block
lbl3:
	// get_local
	s0f64 = l9
	// return
	return s0f64
}
