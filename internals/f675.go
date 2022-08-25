package internals

func f675(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// call
	s0i32 = f470(ctx, s0i32, s1i32, s2i32)
	// call
	f550(ctx)
	// return
	return s0i32
}
