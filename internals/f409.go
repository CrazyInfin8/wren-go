package internals

import (
	"encoding/binary"
)

func f409(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// call
	f534(ctx, s0i32)
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l10
	// call
	s0i32 = f537(ctx, s0i32, s1i32)
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l12 = s0i32
	// const
	s0i32 = 16
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l13
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l14
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l12
	// return
	return s0i32
}