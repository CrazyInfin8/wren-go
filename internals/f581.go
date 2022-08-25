package internals

func f581(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// block
	// get_local
	s0i32 = l2
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
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// tee_local
	l4 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// loop
lbl2:
	// block
	// get_local
	s0i32 = l1
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l4
	// set_local
	l3 = s0i32
	// br
	goto lbl1
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l2
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l4
	// set_local
	l3 = s0i32
	// br
	goto lbl1
	// end_block
lbl4:
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l5
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l4
	// set_local
	l3 = s0i32
	// br
	goto lbl1
	// end_block
lbl5:
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l4
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// end
	// end_block
lbl1:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l1
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// return
	return s0i32
}
