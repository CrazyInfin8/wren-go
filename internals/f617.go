package internals

import (
	"encoding/binary"
	"math"
)

func f617(ctx *Context, l0 float64) float64 {
	var l1 int64
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var l6 int64
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	var s3f64 float64
	_ = s3f64
	// block
	// block
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l1 = s0i64
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
	l2 = s0i32
	// const
	s1i32 = -969
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 63
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
	s0i32 = l2
	// set_local
	l3 = s0i32
	// br
	goto lbl1
	// end_block
lbl2:
	// const
	s0f64 = 1
	// set_local
	l4 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = 969
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
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1033
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
	// const
	s0f64 = 0
	// set_local
	l4 = s0f64
	// get_local
	s0i64 = l1
	// const
	s1i64 = -4503599627370496
	// binary: i64.eq
	if s0i64 == s1i64 {
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
	s0i32 = l2
	// const
	s1i32 = 2047
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl3:
	// block
	// get_local
	s0i64 = l1
	// const
	s1i64 = -1
	// binary: i64.gt_s
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// const
	s0i32 = 0
	// call
	s0f64 = f615(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl4:
	// const
	s0i32 = 0
	// call
	s0f64 = f616(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl1:
	// const
	s0i32 = 0
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+26608):]))
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// const
	s1i32 = 0
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+26616):]))
	// tee_local
	l4 = s1f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l4
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l4 = s0f64
	// const
	s1i32 = 0
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+26632):]))
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l4
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26624):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l0
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l0 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l4 = s0f64
	// get_local
	s1f64 = l4
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26664):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26656):]))
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l4
	// get_local
	s2f64 = l0
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+26648):]))
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+26640):]))
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l5
	// unary: i64.reinterpret/f64
	s2i64 = int64(math.Float64bits(s2f64))
	// tee_local
	l1 = s2i64
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// const
	s3i32 = 4
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// const
	s3i32 = 2032
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l2 = s2i32
	// const
	s3i32 = 26720
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// get_local
	s3f64 = l0
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// get_local
	s0i32 = l2
	// const
	s1i32 = 26728
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i64 = l1
	// const
	s2i64 = 45
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l6 = s0i64
	// block
	// get_local
	s0i32 = l3
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0f64 = l0
	// get_local
	s1i64 = l6
	// get_local
	s2i64 = l1
	// call
	s0f64 = f618(ctx, s0f64, s1i64, s2i64)
	// return
	return s0f64
	// end_block
lbl5:
	// get_local
	s0i64 = l6
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l4 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l4
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l4 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l4
	// return
	return s0f64
}
