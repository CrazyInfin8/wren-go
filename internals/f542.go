package internals

import (
	"math"
)

func f542(ctx *Context, l0 float64, l1 float64) float64 {
	var s0i32 int32
	_ = s0i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// block
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.ne
	if s0f64 != s1f64 {
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
	s0f64 = l1
	// get_local
	s1f64 = l1
	// binary: f64.eq
	if s0f64 == s1f64 {
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
	// return
	return s0f64
	// end_block
lbl1:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l1
	// binary: f64.min
	s0f64 = math.Min(s0f64, s1f64)
	// set_local
	l1 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l1
	// return
	return s0f64
}
