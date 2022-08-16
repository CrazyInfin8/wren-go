package internals

func f587(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// get_local
	s2i32 = l1
	// call
	s0i32 = f586(ctx, s0i32, s1i32, s2i32)
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// return
	return s0i32
}
