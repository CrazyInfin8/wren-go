package internals

import (
	"encoding/binary"
)

func f462(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l11 int32
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int64
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
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 int32
	_ = l30
	var l31 int32
	_ = l31
	var l32 int32
	_ = l32
	var l33 int32
	_ = l33
	var l34 int32
	_ = l34
	var l35 int64
	_ = l35
	var l36 int32
	_ = l36
	var l37 int32
	_ = l37
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
	l4 = s0i32
	// const
	s0i32 = 48
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l6
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 7378
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l9
	// get_local
	s2i32 = l6
	// call
	s0i64 = f314(ctx, s0i32, s1i32, s2i32)
	// set_local
	l10 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l10
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+24):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+24):]))
	// set_local
	l12 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l13 = s0i64
	// get_local
	s0i64 = l12
	// get_local
	s1i64 = l13
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l14 = s0i64
	// get_local
	s0i64 = l14
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l15
	// call
	f131(ctx, s0i32, s1i32)
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+24):]))
	// set_local
	l17 = s0i64
	// get_local
	s0i32 = l16
	// get_local
	s1i64 = l17
	// call
	s0i32 = f431(ctx, s0i32, s1i64)
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l18
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l19
	// call
	f133(ctx, s0i32)
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l20 = s0i32
	// const
	s0i32 = 28
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l24
	// call
	s0i32 = f585(ctx, s0i32)
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l22
	// get_local
	s1i32 = l23
	// get_local
	s2i32 = l25
	// call
	s0i32 = f85(ctx, s0i32, s1i32, s2i32)
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l26
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l29
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l31 = s0i32
	// const
	s0i32 = 3
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l31
	// get_local
	s1i32 = l32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l30
	// get_local
	s1i32 = l33
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l34
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l35 = s0i64
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l28
	// get_local
	s2i64 = l35
	// call
	f444(ctx, s0i32, s1i32, s2i64)
	// const
	s0i32 = 48
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l36
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l37
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
