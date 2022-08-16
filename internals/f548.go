package internals

func f548(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// call
	s0i32 = f11(ctx, s0i32, s1i32)
	// const
	s1i32 = 65535
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// return
	return s0i32
}
