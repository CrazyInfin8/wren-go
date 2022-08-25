package internals

func f653(ctx *Context, l0 int32, l1 int32) float64 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f64 float64
	_ = s0f64
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// call
	s0f64 = f446(ctx, s0i32, s1i32)
	// call
	f550(ctx)
	// return
	return s0f64
}
