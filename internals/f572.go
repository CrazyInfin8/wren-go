package internals

import (
	"encoding/binary"
)

func f572(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 208
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+204):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 160
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 184
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 176
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+168):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+160):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+200):], uint32(s1i32))
	// block
	// block
	// const
	s0i32 = 0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l3
	// const
	s3i32 = 200
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 80
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// get_local
	s4i32 = l3
	// const
	s5i32 = 160
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// call
	s0i32 = f573(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
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
		goto lbl1
	}
	// const
	s0i32 = -1
	// set_local
	l0 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l4 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+60):]))
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
		goto lbl2
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// const
	s2i32 = -33
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl2:
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 80
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
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
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// br
	goto lbl5
	// end_block
lbl6:
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// end_block
lbl5:
	// const
	s0i32 = -1
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l0
	// call
	s0i32 = f567(ctx, s0i32)
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// end_block
lbl4:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l3
	// const
	s3i32 = 200
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 80
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// get_local
	s4i32 = l3
	// const
	s5i32 = 160
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// call
	s0i32 = f573(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// set_local
	l2 = s0i32
	// end_block
lbl3:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l5
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// const
	s2i32 = 0
	// get_local
	s3i32 = l0
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
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1
	// get_local
	s2i32 = l5
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l2 = s0i32
	// end_block
lbl7:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l5 = s1i32
	// get_local
	s2i32 = l1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = -1
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l5
	// const
	s3i32 = 32
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l0 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 208
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
