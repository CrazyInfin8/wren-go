package internals

import (
	"math"
)

func f605(ctx *Context, l0 float64) float64 {
	var l1 int64
	_ = l1
	var l2 int32
	_ = l2
	var l3 float64
	_ = l3
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
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l1 = s0i64
	// const
	s1i64 = 52
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 2047
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1074
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
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
	s1i32 = 1021
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
	// get_local
	s0f64 = l0
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl1:
	// block
	// block
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2i64 = l1
	// const
	s3i64 = -1
	// binary: i64.gt_s
	if s2i64 > s3i64 {
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
	// tee_local
	l0 = s0f64
	// const
	s1f64 = 4.503599627370496e+15
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = -4.503599627370496e+15
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l0
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l3 = s0f64
	// const
	s1f64 = 0.5
	// binary: f64.gt
	if s0f64 > s1f64 {
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
		goto lbl3
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l3
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = -1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// br
	goto lbl2
	// end_block
lbl3:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l3
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// get_local
	s0f64 = l3
	// const
	s1f64 = -0.5
	// binary: f64.le
	if s0f64 <= s1f64 {
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
		goto lbl2
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl2:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// unary: f64.neg
	s1f64 = -s1f64
	// get_local
	s2i64 = l1
	// const
	s3i64 = -1
	// binary: i64.gt_s
	if s2i64 > s3i64 {
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
lbl0:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
