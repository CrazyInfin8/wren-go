package internals

import (
	"math"
)

func f616(ctx *Context, l0 float64, l1 float64) float64 {
	var l2 int64
	_ = l2
	var l3 int64
	_ = l3
	var l4 int64
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 int32
	_ = l7
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
	var s3i64 int64
	_ = s3i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// block
	// block
	// get_local
	s0f64 = l1
	// get_local
	s1f64 = l1
	// binary: f64.ne
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0f64 = l1
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l2 = s0i64
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l3 = s0i64
	// unary: i64.eqz
	if s0i64 == 0 {
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
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l4 = s0i64
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
	s1i32 = 2047
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// end_block
lbl1:
	// get_local
	s0f64 = l0
	// get_local
	s1f64 = l1
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// tee_local
	l1 = s0f64
	// get_local
	s1f64 = l1
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// block
	// get_local
	s0i64 = l4
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l6 = s0i64
	// get_local
	s1i64 = l3
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
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
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// get_local
	s1f64 = l0
	// get_local
	s2i64 = l6
	// get_local
	s3i64 = l3
	// binary: i64.eq
	if s2i64 == s3i64 {
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
lbl2:
	// get_local
	s0i64 = l2
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
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// block
	// get_local
	s0i64 = l4
	// const
	s1i64 = 12
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l3 = s0i64
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// loop
lbl6:
	// get_local
	s0i32 = l5
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i64 = l3
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l3 = s0i64
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
		goto lbl6
	}
	// end
	// end_block
lbl5:
	// get_local
	s0i64 = l4
	// const
	s1i32 = 1
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l3 = s0i64
	// br
	goto lbl3
	// end_block
lbl4:
	// get_local
	s0i64 = l4
	// const
	s1i64 = 4503599627370495
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 4503599627370496
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// set_local
	l3 = s0i64
	// end_block
lbl3:
	// block
	// block
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// block
	// get_local
	s0i64 = l2
	// const
	s1i64 = 12
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l6 = s0i64
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// loop
lbl10:
	// get_local
	s0i32 = l7
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i64 = l6
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l6 = s0i64
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
		goto lbl10
	}
	// end
	// end_block
lbl9:
	// get_local
	s0i64 = l2
	// const
	s1i32 = 1
	// get_local
	s2i32 = l7
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l2 = s0i64
	// br
	goto lbl7
	// end_block
lbl8:
	// get_local
	s0i64 = l2
	// const
	s1i64 = 4503599627370495
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// const
	s1i64 = 4503599627370496
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// set_local
	l2 = s0i64
	// end_block
lbl7:
	// block
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l7
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// loop
lbl12:
	// block
	// get_local
	s0i64 = l3
	// get_local
	s1i64 = l2
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// tee_local
	l6 = s0i64
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i64 = l6
	// set_local
	l3 = s0i64
	// get_local
	s0i64 = l6
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
		goto lbl13
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl13:
	// get_local
	s0i64 = l3
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// set_local
	l3 = s0i64
	// get_local
	s0i32 = l5
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l7
	// binary: i32.gt_s
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// end
	// get_local
	s0i32 = l7
	// set_local
	l5 = s0i32
	// end_block
lbl11:
	// block
	// get_local
	s0i64 = l3
	// get_local
	s1i64 = l2
	// binary: i64.sub
	s0i64 = s0i64 - s1i64
	// tee_local
	l6 = s0i64
	// const
	s1i64 = 0
	// binary: i64.lt_s
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// get_local
	s0i64 = l6
	// set_local
	l3 = s0i64
	// get_local
	s0i64 = l6
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
		goto lbl14
	}
	// get_local
	s0f64 = l0
	// const
	s1f64 = 0
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// return
	return s0f64
	// end_block
lbl14:
	// block
	// block
	// get_local
	s0i64 = l3
	// const
	s1i64 = 4503599627370495
	// binary: i64.le_u
	if uint64(s0i64) <= uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl16
	}
	// get_local
	s0i64 = l3
	// set_local
	l6 = s0i64
	// br
	goto lbl15
	// end_block
lbl16:
	// loop
lbl17:
	// get_local
	s0i32 = l5
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i64 = l3
	// const
	s1i64 = 2251799813685248
	// binary: i64.lt_u
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l7 = s0i32
	// get_local
	s0i64 = l3
	// const
	s1i64 = 1
	// binary: i64.shl
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	// tee_local
	l6 = s0i64
	// set_local
	l3 = s0i64
	// get_local
	s0i32 = l7
	// br_if
	if s0i32 != 0 {
		goto lbl17
	}
	// end
	// end_block
lbl15:
	// get_local
	s0i64 = l4
	// const
	s1i64 = -9223372036854775808
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l3 = s0i64
	// block
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 1
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl19
	}
	// get_local
	s0i64 = l6
	// const
	s1i64 = -4503599627370496
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// get_local
	s1i32 = l5
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// const
	s2i64 = 52
	// binary: i64.shl
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// set_local
	l6 = s0i64
	// br
	goto lbl18
	// end_block
lbl19:
	// get_local
	s0i64 = l6
	// const
	s1i32 = 1
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// unary: i64.extend_u/i32
	s1i64 = int64(uint32(s1i32))
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// set_local
	l6 = s0i64
	// end_block
lbl18:
	// get_local
	s0i64 = l6
	// get_local
	s1i64 = l3
	// binary: i64.or
	s0i64 = s0i64 | s1i64
	// unary: f64.reinterpret/i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	// return
	return s0f64
}
