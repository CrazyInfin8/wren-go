package internals

import (
	"encoding/binary"
)

func f520(ctx *Context, l0 int64) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int64
	_ = l4
	var l5 int64
	_ = l5
	var l6 int64
	_ = l6
	var l7 int64
	_ = l7
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 int64
	_ = l18
	var l19 int64
	_ = l19
	var l20 int64
	_ = l20
	var l21 int64
	_ = l21
	var l22 int64
	_ = l22
	var l23 int64
	_ = l23
	var l24 int64
	_ = l24
	var l25 int64
	_ = l25
	var l26 int64
	_ = l26
	var l27 int64
	_ = l27
	var l28 int64
	_ = l28
	var l29 int64
	_ = l29
	var l30 int64
	_ = l30
	var l31 int64
	_ = l31
	var l32 int64
	_ = l32
	var l33 int64
	_ = l33
	var l34 int64
	_ = l34
	var l35 int64
	_ = l35
	var l36 int64
	_ = l36
	var l37 int32
	_ = l37
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l1 = s0i32
	// const
	s0i32 = 16
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l4 = s0i64
	// const
	s0i64 = -1
	// set_local
	l5 = s0i64
	// get_local
	s0i64 = l4
	// get_local
	s1i64 = l5
	// binary: i64.xor
	s0i64 = s0i64 ^ s1i64
	// set_local
	l6 = s0i64
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l7 = s0i64
	// const
	s0i64 = 18
	// set_local
	l8 = s0i64
	// get_local
	s0i64 = l7
	// get_local
	s1i64 = l8
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l9 = s0i64
	// get_local
	s0i64 = l6
	// get_local
	s1i64 = l9
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l10 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l10
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l11 = s0i64
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l12 = s0i64
	// const
	s0i64 = 31
	// set_local
	l13 = s0i64
	// get_local
	s0i64 = l12
	// get_local
	s1i64 = l13
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l14 = s0i64
	// get_local
	s0i64 = l11
	// get_local
	s1i64 = l14
	// binary: i64.xor
	s0i64 = s0i64 ^ s1i64
	// set_local
	l15 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l15
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l16 = s0i64
	// const
	s0i64 = 21
	// set_local
	l17 = s0i64
	// get_local
	s0i64 = l16
	// get_local
	s1i64 = l17
	// binary: i64.mul
	s0i64 = s0i64 * s1i64
	// set_local
	l18 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l18
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l19 = s0i64
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l20 = s0i64
	// const
	s0i64 = 11
	// set_local
	l21 = s0i64
	// get_local
	s0i64 = l20
	// get_local
	s1i64 = l21
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l22 = s0i64
	// get_local
	s0i64 = l19
	// get_local
	s1i64 = l22
	// binary: i64.xor
	s0i64 = s0i64 ^ s1i64
	// set_local
	l23 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l23
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l24 = s0i64
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l25 = s0i64
	// const
	s0i64 = 6
	// set_local
	l26 = s0i64
	// get_local
	s0i64 = l25
	// get_local
	s1i64 = l26
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l27 = s0i64
	// get_local
	s0i64 = l24
	// get_local
	s1i64 = l27
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l28 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l28
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l29 = s0i64
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l30 = s0i64
	// const
	s0i64 = 22
	// set_local
	l31 = s0i64
	// get_local
	s0i64 = l30
	// get_local
	s1i64 = l31
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l32 = s0i64
	// get_local
	s0i64 = l29
	// get_local
	s1i64 = l32
	// binary: i64.xor
	s0i64 = s0i64 ^ s1i64
	// set_local
	l33 = s0i64
	// get_local
	s0i32 = l3
	// get_local
	s1i64 = l33
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l34 = s0i64
	// const
	s0i64 = 1073741823
	// set_local
	l35 = s0i64
	// get_local
	s0i64 = l34
	// get_local
	s1i64 = l35
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l36 = s0i64
	// get_local
	s0i64 = l36
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l37
	// return
	return s0i32
}
