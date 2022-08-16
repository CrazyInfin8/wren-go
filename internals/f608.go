package internals

import (
	"encoding/binary"
	"math"
)

func f608(ctx *Context, l0 float64, l1 float64) float64 {
	var l2 int64
	_ = l2
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
	var l8 float64
	_ = l8
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	// block
	// get_local
	s0f64 = l1
	// get_local
	s1f64 = l1
	// binary: f64.eq
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// binary: f64.eq
	if s1f64 == s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// block
	// get_local
	s0f64 = l1
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l2 = s0i64
	// const
	s1i64 = 32
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l3 = s0i32
	// const
	s1i32 = -1072693248
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i64 = l2
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0f64 = l0
	// call
	s0f64 = f607(ctx, s0f64)
	// return
	return s0f64
	// end_block
lbl1:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 30
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// const
	s1i32 = 2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1f64 = l0
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// tee_local
	l2 = s1i64
	// const
	s2i64 = 32
	// binary: i64.shr_u
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// tee_local
	l6 = s1i32
	// const
	s2i32 = 31
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i64 = l2
	// unary: i32.wrap/i64
	s1i32 = int32(s1i64)
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0f64 = l0
	// set_local
	l8 = s0f64
	// block
	// block
	// get_local
	s0i32 = l7
	// br_table
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl2
	case 2:
		goto lbl5
	case 3:
		goto lbl4
	default:
		goto lbl2
	}
	// end_block
lbl5:
	// const
	s0f64 = 3.141592653589793
	// return
	return s0f64
	// end_block
lbl4:
	// const
	s0f64 = -3.141592653589793
	// return
	return s0f64
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// const
	s0f64 = 1.5707963267948966
	// get_local
	s1f64 = l0
	// binary: f64.copysign
	s0f64 = math.Copysign(s0f64, s1f64)
	// return
	return s0f64
	// end_block
lbl6:
	// block
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2146435072
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2146435072
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
	s0i32 = l7
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 23760
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// return
	return s0f64
	// end_block
lbl8:
	// block
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2146435072
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 67108864
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// end_block
lbl10:
	// const
	s0f64 = 1.5707963267948966
	// get_local
	s1f64 = l0
	// binary: f64.copysign
	s0f64 = math.Copysign(s0f64, s1f64)
	// return
	return s0f64
	// end_block
lbl9:
	// block
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
		goto lbl12
	}
	// const
	s0f64 = 0
	// set_local
	l8 = s0f64
	// get_local
	s0i32 = l6
	// const
	s1i32 = 67108864
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// end_block
lbl12:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l1
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// unary: f64.abs
	s0f64 = math.Abs(s0f64)
	// call
	s0f64 = f607(ctx, s0f64)
	// set_local
	l8 = s0f64
	// end_block
lbl11:
	// block
	// block
	// block
	// get_local
	s0i32 = l7
	// br_table
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl15
	case 2:
		goto lbl14
	default:
		goto lbl13
	}
	// end_block
lbl15:
	// get_local
	s0f64 = l8
	// unary: f64.neg
	s0f64 = -s0f64
	// return
	return s0f64
	// end_block
lbl14:
	// const
	s0f64 = 3.141592653589793
	// get_local
	s1f64 = l8
	// const
	s2f64 = -1.2246467991473532e-16
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// return
	return s0f64
	// end_block
lbl13:
	// get_local
	s0f64 = l8
	// const
	s1f64 = -1.2246467991473532e-16
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = -3.141592653589793
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl7:
	// get_local
	s0i32 = l7
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 23792
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l8 = s0f64
	// end_block
lbl2:
	// get_local
	s0f64 = l8
	// return
	return s0f64
}
