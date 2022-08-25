package internals

func f16(ctx *Context) {
	var l0 int32
	_ = l0
	var s0i32 int32
	_ = s0i32
	// block
	// call
	s0i32 = f18(ctx)
	// tee_local
	l0 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// call
	f551(ctx, s0i32)
	// unreachable
	panic("unreachable executed")
	// end_block
lbl0:
}
