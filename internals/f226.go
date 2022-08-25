package internals

import (
	"encoding/binary"
	"math"
)

func f226(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 float64
	_ = l17
	var l18 int32
	_ = l18
	var l19 int64
	_ = l19
	var l20 float64
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 float64
	_ = l26
	var l27 float64
	_ = l27
	var l28 float64
	_ = l28
	var l29 float64
	_ = l29
	var l30 int64
	_ = l30
	var l31 int32
	_ = l31
	var l32 int32
	_ = l32
	var l33 int32
	_ = l33
	var l34 int32
	_ = l34
	var l35 int32
	_ = l35
	var l36 int32
	_ = l36
	var l37 int32
	_ = l37
	var l38 int32
	_ = l38
	var l39 int32
	_ = l39
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
	// set_local
	l2 = s0i32
	// const
	s0i32 = 32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l7 = s0i64
	// const
	s0i32 = 1651
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i64 = l7
	// get_local
	s2i32 = l8
	// call
	s0i32 = f330(ctx, s0i32, s1i64, s2i32)
	// set_local
	l9 = s0i32
	// const
	s0i32 = 1
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l9
	// get_local
	s1i32 = l10
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l11 = s0i32
	// block
	// block
	// get_local
	s0i32 = l11
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// set_local
	l12 = s0i32
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l13
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l14
	// store: i32.store8
	ctx.Mem[int(s0i32+31)] = uint8(s1i32)
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l15
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l16 = s0i64
	// get_local
	s0i64 = l16
	// call
	s0f64 = f98(ctx, s0i64)
	// set_local
	l17 = s0f64
	// get_local
	s0i32 = l4
	// get_local
	s1f64 = l17
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l18
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l19 = s0i64
	// get_local
	s0i64 = l19
	// call
	s0f64 = f98(ctx, s0i64)
	// set_local
	l20 = s0f64
	// get_local
	s0i32 = l4
	// get_local
	s1f64 = l20
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l21 = s0f64
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l22 = s0f64
	// get_local
	s0f64 = l21
	// get_local
	s1f64 = l22
	// binary: f64.le
	if s0f64 <= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l23 = s0i32
	// const
	s0i32 = 1
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l24
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l25 = s0i32
	// block
	// block
	// get_local
	s0i32 = l25
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l26 = s0f64
	// get_local
	s0f64 = l26
	// set_local
	l27 = s0f64
	// br
	goto lbl2
	// end_block
lbl3:
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l28 = s0f64
	// get_local
	s0f64 = l28
	// set_local
	l27 = s0f64
	// end_block
lbl2:
	// get_local
	s0f64 = l27
	// set_local
	l29 = s0f64
	// get_local
	s0f64 = l29
	// call
	s0i64 = f326(ctx, s0f64)
	// set_local
	l30 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l31
	// get_local
	s1i64 = l30
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// const
	s0i32 = 1
	// set_local
	l32 = s0i32
	// const
	s0i32 = 1
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l33
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l34
	// store: i32.store8
	ctx.Mem[int(s0i32+31)] = uint8(s1i32)
	// end_block
lbl0:
	// get_local
	s0i32 = l4
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+31)])
	// set_local
	l35 = s0i32
	// const
	s0i32 = 1
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l35
	// get_local
	s1i32 = l36
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l37 = s0i32
	// const
	s0i32 = 32
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l38
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l39
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l37
	// return
	return s0i32
}
