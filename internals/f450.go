package internals

import (
	"encoding/binary"
)

func f450(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int64
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
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
	// get_global
	s0i32 = ctx.G0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 16
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l5
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// store: i32.store8
	ctx.Mem[int(s0i32+7)] = uint8(s1i32)
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+7)])
	// set_local
	l9 = s0i32
	// const
	s0i64 = 9222246136947933187
	// set_local
	l10 = s0i64
	// const
	s0i64 = 9222246136947933186
	// set_local
	l11 = s0i64
	// const
	s0i32 = 1
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l12
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l13 = s0i32
	// get_local
	s0i64 = l10
	// get_local
	s1i64 = l11
	// get_local
	s2i32 = l13
	// select
	if s2i32 != 0 {
		// s0i64 = s0i64
	} else {
		s0i64 = s1i64
	}
	// set_local
	l14 = s0i64
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// get_local
	s2i64 = l14
	// call
	f451(ctx, s0i32, s1i32, s2i64)
	// const
	s0i32 = 16
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l15
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
