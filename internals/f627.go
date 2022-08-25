package internals

import (
	"math"
)

func f627(ctx *Context, l0 float64, l1 int32) float64 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1024
	// binary: i32.lt_s
	if s0i32 < s1i32 {
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
	s1f64 = 8.98846567431158e+307
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l0 = s0f64
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2047
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1023
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// br
	goto lbl0
	// end_block
lbl2:
	// get_local
	s0f64 = l0
	// const
	s1f64 = 8.98846567431158e+307
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l1
	// const
	s1i32 = 3069
	// get_local
	s2i32 = l1
	// const
	s3i32 = 3069
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
	// const
	s1i32 = -2046
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1023
	// binary: i32.gt_s
	if s0i32 > s1i32 {
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
	// const
	s1f64 = 2.004168360008973e-292
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l0 = s0f64
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1992
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 969
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0f64 = l0
	// const
	s1f64 = 2.004168360008973e-292
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l1
	// const
	s1i32 = -2960
	// get_local
	s2i32 = l1
	// const
	s3i32 = -2960
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
	// const
	s1i32 = 1938
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1023
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// const
	s2i64 = 52
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
}
