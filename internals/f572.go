package internals

func f572(ctx *Context) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// const
	s0i32 = 20644
	// const
	s1i32 = 41288
	// call
	s0i32 = f567(ctx, s0i32, s1i32)
	// call
	f542(ctx)
	// unreachable
	panic("unreachable executed")
}