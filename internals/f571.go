package internals

func f571(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
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
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 256
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 73728
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l3
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l2 = s2i32
	// const
	s3i32 = 256
	// get_local
	s4i32 = l2
	// const
	s5i32 = 256
	// binary: i32.lt_u
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	// tee_local
	l4 = s4i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l4
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// loop
lbl2:
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 256
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl3:
	// get_local
	s0i32 = l2
	// const
	s1i32 = -256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// end
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l0
	// call
	s0i32 = f564(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl0:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 256
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
}
