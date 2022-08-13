package internals

import (
	"encoding/binary"
)

func f637(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+44):]))
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l4
	// const
	s3i32 = 0
	// binary: i32.ne
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l6 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// block
	// block
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l3
	// const
	s2i32 = 2
	// call
	s0i32 = f635(ctx, s0i32, s1i32, s2i32)
	// set_local
	l4 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+56):]))
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l4
	// call
	s0i32 = f636(ctx, s0i32, s1i32, s2i32)
	// set_local
	l4 = s0i32
	// end_block
lbl0:
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// block
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 0
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 32
	// const
	s3i32 = 16
	// get_local
	s4i32 = l4
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl2
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// tee_local
	l5 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l4
	// set_local
	l6 = s0i32
	// br
	goto lbl2
	// end_block
lbl4:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// tee_local
	l6 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l4
	// get_local
	s3i32 = l5
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l6
	// const
	s2i32 = 1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// load: i32.load8_u
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// end_block
lbl5:
	// get_local
	s0i32 = l2
	// set_local
	l6 = s0i32
	// end_block
lbl2:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l6
	// return
	return s0i32
}
