package internals

func f611(ctx *Context, l0 int32, l1 float64) float64 {
	var s2i32 int32
	_ = s2i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// get_local
	s0f64 = l1
	// unary: f64.neg
	s0f64 = -s0f64
	// get_local
	s1f64 = l1
	// get_local
	s2i32 = l0
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// get_local
	s1f64 = l1
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
}
