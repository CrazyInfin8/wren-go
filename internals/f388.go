package internals

import (
	"encoding/binary"
)

func f388(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s0i32 = l5
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l9 = s0i32
	// const
	s0i32 = 16
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l10
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// const
	s0i32 = 0
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l12
	// get_local
	s2i32 = l12
	// get_local
	s3i32 = l11
	// call
	s0i32 = f303(ctx, s0i32, s1i32, s2i32, s3i32)
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l13
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l16 = s0i32
	// const
	s0i32 = 4
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l15
	// get_local
	s2i32 = l17
	// get_local
	s3i32 = l16
	// call
	f304(ctx, s0i32, s1i32, s2i32, s3i32)
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l18 = s0i32
	// const
	s0i32 = 16
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l18
	// get_local
	s1i32 = l19
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l21 = s0i32
	// const
	s0i32 = 0
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l22
	// get_local
	s2i32 = l21
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l23 = s0i32
	// const
	s0i32 = 16
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l25
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l23
	// return
	return s0i32
}
