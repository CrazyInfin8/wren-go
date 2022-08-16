package internals

import (
	"math"
)

func f622(ctx *Context, l0 int32) float64 {
	var s2i32 int32
	_ = s2i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// const
	s0f64 = math.Inf(0)
	// const
	s1f64 = math.Inf(0)
	// get_local
	s2i32 = l0
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// return
	return s0f64
}
