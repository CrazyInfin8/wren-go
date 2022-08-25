package internals

func f567(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	// get_local
	s0i32 = l0
	// call
	s0i32 = f580(ctx, s0i32)
	// set_local
	l2 = s0i32
	// const
	s0i32 = -1
	// const
	s1i32 = 0
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l0
	// const
	s4i32 = 1
	// get_local
	s5i32 = l2
	// get_local
	s6i32 = l1
	// call
	s3i32 = f565(ctx, s3i32, s4i32, s5i32, s6i32)
	// binary: i32.ne
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// return
	return s0i32
}
