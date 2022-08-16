package internals

func f623(ctx *Context, l0 float64) float64 {
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l0 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// return
	return s0f64
}
