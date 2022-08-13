package internals

import (
	"encoding/binary"
	"math"
)

func f609(ctx *Context, l0 float64) float64 {
	var l1 int64
	_ = l1
	var l2 float64
	_ = l2
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
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
	var s7i32 int32
	_ = s7i32
	var s9i32 int32
	_ = s9i32
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
	var s8f64 float64
	_ = s8f64
	var s9f64 float64
	_ = s9f64
	// block
	// get_local
	s0f64 = l0
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// tee_local
	l1 = s0i64
	// const
	s1i64 = -4606619468846596096
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// const
	s1i64 = 854320534781951
	// binary: i64.gt_u
	if uint64(s0i64) > uint64(s1i64) {
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
	s1f64 = -1
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l0 = s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l0
	// const
	s3f64 = 1.34217728e+08
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l2 = s2f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l2
	// binary: f64.sub
	s1f64 = s1f64 - s2f64
	// tee_local
	l2 = s1f64
	// get_local
	s2f64 = l2
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+25960):]))
	// tee_local
	l3 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l4 = s1f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l5 = s0f64
	// get_local
	s1f64 = l0
	// get_local
	s2f64 = l2
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l2
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// get_local
	s3f64 = l3
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// get_local
	s2f64 = l4
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l5
	// binary: f64.sub
	s3f64 = s3f64 - s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l0
	// binary: f64.mul
	s3f64 = s3f64 * s4f64
	// tee_local
	l2 = s3f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l3 = s2f64
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+25968):]))
	// get_local
	s4f64 = l0
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+25976):]))
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// get_local
	s4f64 = l2
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+25984):]))
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// get_local
	s4f64 = l3
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+25992):]))
	// get_local
	s6f64 = l0
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+26000):]))
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// get_local
	s6f64 = l2
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+26008):]))
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// get_local
	s6f64 = l3
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+26016):]))
	// get_local
	s8f64 = l0
	// const
	s9i32 = 0
	// load: f64.load
	s9f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s9i32+26024):]))
	// binary: f64.mul
	s8f64 = s8f64 * s9f64
	// binary: f64.add
	s7f64 = s7f64 + s8f64
	// get_local
	s8f64 = l2
	// const
	s9i32 = 0
	// load: f64.load
	s9f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s9i32+26032):]))
	// binary: f64.mul
	s8f64 = s8f64 * s9f64
	// binary: f64.add
	s7f64 = s7f64 + s8f64
	// get_local
	s8f64 = l3
	// const
	s9i32 = 0
	// load: f64.load
	s9f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s9i32+26040):]))
	// binary: f64.mul
	s8f64 = s8f64 * s9f64
	// binary: f64.add
	s7f64 = s7f64 + s8f64
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// return
	return s0f64
	// end_block
lbl0:
	// block
	// block
	// get_local
	s0i64 = l1
	// const
	s1i64 = 48
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// tee_local
	l6 = s0i32
	// const
	s1i32 = -16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 32736
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
	// block
	// get_local
	s0i64 = l1
	// const
	s1i64 = 9223372036854775807
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
		goto lbl3
	}
	// const
	s0i32 = 1
	// call
	s0f64 = f629(ctx, s0i32)
	// return
	return s0f64
	// end_block
lbl3:
	// get_local
	s0i64 = l1
	// const
	s1i64 = 9218868437227405312
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// block
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 32768
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 32752
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 32752
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	// call
	s0f64 = f627(ctx, s0f64)
	// return
	return s0f64
	// end_block
lbl4:
	// get_local
	s0f64 = l0
	// const
	s1f64 = 4.503599627370496e+15
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// unary: i64.reinterpret/f64
	s0i64 = int64(math.Float64bits(s0f64))
	// const
	s1i64 = -234187180623265792
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// set_local
	l1 = s0i64
	// end_block
lbl2:
	// get_local
	s0i64 = l1
	// const
	s1i64 = -4604367669032910848
	// binary: i64.add
	s0i64 = s0i64 + s1i64
	// tee_local
	l7 = s0i64
	// const
	s1i64 = 45
	// binary: i64.shr_u
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// const
	s1i32 = 127
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 4
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 26056
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// const
	s1i32 = 0
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+25904):]))
	// get_local
	s2i64 = l7
	// const
	s3i64 = 52
	// binary: i64.shr_s
	s2i64 = s2i64 >> (uint64(s3i64) & 63)
	// unary: i32.wrap/i64
	s2i32 = int32(s2i64)
	// unary: f64.convert_s/i32
	s2f64 = float64(s2i32)
	// tee_local
	l3 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l4 = s0f64
	// get_local
	s1i32 = l6
	// const
	s2i32 = 26048
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i64 = l1
	// get_local
	s3i64 = l7
	// const
	s4i64 = -4503599627370496
	// binary: i64.and
	s3i64 = s3i64 & s4i64
	// binary: i64.sub
	s2i64 = s2i64 - s3i64
	// unary: f64.reinterpret/i64
	s2f64 = math.Float64frombits(uint64(s2i64))
	// get_local
	s3i32 = l6
	// const
	s4i32 = 28096
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+0):]))
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// get_local
	s3i32 = l6
	// const
	s4i32 = 28104
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+0):]))
	// binary: f64.sub
	s2f64 = s2f64 - s3f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// tee_local
	l0 = s1f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// tee_local
	l5 = s0f64
	// const
	s1i32 = 0
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+25920):]))
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l0
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// tee_local
	l2 = s2f64
	// binary: f64.mul
	s1f64 = s1f64 * s2f64
	// const
	s2i32 = 0
	// load: f64.load
	s2f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s2i32+25912):]))
	// get_local
	s3f64 = l3
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// get_local
	s3f64 = l0
	// get_local
	s4f64 = l4
	// get_local
	s5f64 = l5
	// binary: f64.sub
	s4f64 = s4f64 - s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.add
	s2f64 = s2f64 + s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// get_local
	s2f64 = l0
	// get_local
	s3f64 = l2
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// const
	s3i32 = 0
	// load: f64.load
	s3f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s3i32+25928):]))
	// get_local
	s4f64 = l0
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+25936):]))
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// get_local
	s4f64 = l2
	// const
	s5i32 = 0
	// load: f64.load
	s5f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s5i32+25944):]))
	// get_local
	s6f64 = l0
	// const
	s7i32 = 0
	// load: f64.load
	s7f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s7i32+25952):]))
	// binary: f64.mul
	s6f64 = s6f64 * s7f64
	// binary: f64.add
	s5f64 = s5f64 + s6f64
	// binary: f64.mul
	s4f64 = s4f64 * s5f64
	// binary: f64.add
	s3f64 = s3f64 + s4f64
	// binary: f64.mul
	s2f64 = s2f64 * s3f64
	// binary: f64.add
	s1f64 = s1f64 + s2f64
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l0 = s0f64
	// end_block
lbl1:
	// get_local
	s0f64 = l0
	// return
	return s0f64
}
