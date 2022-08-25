package internals

func f606(ctx *Context, l0 float64, l1 float64) float64 {
	var l2 float64
	_ = l2
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
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
	// const
	s0f64 = 1
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l2 = s1f64
	// const
	s2f64 = 0.5
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l3 = s1f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l4 = s0f64
	// const
	s1f64 = 1
	// get_local
	s2f64 = l4
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l3
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l2
	// get_local
	s3f64 = l2
	// get_local
	s4f64 = l2
	// get_local
	s5f64 = l2
	// const
	s6f64 = 2.480158728947673e-05
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = -0.001388888888887411
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = 0.0416666666666666
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// get_local
	s4f64 = l2
	// get_local
	s5f64 = l2
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// tee_local
	l3 = s4f64
	// get_local
	s5f64 = l3
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// get_local
	s5f64 = l2
	// get_local
	s6f64 = l2
	// const
	s7f64 = -1.1359647557788195e-11
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7f64 = 2.087572321298175e-09
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = -2.7557314351390663e-07
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l1
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
}
