package internals

func f543(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_local
	s0i32 = l0
	// call
	s0i32 = f11(ctx, s0i32)
	// const
	s1i32 = 65535
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// return
	return s0i32
}
