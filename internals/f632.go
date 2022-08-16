package internals

import (
	"encoding/binary"
	"math"
)

func f632(ctx *Context, l0 float64) float64 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
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
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l1 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// const
	s1i64 = 32
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 2147483647
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1072243195
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1045430272
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
	s0f64 = l0
	// const
	s1f64 = 0
	// const
	s2i32 = 0
	// call
	s0f64 = f613(ctx, s0f64, s1f64, s2i32)
	// set_local
	l0 = s0f64
	// br
	goto lbl0
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 2146435072
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// set_local
	l0 = s0f64
	// br
	goto lbl0
	// end_block
lbl2:
	// block
	// block
	// block
	// block
	// get_local
	s0f64 = l0
	// get_local
	s1i32 = l1
	// call
	s0i32 = f612(ctx, s0f64, s1i32)
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl6
	case 1:
		goto lbl5
	case 2:
		goto lbl4
	default:
		goto lbl3
	}
	// end_block
lbl6:
	// get_local
	s0i32 = l1
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l1
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// const
	s2i32 = 1
	// call
	s0f64 = f613(ctx, s0f64, s1f64, s2i32)
	// set_local
	l0 = s0f64
	// br
	goto lbl0
	// end_block
lbl5:
	// get_local
	s0i32 = l1
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l1
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// call
	s0f64 = f610(ctx, s0f64, s1f64)
	// set_local
	l0 = s0f64
	// br
	goto lbl0
	// end_block
lbl4:
	// get_local
	s0i32 = l1
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l1
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// const
	s2i32 = 1
	// call
	s0f64 = f613(ctx, s0f64, s1f64, s2i32)
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l0 = s0f64
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0i32 = l1
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l1
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+8):]))
	// call
	s0f64 = f610(ctx, s0f64, s1f64)
	// unary: f64.neg
	s0f64 = -s0f64
	// set_local
	l0 = s0f64
	// end_block
lbl0:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
