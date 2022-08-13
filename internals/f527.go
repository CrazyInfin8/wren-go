package internals

import (
	"encoding/binary"
)

func f527(ctx *Context) int32 {
	var l0 int32
	_ = l0
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
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
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_global
	s0i32 = ctx.G0
	// set_local
	l0 = s0i32
	// const
	s0i32 = 48
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l2
	// set_global
	ctx.G0 = s0i32
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// call
	f408(ctx, s0i32)
	// const
	s0i32 = 194
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 195
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// const
	s0i32 = 196
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// const
	s0i32 = 197
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// const
	s0i32 = 198
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// const
	s0i32 = 199
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l11
	// call
	s0i32 = f410(ctx, s0i32)
	// set_local
	l12 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l13
	// get_local
	s1i32 = l12
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41624):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l14 = s0i32
	// const
	s0i32 = 48
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l2
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
	// get_local
	s0i32 = l14
	// return
	return s0i32
}
