package internals

func f628(ctx *Context, l0 float64, l1 float64, l2 int32) float64 {
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var s0i32 int32
	_ = s0i32
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
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l3 = s0f64
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l3
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l3
	// const
	s2f64 = 1.58969099521155e-10
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = -2.5050760253406863e-08
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l3
	// const
	s3f64 = 2.7557313707070068e-06
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = -0.0001984126982985795
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 0.00833333333332249
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l4 = s0f64
	// get_local
	s0f64 = l3
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l5 = s0f64
	// block
	// get_local
	s0i32 = l2
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0f64 = l5
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l4
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = -0.16666666666666632
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l3
	// get_local
	s2f64 = l1
	// const
	s3f64 = 0.5
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l5
	// get_local
	s4f64 = l4
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l1
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l5
	// const
	s3f64 = 0.16666666666666632
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// return
	return s0f64
}
