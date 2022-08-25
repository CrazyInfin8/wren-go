package internals

import (
	"encoding/binary"
	"math"
)

func f311(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int32 {
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
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int64
	_ = l18
	var l19 int64
	_ = l19
	var l20 int64
	_ = l20
	var l21 int64
	_ = l21
	var l22 int64
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
	var l28 int64
	_ = l28
	var l29 int64
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
	var l50 int32
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
	var l56 int64
	_ = l56
	var l57 int32
	_ = l57
	var l58 int32
	_ = l58
	var l59 int32
	_ = l59
	var l60 int32
	_ = l60
	var l61 int32
	_ = l61
	var l62 int32
	_ = l62
	var l63 int64
	_ = l63
	var l64 int64
	_ = l64
	var l65 int64
	_ = l65
	var l66 int64
	_ = l66
	var l67 int64
	_ = l67
	var l68 int64
	_ = l68
	var l69 int32
	_ = l69
	var l70 int32
	_ = l70
	var l71 int32
	_ = l71
	var l72 int32
	_ = l72
	var l73 int32
	_ = l73
	var l74 int32
	_ = l74
	var l75 int32
	_ = l75
	var l76 int32
	_ = l76
	var l77 int32
	_ = l77
	var l78 int32
	_ = l78
	var l79 int32
	_ = l79
	var l80 int32
	_ = l80
	var l81 int32
	_ = l81
	var l82 int32
	_ = l82
	var l83 int32
	_ = l83
	var l84 int32
	_ = l84
	var l85 int64
	_ = l85
	var l86 float64
	_ = l86
	var l87 float64
	_ = l87
	var l88 float64
	_ = l88
	var l89 int32
	_ = l89
	var l90 int32
	_ = l90
	var l91 int32
	_ = l91
	var l92 int32
	_ = l92
	var l93 int32
	_ = l93
	var l94 int32
	_ = l94
	var l95 int32
	_ = l95
	var l96 int64
	_ = l96
	var l97 int32
	_ = l97
	var l98 int32
	_ = l98
	var l99 int32
	_ = l99
	var l100 int32
	_ = l100
	var l101 int32
	_ = l101
	var l102 int32
	_ = l102
	var l103 int32
	_ = l103
	var l104 int32
	_ = l104
	var l105 int32
	_ = l105
	var l106 int32
	_ = l106
	var l107 int32
	_ = l107
	var l108 int32
	_ = l108
	var l109 int64
	_ = l109
	var l110 int64
	_ = l110
	var l111 int64
	_ = l111
	var l112 int64
	_ = l112
	var l113 int64
	_ = l113
	var l114 int64
	_ = l114
	var l115 int32
	_ = l115
	var l116 int32
	_ = l116
	var l117 int32
	_ = l117
	var l118 int32
	_ = l118
	var l119 int32
	_ = l119
	var l120 int32
	_ = l120
	var l121 int32
	_ = l121
	var l122 int32
	_ = l122
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
	// get_global
	s0i32 = ctx.G0
	// set_local
	l6 = s0i32
	// const
	s0i32 = 48
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l8
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i64 = l4
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l9
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l10 = s0i32
	// const
	s0i32 = 65536
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l10
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l11
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l13
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l14 = s0i32
	// const
	s0i32 = 1
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l15
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l16 = s0i32
	// block
	// block
	// get_local
	s0i32 = l16
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
	// const
	s0i32 = -2
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l17
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l8
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l18 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l19 = s0i64
	// get_local
	s0i64 = l18
	// get_local
	s1i64 = l19
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l20 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l21 = s0i64
	// get_local
	s0i64 = l20
	// set_local
	l22 = s0i64
	// get_local
	s0i64 = l21
	// set_local
	l23 = s0i64
	// get_local
	s0i64 = l22
	// get_local
	s1i64 = l23
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l24 = s0i32
	// const
	s0i32 = 1
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l25
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l26 = s0i32
	// block
	// get_local
	s0i32 = l26
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
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l8
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l28 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l29 = s0i64
	// get_local
	s0i64 = l28
	// get_local
	s1i64 = l29
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l30 = s0i64
	// get_local
	s0i64 = l30
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l31
	// call
	f136(ctx, s0i32, s1i32)
	// end_block
lbl2:
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l32 = s0i32
	// const
	s0i32 = 28
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l33
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l34
	// get_local
	s1i32 = l35
	// get_local
	s2i32 = l36
	// call
	s0i32 = f90(ctx, s0i32, s1i32, s2i32)
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l37
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l38 = s0i32
	// const
	s0i32 = -1
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l38
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l39
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l40
	// get_local
	s1i32 = l41
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l42 = s0i32
	// const
	s0i32 = 1
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l42
	// get_local
	s1i32 = l43
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l44 = s0i32
	// block
	// block
	// get_local
	s0i32 = l44
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l46 = s0i32
	// const
	s0i32 = 28
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l46
	// get_local
	s1i32 = l47
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l45
	// get_local
	s1i32 = l48
	// get_local
	s2i32 = l49
	// get_local
	s3i32 = l50
	// call
	s0i32 = f312(ctx, s0i32, s1i32, s2i32, s3i32)
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l51
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l53 = s0i32
	// const
	s0i32 = 16
	// set_local
	l54 = s0i32
	// get_local
	s0i32 = l53
	// get_local
	s1i32 = l54
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l8
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l56 = s0i64
	// get_local
	s0i32 = l52
	// get_local
	s1i32 = l55
	// get_local
	s2i64 = l56
	// call
	f343(ctx, s0i32, s1i32, s2i64)
	// br
	goto lbl3
	// end_block
lbl4:
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l57
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l58 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l59 = s0i32
	// const
	s0i32 = 3
	// set_local
	l60 = s0i32
	// get_local
	s0i32 = l59
	// get_local
	s1i32 = l60
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l61 = s0i32
	// get_local
	s0i32 = l58
	// get_local
	s1i32 = l61
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l62 = s0i32
	// get_local
	s0i32 = l62
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l63 = s0i64
	// const
	s0i64 = 9222246136947933184
	// set_local
	l64 = s0i64
	// get_local
	s0i64 = l63
	// get_local
	s1i64 = l64
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l65 = s0i64
	// const
	s0i64 = 9222246136947933184
	// set_local
	l66 = s0i64
	// get_local
	s0i64 = l65
	// set_local
	l67 = s0i64
	// get_local
	s0i64 = l66
	// set_local
	l68 = s0i64
	// get_local
	s0i64 = l67
	// get_local
	s1i64 = l68
	// binary: i64.ne
	if s0i64 != s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l69 = s0i32
	// const
	s0i32 = 1
	// set_local
	l70 = s0i32
	// get_local
	s0i32 = l69
	// get_local
	s1i32 = l70
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l71 = s0i32
	// block
	// block
	// get_local
	s0i32 = l71
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l72 = s0i32
	// const
	s0i32 = 0
	// set_local
	l73 = s0i32
	// get_local
	s0i32 = l72
	// set_local
	l74 = s0i32
	// get_local
	s0i32 = l73
	// set_local
	l75 = s0i32
	// get_local
	s0i32 = l74
	// get_local
	s1i32 = l75
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l76 = s0i32
	// const
	s0i32 = 1
	// set_local
	l77 = s0i32
	// get_local
	s0i32 = l76
	// get_local
	s1i32 = l77
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l78 = s0i32
	// block
	// get_local
	s0i32 = l78
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l79 = s0i32
	// get_local
	s0i32 = l79
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l80 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l81 = s0i32
	// const
	s0i32 = 3
	// set_local
	l82 = s0i32
	// get_local
	s0i32 = l81
	// get_local
	s1i32 = l82
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l83 = s0i32
	// get_local
	s0i32 = l80
	// get_local
	s1i32 = l83
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l84 = s0i32
	// get_local
	s0i32 = l84
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l85 = s0i64
	// get_local
	s0i64 = l85
	// call
	s0f64 = f98(ctx, s0i64)
	// set_local
	l86 = s0f64
	// get_local
	s0f64 = l86
	// unary: f64.abs
	s0f64 = math.Abs(s0f64)
	// set_local
	l87 = s0f64
	// const
	s0f64 = 2.147483648e+09
	// set_local
	l88 = s0f64
	// get_local
	s0f64 = l87
	// get_local
	s1f64 = l88
	// binary: f64.lt
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l89 = s0i32
	// get_local
	s0i32 = l89
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l90 = s0i32
	// block
	// block
	// get_local
	s0i32 = l90
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0f64 = l86
	// unary: i32.trunc_s/f64
	s0i32 = int32(math.Trunc(s0f64))
	// set_local
	l91 = s0i32
	// get_local
	s0i32 = l91
	// set_local
	l92 = s0i32
	// br
	goto lbl8
	// end_block
lbl9:
	// const
	s0i32 = -2147483648
	// set_local
	l93 = s0i32
	// get_local
	s0i32 = l93
	// set_local
	l92 = s0i32
	// end_block
lbl8:
	// get_local
	s0i32 = l92
	// set_local
	l94 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l95 = s0i32
	// get_local
	s0i32 = l95
	// get_local
	s1i32 = l94
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl7:
	// get_local
	s0i32 = l8
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l96 = s0i64
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l97 = s0i32
	// get_local
	s0i32 = l97
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l98 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l99 = s0i32
	// const
	s0i32 = 3
	// set_local
	l100 = s0i32
	// get_local
	s0i32 = l99
	// get_local
	s1i32 = l100
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l101 = s0i32
	// get_local
	s0i32 = l98
	// get_local
	s1i32 = l101
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l102 = s0i32
	// get_local
	s0i32 = l102
	// get_local
	s1i64 = l96
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l103 = s0i32
	// get_local
	s0i32 = l103
	// call
	s0i32 = f89(ctx, s0i32)
	// set_local
	l104 = s0i32
	// const
	s0i32 = 1
	// set_local
	l105 = s0i32
	// get_local
	s0i32 = l104
	// get_local
	s1i32 = l105
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l106 = s0i32
	// block
	// get_local
	s0i32 = l106
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// const
	s0i32 = -3
	// set_local
	l107 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l107
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl10:
	// br
	goto lbl5
	// end_block
lbl6:
	// const
	s0i32 = -1
	// set_local
	l108 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l108
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl5:
	// end_block
lbl3:
	// get_local
	s0i32 = l8
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l109 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l110 = s0i64
	// get_local
	s0i64 = l109
	// get_local
	s1i64 = l110
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l111 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l112 = s0i64
	// get_local
	s0i64 = l111
	// set_local
	l113 = s0i64
	// get_local
	s0i64 = l112
	// set_local
	l114 = s0i64
	// get_local
	s0i64 = l113
	// get_local
	s1i64 = l114
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l115 = s0i32
	// const
	s0i32 = 1
	// set_local
	l116 = s0i32
	// get_local
	s0i32 = l115
	// get_local
	s1i32 = l116
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l117 = s0i32
	// block
	// get_local
	s0i32 = l117
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l118 = s0i32
	// get_local
	s0i32 = l118
	// call
	f138(ctx, s0i32)
	// end_block
lbl11:
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l119 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l119
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l120 = s0i32
	// const
	s0i32 = 48
	// set_local
	l121 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l121
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l122 = s0i32
	// get_local
	s0i32 = l122
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l120
	// return
	return s0i32
}
