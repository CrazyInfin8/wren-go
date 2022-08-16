package internals

import (
	"math"
)

func f619(ctx *Context, l0 float64, l1 int64, l2 int64) float64 {
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
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
	// block
	// get_local
	s0i64 = l2
	// const
	s1i64 = 2147483648
	// binary: i64.and
	s0i64 = s0i64 & s1i64
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
		goto lbl0
	}
	// get_local
	s0i64 = l1
	// const
	s1i64 = -4544132024016830464
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l3 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l3
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = 5.486124068793689e+303
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// block
	// get_local
	s0i64 = l1
	// const
	s1i64 = 4602678819172646912
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l3 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l4 = s0f64
	// get_local
	s1f64 = l3
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l0 = s0f64
	// const
	s1f64 = 1
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
		goto lbl1
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l4
	// get_local
	s2f64 = l3
	// get_local
	s3f64 = l0
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l0
	// const
	s3f64 = 1
	// get_local
	s4f64 = l5
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = -1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl1:
	// get_local
	s0f64 = l0
	// const
	s1f64 = 2.2250738585072014e-308
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
}
