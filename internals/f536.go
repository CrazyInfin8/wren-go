package internals

func f536(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	// block
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// unary: i64.extend_u/i32
	s0i64 = int64(uint32(s0i32))
	// get_local
	s1i32 = l1
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.mul
	s0i64 = s0i64 * s1i64
	// tee_local
	l3 = s0i64
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l0
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 65536
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = -1
	// get_local
	s1i32 = l2
	// get_local
	s2i64 = l3
	// const
	s3i64 = 32
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// const
	s3i32 = 0
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
	// set_local
	l2 = s0i32
	// end_block
lbl0:
	// block
	// get_local
	s0i32 = l2
	// call
	s0i32 = f533(ctx, s0i32)
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
		goto lbl2
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// get_local
	s2i32 = l2
	// call
	s0i32 = f580(ctx, s0i32, s1i32, s2i32)
	// end_block
lbl2:
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
