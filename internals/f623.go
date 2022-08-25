package internals

import (
	"encoding/binary"
	"math"
)

func f623(ctx *Context, l0 float64, l1 float64) float64 {
	var l2 int64
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
	_ = l9
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 float64
	_ = l12
	var l13 float64
	_ = l13
	var l14 float64
	_ = l14
	var l15 float64
	_ = l15
	var l16 float64
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	var s5f64 float64
	_ = s5f64
	var s6f64 float64
	_ = s6f64
	var s7f64 float64
	_ = s7f64
	// get_local
	s0f64 = l1
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
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 2047
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = -1086
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// block
	// block
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l6 = s0i64
	// const
	s1i64 = 52
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l7 = s0i32
	// const
	s1i32 = -2047
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = -2046
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
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = -129
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
	// end_block
lbl2:
	// block
	// get_local
	s0i64 = l2
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l9 = s0i64
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = -9007199254740993
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// const
	s0f64 = 1
	// set_local
	l10 = s0f64
	// get_local
	s0i64 = l6
	// const
	s1i64 = 4607182418800017408
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
	// get_local
	s0i64 = l9
	// unary: i64.eqz
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// block
	// block
	// get_local
	s0i64 = l6
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l6 = s0i64
	// const
	s1i64 = -9007199254740992
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i64 = l9
	// const
	s1i64 = -9007199254740991
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// end_block
lbl5:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl4:
	// get_local
	s0i64 = l6
	// const
	s1i64 = 9214364837600034816
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
	// const
	s0f64 = 0
	// get_local
	s1f64 = l1
	// get_local
	s2f64 = l1
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2i64 = l2
	// const
	s3i64 = 63
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// const
	s3i32 = 1
	// binary: i32.xor
	s2i32 = s2i32 ^ s3i32
	// get_local
	s3i64 = l6
	// const
	s4i64 = 9214364837600034816
	// binary: i64.lt_u
	if uint64(s3i64) < uint64(s4i64) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// return
	return s0f64
	// end_block
lbl3:
	// block
	// get_local
	s0i64 = l6
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// const
	s1i64 = -1
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = -9007199254740993
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l10 = s0f64
	// block
	// get_local
	s0i64 = l6
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
		goto lbl7
	}
	// get_local
	s0f64 = l10
	// unary: f64.neg
	s0f64 = -s0f64
	// get_local
	s1f64 = l10
	// get_local
	s2i64 = l2
	// call
	s2i32 = f624(ctx, s2i64)
	// const
	s3i32 = 1
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// set_local
	l10 = s0f64
	// end_block
lbl7:
	// get_local
	s0i64 = l2
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
		goto lbl0
	}
	// const
	s0f64 = 1
	// get_local
	s1f64 = l10
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// return
	return s0f64
	// end_block
lbl6:
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// block
	// get_local
	s0i64 = l6
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
		goto lbl8
	}
	// block
	// get_local
	s0i64 = l2
	// call
	s0i32 = f624(ctx, s0i64)
	// tee_local
	l8 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0f64 = l0
	// call
	s0f64 = f619(ctx, s0f64)
	// return
	return s0f64
	// end_block
lbl9:
	// get_local
	s0i32 = l7
	// const
	s1i32 = 2047
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i64 = l6
	// const
	s1i64 = 9223372036854775807
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l6 = s0i64
	// get_local
	s0i32 = l8
	// const
	s1i32 = 1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// const
	s1i32 = 18
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l8 = s0i32
	// end_block
