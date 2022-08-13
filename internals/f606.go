package internals

import (
	"encoding/binary"
	"math"
)

func f606(ctx *Context, l0 float64) float64 {
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
	var l6 float64
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
	var s3i64 int64
	_ = s3i64
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
	// block
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l1 = s0i64
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
	s1i32 = 1141899264
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
	s0f64 = l0
	// get_local
	s1f64 = l0
	// binary: f64.ne
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
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
lbl1:
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1071382527
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
	s0i32 = l2
	// const
	s1i32 = 1044381696
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
	s0i32 = -1
	// set_local
	l3 = s0i32
	// const
	s0i32 = 1
	// set_local
	l2 = s0i32
	// br
	goto lbl2
	// end_block
lbl3:
	// get_local
	s0f64 = l0
	// unary: f64.abs
	s0f64 = math.Abs(s0f64)
	// set_local
	l0 = s0f64
	// block
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1072889855
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1072037887
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
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
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// const
	s1f64 = -1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l0
	// const
	s2f64 = 2
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l0 = s0f64
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// br
	goto lbl2
	// end_block
lbl6:
	// get_local
	s0f64 = l0
	// const
	s1f64 = -1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l0
	// const
	s2f64 = 1
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l0 = s0f64
	// const
	s0i32 = 1
	// set_local
	l3 = s0i32
	// br
	goto lbl4
	// end_block
lbl5:
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1073971199
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = -1.5
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// get_local
	s1f64 = l0
	// const
	s2f64 = 1.5
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 1
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l0 = s0f64
	// const
	s0i32 = 2
	// set_local
	l3 = s0i32
	// br
	goto lbl4
	// end_block
lbl7:
	// const
	s0f64 = -1
	// get_local
	s1f64 = l0
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l0 = s0f64
	// const
	s0i32 = 3
	// set_local
	l3 = s0i32
	// end_block
lbl4:
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// end_block
lbl2:
	// get_local
	s0f64 = l0
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
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l5
	// get_local
	s2f64 = l5
	// get_local
	s3f64 = l5
	// get_local
	s4f64 = l5
	// const
	s5f64 = -0.036531572744216916
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = -0.058335701337905735
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = -0.0769187620504483
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = -0.11111110405462356
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = -0.19999999999876483
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l6 = s0f64
	// get_local
	s0f64 = l4
	// get_local
	s1f64 = l5
	// get_local
	s2f64 = l5
	// get_local
	s3f64 = l5
	// get_local
	s4f64 = l5
	// get_local
	s5f64 = l5
	// const
	s6f64 = 0.016285820115365782
	// binary: f64.mul
	s5f64 = s5f64 * s6f64
	// const
	s6f64 = 0.049768779946159324
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// const
	s5f64 = 0.06661073137387531
	// binary: f64.add
	s4f64 = s4f64 + s5f64
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// const
	s4f64 = 0.09090887133436507
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3f64 = 0.14285714272503466
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2f64 = 0.3333333333333293
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l5 = s0f64
	// block
	// get_local
	s0i32 = l2
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// get_local
	s3f64 = l5
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// return
	return s0f64
	// end_block
lbl8:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 23632
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l6
	// get_local
	s3f64 = l5
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2i32 = l2
	// const
	s3i32 = 23664
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+0):]))
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// get_local
	s2f64 = l0
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// binary: f64.sub
	s0f64 = s0f64 - s1f64
	// tee_local
	l0 = s0f64
	// unary: f64.neg
	s0f64 = -s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2i64 = l1
	// const
	s3i64 = 0
	// binary: i64.lt_s
	if s2i64 < s3i64 {
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
	l0 = s0f64
	// end_block
lbl0:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
