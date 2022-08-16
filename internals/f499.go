package internals

import (
	"encoding/binary"
)

func f499(ctx *Context, l0 int32, l1 int32) {
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
	var l7 int32
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
	var l15 int64
	_ = l15
	var l16 float64
	_ = l16
	var l17 int64
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 float64
	_ = l22
	var l23 int64
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 int32
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
	var l37 float64
	_ = l37
	var l38 int64
	_ = l38
	var l39 int32
	_ = l39
	var l40 int32
	_ = l40
	var l41 int32
	_ = l41
	var l42 int32
	_ = l42
	var l43 int32
	_ = l43
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
	// get_global
	s0i32 = ctx.G0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 16
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store8
	ctx.Mem[int(s0i32+11)] = uint8(s1i32)
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42076):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+11)])
	// set_local
	l8 = s0i32
	// const
	s0i32 = 1
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l9
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l10 = s0i32
	// block
	// block
	// get_local
	s0i32 = l10
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l11
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l12 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// const
	s0i32 = 16
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l13
	// get_local
	s2i32 = l14
	// call
	s0i64 = f600(ctx, s0i32, s1i32, s2i32)
	// set_local
	l15 = s0i64
	// get_local
	s0i64 = l15
	// unary: f64.convert_s/i64
	s0f64 = float64(s0i64)
	// set_local
	l16 = s0f64
	// get_local
	s0f64 = l16
	// call
	s0i64 = f321(ctx, s0f64)
	// set_local
	l17 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l18
	// get_local
	s1i64 = l17
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+40):], uint64(s1i64))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l19
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l20 = s0i32
	// const
	s0i32 = 0
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// call
	s0f64 = f598(ctx, s0i32, s1i32)
	// set_local
	l22 = s0f64
	// get_local
	s0f64 = l22
	// call
	s0i64 = f321(ctx, s0f64)
	// set_local
	l23 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l24
	// get_local
	s1i64 = l23
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+40):], uint64(s1i64))
	// end_block
lbl0:
	// const
	s0i32 = 0
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l25
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42076):]))
	// set_local
	l26 = s0i32
	// const
	s0i32 = 68
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l26
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l27
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l28
	// get_local
	s1i32 = l29
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l30 = s0i32
	// const
	s0i32 = 1
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l30
	// get_local
	s1i32 = l31
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l32 = s0i32
	// block
	// get_local
	s0i32 = l32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l33 = s0i32
	// const
	s0i32 = 4
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l34
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 5233
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l33
	// get_local
	s1i32 = l35
	// get_local
	s2i32 = l4
	// call
	f110(ctx, s0i32, s1i32, s2i32)
	// const
	s0i32 = 0
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l36
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l37 = s0f64
	// get_local
	s0f64 = l37
	// call
	s0i64 = f321(ctx, s0f64)
	// set_local
	l38 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l39
	// get_local
	s1i64 = l38
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+40):], uint64(s1i64))
	// end_block
lbl2:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l40 = s0i32
	// const
	s0i32 = 58
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l40
	// get_local
	s1i32 = l41
	// call
	f97(ctx, s0i32, s1i32)
	// const
	s0i32 = 16
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l42
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l43
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
