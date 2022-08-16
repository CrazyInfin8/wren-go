package internals

func f602(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 0
	// return
	return s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 0
	// call
	s0i32 = f601(ctx, s0i32, s1i32, s2i32)
	// return
	return s0i32
}
