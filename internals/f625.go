package internals

import (
	"encoding/binary"
	"math"
)

func f625(ctx *Context, l0 float64, l1 int32) float64 {
	var l2 int64
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
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
	s1i32 = -1023
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1075
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 1024
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
	// get_local
	s0i64 = l2
	// const
	s1i64 = 4503599627370495
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 0
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// end_block
lbl2:
	// get_local
	s0i64 = l2
	// const
	s1i64 = -9223372036854775808
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// return
	return s0f64
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1022
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i64 = l2
	// const
	s2i64 = -9223372036854775808
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0f64 = l0
	// return
	return s0f64
	// end_block
lbl3:
	// block
	// get_local
	s0i64 = l2
	// get_local
	s1i32 = l4
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// tee_local
	l5 = s1i64
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// const
	s1i64 = 4503599627370495
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 0
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l1
	// get_local
	s1f64 = l0
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i64 = l2
	// const
	s1i64 = -9223372036854775808
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// return
	return s0f64
	// end_block
lbl4:
	// get_local
	s0i32 = l1
	// const
	s1i64 = -4503599627370496
	// get_local
	s2i64 = l5
	// binary: i64.shr_s
	s1i64 = s1i64 >> (uint64(s2i64) & 63)
	// get_local
	s2i64 = l2
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// tee_local
	l2 = s1i64
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0f64 = l0
	// get_local
	s1i64 = l2
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
