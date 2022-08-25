package internals

import (
	"encoding/binary"
)

func f573(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
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
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 128
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l4 = s0i32
	// set_global
	ctx.G0 = s0i32
	// const
	s0i32 = -1
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// const
	s2i32 = -1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l1
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+116):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l4
	// const
	s3i32 = 126
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l1
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l0 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+112):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 0
	// const
	s2i32 = 112
	// call
	s0i32 = f578(ctx, s0i32, s1i32, s2i32)
	// tee_local
	l4 = s0i32
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+64):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 208
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// const
	s2i32 = 112
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+68):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// const
	s2i32 = 127
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 61
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l3
	// call
	s0i32 = f568(ctx, s0i32, s1i32, s2i32)
	// set_local
	l5 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 128
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l5
	// return
	return s0i32
}
