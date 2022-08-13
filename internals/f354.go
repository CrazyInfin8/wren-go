package internals

import (
	"encoding/binary"
	"math"
)

func f354(ctx *Context, l0 int32, l1 int32, l2 float64, l3 int32) int32 {
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
	var l14 float64
	_ = l14
	var l15 int32
	_ = l15
	var l16 float64
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 float64
	_ = l23
	var l24 float64
	_ = l24
	var l25 int32
	_ = l25
	var l26 float64
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 float64
	_ = l30
	var l31 int32
	_ = l31
	var l32 float64
	_ = l32
	var l33 int32
	_ = l33
	var l34 int32
	_ = l34
	var l35 int32
	_ = l35
	var l36 float64
	_ = l36
	var l37 float64
	_ = l37
	var l38 int32
	_ = l38
	var l39 float64
	_ = l39
	var l40 int32
	_ = l40
	var l41 int32
	_ = l41
	var l42 int32
	_ = l42
	var l43 int32
	_ = l43
	var l44 int32
	_ = l44
	var l45 int32
	_ = l45
	var l46 int32
	_ = l46
	var l47 int32
	_ = l47
	var l48 int32
	_ = l48
	var l49 int32
	_ = l49
	var l50 int64
	_ = l50
	var l51 int32
	_ = l51
	var l52 int32
	_ = l52
	var l53 int32
	_ = l53
	var l54 int32
	_ = l54
	var l55 int32
	_ = l55
	var l56 int32
	_ = l56
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
	l4 = s0i32
	// const
	s0i32 = 32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l6
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1f64 = l2
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l8 = s0f64
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1f64 = l8
	// get_local
	s2i32 = l9
	// call
	s0i32 = f351(ctx, s0i32, s1f64, s2i32)
	// set_local
	l10 = s0i32
	// const
	s0i32 = 1
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l11
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l12 = s0i32
	// block
	// block
	// get_local
	s0i32 = l12
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = -1
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l13
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l14 = s0f64
	// const
	s0i32 = 0
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l15
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l16 = s0f64
	// get_local
	s0f64 = l14
	// get_local
	s1f64 = l16
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l17 = s0i32
	// const
	s0i32 = 1
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l18
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l19 = s0i32
	// block
	// get_local
	s0i32 = l19
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
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l20
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l21 = s0f64
	// get_local
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l22 = s0f64
	// get_local
	s0f64 = l21
	// get_local
	s1f64 = l22
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l23 = s0f64
	// get_local
	s0i32 = l6
	// get_local
	s1f64 = l23
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// end_block
lbl2:
	// get_local
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l24 = s0f64
	// const
	s0i32 = 0
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l25
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l26 = s0f64
	// get_local
	s0f64 = l24
	// get_local
	s1f64 = l26
	// binary: f64.ge
	if s0f64 >= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l27 = s0i32
	// const
	s0i32 = 1
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l28
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l29 = s0i32
	// block
	// get_local
	s0i32 = l29
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
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l30 = s0f64
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l31
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l32 = s0f64
	// get_local
	s0f64 = l30
	// get_local
	s1f64 = l32
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l33 = s0i32
	// const
	s0i32 = 1
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l33
	// get_local
	s1i32 = l34
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l35
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
	s0i32 = l6
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l36 = s0f64
	// const
	s0f64 = 4.294967296e+09
	// set_local
	l37 = s0f64
	// get_local
	s0f64 = l36
	// get_local
	s1f64 = l37
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l38 = s0i32
	// const
	s0f64 = 0
	// set_local
	l39 = s0f64
	// get_local
	s0f64 = l36
	// get_local
	s1f64 = l39
	// binary: f64.ge
	if s0f64 >= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l38
	// get_local
	s1i32 = l40
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l41
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l42 = s0i32
	// block
	// block
	// get_local
	s0i32 = l42
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0f64 = l36
	// unary: i32.trunc_u/f64
	s0i32 = int32(uint32(math.Trunc(s0f64)))
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l43
	// set_local
	l44 = s0i32
	// br
	goto lbl4
	// end_block
lbl5:
	// const
	s0i32 = 0
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l45
	// set_local
	l44 = s0i32
	// end_block
lbl4:
	// get_local
	s0i32 = l44
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l46
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 3105
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l49
	// get_local
	s2i32 = l6
	// call
	s0i64 = f314(ctx, s0i32, s1i32, s2i32)
	// set_local
	l50 = s0i64
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l51
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l52
	// get_local
	s1i64 = l50
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+48):], uint64(s1i64))
	// const
	s0i32 = -1
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l53
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l54 = s0i32
	// const
	s0i32 = 32
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l55
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l56
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l54
	// return
	return s0i32
}
