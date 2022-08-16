package internals

func f549(ctx *Context, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l1
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l3
	// call
	s0i32 = f12(ctx, s0i32, s1i64, s2i32, s3i32)
	// const
	s1i32 = 65535
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// return
	return s0i32
}
