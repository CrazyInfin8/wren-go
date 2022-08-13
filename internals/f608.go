package internals

import (
	"math"
)

func f608(ctx *Context, l0 float64) float64 {
	var l1 int64
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
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
	s1i32 = 2146435072
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
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// const
	s0i32 = 715094163
	// set_local
	l3 = s0i32
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1048575
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// const
	s0i32 = 696219795
	// set_local
	l3 = s0i32
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1.8014398509481984e+16
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
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
	// end_block
lbl2:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 3
	// binary: i32.div_u
	s0i32 = i32DivU(s0i32, s1i32)
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// const
	s1i64 = 32
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// get_local
	s1i64 = l1
	// const
	s2i64 = -9223372036854775808
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l4 = s0f64
	// get_local
	s1f64 = l4
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l4
	// get_local
	s2f64 = l0
	// binary: f64.div
	s1f64 = s1f64 / s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l5
	// get_local
	s2f64 = l5
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l5
	// const
	s2f64 = 0.14599619288661245
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = -0.758397934778766
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l5
	// get_local
	s2f64 = l5
	// const
	s3f64 = 1.6214297201053545
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = -1.8849797954337717
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 1.87595182427177
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l4
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// const
	s1i64 = 2147483648
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = -1073741824
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l5
	// get_local
	s3f64 = l5
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.div
	s1f64 = s1f64 / s2f64
	// tee_local
	l0 = s1f64
	// get_local
	s2f64 = l5
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l5
	// get_local
	s3f64 = l5
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// get_local
	s3f64 = l0
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.div
	s1f64 = s1f64 / s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l5
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl1:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
