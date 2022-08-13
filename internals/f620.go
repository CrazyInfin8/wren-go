package internals

import (
	"encoding/binary"
	"math"
)

func f620(ctx *Context, l0 float64, l1 int32) float64 {
	var l2 int64
	_ = l2
	var l3 int32
	_ = l3
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l2 = s0i64
	// const
	s1i64 = 52
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 2047
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 2047
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
	// block
	// get_local
	s0i32 = l3
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// block
	// get_local
	s0f64 = l0
	// const
	s1f64 = 0
	// binary: f64.ne
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0f64 = l0
	// return
	return s0f64
	// end_block
lbl2:
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1.8446744073709552e+19
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1i32 = l1
	// call
	s0f64 = f620(ctx, s0f64, s1i32)
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l1
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = -64
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0f64 = l0
	// return
	return s0f64
	// end_block
lbl1:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l3
	// const
	s2i32 = -1022
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i64 = l2
	// const
	s1i64 = -9218868437227405313
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 4602678819172646912
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// set_local
	l0 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
