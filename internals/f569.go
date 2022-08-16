package internals

import (
	"encoding/binary"
)

func f569(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l4 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l3
	// call
	s0i32 = f567(ctx, s0i32)
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l5 = s0i32
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+20):]))
	// tee_local
	l6 = s1i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l4
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l4
	// get_local
	s3i32 = l3
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+32):]))
	// call_indirect
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].f == nil {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = table[s3i32].f.(func(*Context, int32, int32, int32) int32)(ctx, s0i32, s1i32, s2i32)
	// set_local
	l5 = s0i32
	// br
	goto lbl0
	// end_block
lbl2:
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+64):]))
	// const
	s1i32 = 0
	// binary: i32.ge_s
	if s0i32 >= s1i32 {
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
	l5 = s0i32
	// br
	goto lbl3
	// end_block
lbl4:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l8 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// loop
lbl5:
	// block
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l4
	// set_local
	l5 = s0i32
	// br
	goto lbl3
	// end_block
lbl6:
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l10 = s0i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l9
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 10
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// end
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l4
	// get_local
	s3i32 = l10
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 1
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// tee_local
	l7 = s2i32
	// get_local
	s3i32 = l3
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+32):]))
	// call_indirect
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].f == nil {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = table[s3i32].f.(func(*Context, int32, int32, int32) int32)(ctx, s0i32, s1i32, s2i32)
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l7
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l10
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l10
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l6 = s0i32
	// end_block
lbl3:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l5
	// call
	s0i32 = f581(ctx, s0i32, s1i32, s2i32)
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+20):]))
	// get_local
	s2i32 = l5
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl0:
	// block
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// get_local
	s2i32 = l1
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// return
	return s0i32
	// end_block
lbl7:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// binary: i32.div_u
	s0i32 = i32DivU(s0i32, s1i32)
	// return
	return s0i32
}