lbl8:
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = -129
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// const
	s0f64 = 1
	// set_local
	l10 = s0f64
	// get_local
	s0i32 = l4
	// const
	s1i32 = 958
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
	s0i64 = l6
	// const
	s1i64 = 4607182418800017408
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
	s0i32 = l3
	// const
	s1i32 = 2048
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s1i64 = l6
	// const
	s2i64 = 4607182418800017409
	// binary: i64.lt_u
	if uint64(s1i64) < uint64(s2i64) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// const
	s0i32 = 0
	// call
	s0f64 = f613(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl11:
	// const
	s0i32 = 0
	// call
	s0f64 = f612(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl10:
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 4.503599627370496e+15
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// const
	s1i64 = 9223372036854775807
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = -234187180623265792
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l6 = s0i64
	// end_block
lbl1:
	// block
	// get_local
	s0i64 = l2
	// const
	s1i64 = -134217728
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l11 = s0f64
	// get_local
	s1i64 = l6
	// get_local
	s2i64 = l6
	// const
	s3i64 = -4604531861337669632
	// binary: i64.add
	s2i64 = s2i64 + s3i64
	// tee_local
	l2 = s2i64
	// const
	s3i64 = -4503599627370496
	// binary: i64.and
	s2i64 = s2i64 & s3i64
	// binary: i64.sub
	s1i64 = s1i64 - s2i64
	// tee_local
	l6 = s1i64
	// const
	s2i64 = 2147483648
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// const
	s2i64 = -4294967296
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// tee_local
	l10 = s1f64
	// get_local
	s2i64 = l2
	// const
	s3i64 = 45
	// binary: i64.shr_u
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// const
	s3i32 = 127
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 5
	// binary: i32.shl
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	// tee_local
	l5 = s2i32
	// const
	s3i32 = 35336
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// tee_local
	l12 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = -1
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// get_local
	s2f64 = l0
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+35280):]))
	// tee_local
	l13 = s3f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l14 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l15 = s1f64
	// get_local
	s2i64 = l2
	// const
	s3i64 = 52
	// binary: i64.shr_s
	s2i64 = s2i64 >> (uint64(s3i64) & 63)
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// unary: f64.convert_s/i32
	s2f64 = float64(s2i32)
	// tee_local
	l16 = s2f64
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+35264):]))
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3i32 = l5
	// const
	s4i32 = 35352
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+0):]))
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// tee_local
	l17 = s2f64
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l12
	// get_local
	s5i64 = l6
	// unary: f64.reinterpret/i64
	s5f64 = math.Float64frombits(uint64(s5i64))
	// get_local
	s6f64 = l10
	// binary: f64.sub
	s5f64 = s5f64 - s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// tee_local
	l18 = s4f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// tee_local
	l0 = s3f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// tee_local
	l10 = s2f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l12 = s1f64
	// get_local
	s2f64 = l15
	// get_local
	s3f64 = l10
	// get_local
	s4f64 = l12
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// get_local
	s3f64 = l18
	// get_local
	s4f64 = l14
	// get_local
	s5f64 = l13
	// get_local
	s6f64 = l0
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// tee_local
	l13 = s5f64
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// get_local
	s4f64 = l16
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+35272):]))
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// get_local
	s5i32 = l5
	// const
	s6i32 = 35360
	// binary: i32.add
	s5i32 = s5i32 + s6i32
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+0):]))
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// get_local
	s5f64 = l0
	// get_local
	s6f64 = l17
	// get_local
	s7f64 = l10
	// binary: f64.sub
	s6f64 = s6f64 - s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l0
	// get_local
	s5f64 = l13
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// tee_local
	l10 = s4f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// get_local
	s4f64 = l10
	// get_local
	s5f64 = l10
	// get_local
	s6f64 = l0
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+35328):]))
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+35320):]))
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// get_local
	s6f64 = l0
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+35312):]))
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+35304):]))
	// binary: f64.add
	s6f64 = s6f64 + s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// get_local
	s5f64 = l0
	// const
	s6i32 = 0
	// load: f64.load
	s6f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s6i32+35296):]))
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6i32 = 0
	// load: f64.load
	s6f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s6i32+35288):]))
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// tee_local
	l14 = s2f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l0 = s1f64
	// unary: i64.reinterpret/f64
	s1i64 = int64(math.Float64bits(s1f64))
	// const
	s2i64 = -134217728
	// binary: i64.and
	s1i64 = s1i64 & s2i64
	// unary: f64.reinterpret/i64
	s1f64 = math.Float64frombits(uint64(s1i64))
	// tee_local
	l10 = s1f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l13 = s0f64
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l6 = s0i64
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
	l5 = s0i32
	// const
	s1i32 = -969
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 63
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 968
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// const
	s0f64 = -1
	// const
	s1f64 = 1
	// get_local
	s2i32 = l8
	// select
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	// return
	return s0f64
	// end_block
lbl13:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1033
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// block
	// get_local
	s0i64 = l6
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
		goto lbl14
	}
	// get_local
	s0i32 = l8
	// call
	s0f64 = f612(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl14:
	// get_local
	s0i32 = l8
	// call
	s0f64 = f613(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl12:
	// get_local
	s0f64 = l1
	// get_local
	s1f64 = l11
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// get_local
	s1f64 = l10
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l14
	// get_local
	s2f64 = l12
	// get_local
	s3f64 = l0
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l10
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l1
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l13
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26672):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26680):]))
	// tee_local
	l0 = s2f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// tee_local
	l10 = s1f64
	// get_local
	s2f64 = l0
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l0 = s1f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26696):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l0
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+26688):]))
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l13
	// binary: f64.add
	s2f64 = s2f64 + s3f64
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
	l1 = s0f64
	// get_local
	s1f64 = l1
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26728):]))
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+26720):]))
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l1
	// get_local
	s2f64 = l0
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+26712):]))
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+26704):]))
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l10
	// unary: i64.reinterpret/f64
	s2i64 = int64(math.Float64bits(s2f64))
	// tee_local
	l6 = s2i64
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
	l7 = s2i32
	// const
	s3i32 = 26784
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
	s0i32 = l7
	// const
	s1i32 = 26792
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i64 = l6
	// get_local
	s2i32 = l8
	// unary: i64.extend_u/i32
	s2i64 = int64(uint32(s2i32))
	// binary: i64.add
	s1i64 = s1i64 + s2i64
	// const
	s2i64 = 45
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l2 = s0i64
	// block
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl15
	}
	// get_local
	s0f64 = l0
	// get_local
	s1i64 = l2
	// get_local
	s2i64 = l6
	// call
	s0f64 = f625(ctx, s0f64, s1i64, s2i64)
	// return
	return s0f64
	// end_block
lbl15:
	// get_local
	s0i64 = l2
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// tee_local
	l1 = s0f64
	// get_local
	s1f64 = l0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l10 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l10
	// return
	return s0f64
}
