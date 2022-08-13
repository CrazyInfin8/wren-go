package internals

func f597(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// const
	s2i32 = 31
	// binary: i32.shr_s
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	// tee_local
	l1 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// return
	return s0i32
}
