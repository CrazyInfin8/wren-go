package internals

func f628(ctx *Context, l0 int64) int32 {
	var l1 int32
	_ = l1
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i64 = l0
	// const
	s1i64 = 52
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 2047
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1023
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
	s0i32 = 2
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1075
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// const
	s0i64 = 1
	// const
	s1i32 = 1075
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l3 = s0i64
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// get_local
	s1i64 = l0
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 0
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 2
	// const
	s1i32 = 1
	// get_local
	s2i64 = l3
	// get_local
	s3i64 = l0
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// unary: i64.eqz
	if s2i64 == 0 {
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
	l1 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l1
	// return
	return s0i32
}
