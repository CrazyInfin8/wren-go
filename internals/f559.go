package internals

import (
	"encoding/binary"
)

func f559(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
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
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+24):]))
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+20):]))
	// get_local
	s2i32 = l1
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 2
	// set_local
	l4 = s0i32
	// block
	// block
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+56):]))
	// get_local
	s2i32 = l3
	// const
	s3i32 = 2
	// call
	s1i32 = f558(ctx, s1i32, s2i32, s3i32)
	// tee_local
	l6 = s1i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l3
	// set_local
	l1 = s0i32
	// loop
lbl2:
	// block
	// get_local
	s0i32 = l6
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
		goto lbl3
	}
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 2
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l6 = s0i32
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l1
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// tee_local
	l7 = s2i32
	// binary: i32.gt_u
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l9 = s0i32
	// get_local
	s1i32 = l9
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l6
	// get_local
	s3i32 = l7
	// const
	s4i32 = 0
	// get_local
	s5i32 = l8
	// select
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// tee_local
	l7 = s2i32
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 12
	// const
	s2i32 = 4
	// get_local
	s3i32 = l8
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l9 = s0i32
	// get_local
	s1i32 = l9
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l7
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+56):]))
	// get_local
	s2i32 = l1
	// const
	s3i32 = 8
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l1
	// get_local
	s4i32 = l8
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// tee_local
	l1 = s2i32
	// get_local
	s3i32 = l4
	// get_local
	s4i32 = l8
	// binary: i32.sub
	s3i32 = s3i32 - s4i32
	// tee_local
	l4 = s3i32
	// call
	s1i32 = f558(ctx, s1i32, s2i32, s3i32)
	// tee_local
	l6 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// end
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+40):]))
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+44):]))
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l6 = s0i32
	// end_block
lbl0:
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
