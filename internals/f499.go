package internals

import (
	"encoding/binary"
)

func f499(ctx *Context, l0 int32, l1 int64, l2 int64, l3 int64) {
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
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int64
	_ = l20
	var l21 int64
	_ = l21
	var l22 int64
	_ = l22
	var l23 int32
	_ = l23
	var l24 int64
	_ = l24
	var l25 int64
	_ = l25
	var l26 int64
	_ = l26
	var l27 int64
	_ = l27
	var l28 int64
	_ = l28
	var l29 int64
	_ = l29
	var l30 int32
	_ = l30
	var l31 int32
	_ = l31
	var l32 int32
	_ = l32
	var l33 int32
	_ = l33
	var l34 int64
	_ = l34
	var l35 int64
	_ = l35
	var l36 int64
	_ = l36
	var l37 int32
	_ = l37
	var l38 int64
	_ = l38
	var l39 int64
	_ = l39
	var l40 int64
	_ = l40
	var l41 int64
	_ = l41
	var l42 int64
	_ = l42
	var l43 int64
	_ = l43
	var l44 int32
	_ = l44
	var l45 int32
	_ = l45
	var l46 int32
	_ = l46
	var l47 int32
	_ = l47
	var l48 int64
	_ = l48
	var l49 int64
	_ = l49
	var l50 int64
	_ = l50
	var l51 int32
	_ = l51
	var l52 int32
	_ = l52
	var l53 int32
	_ = l53
	var l54 int64
	_ = l54
	var l55 int64
	_ = l55
	var l56 int64
	_ = l56
	var l57 int64
	_ = l57
	var l58 int64
	_ = l58
	var l59 int64
	_ = l59
	var l60 int32
	_ = l60
	var l61 int32
	_ = l61
	var l62 int32
	_ = l62
	var l63 int32
	_ = l63
	var l64 int32
	_ = l64
	var l65 int64
	_ = l65
	var l66 int32
	_ = l66
	var l67 int32
	_ = l67
	var l68 int32
	_ = l68
	var l69 int64
	_ = l69
	var l70 int64
	_ = l70
	var l71 int64
	_ = l71
	var l72 int64
	_ = l72
	var l73 int64
	_ = l73
	var l74 int32
	_ = l74
	var l75 int32
	_ = l75
	var l76 int64
	_ = l76
	var l77 int64
	_ = l77
	var l78 int64
	_ = l78
	var l79 int64
	_ = l79
	var l80 int64
	_ = l80
	var l81 int64
	_ = l81
	var l82 int32
	_ = l82
	var l83 int32
	_ = l83
	var l84 int32
	_ = l84
	var l85 int32
	_ = l85
	var l86 int32
	_ = l86
	var l87 int32
	_ = l87
	var l88 int64
	_ = l88
	var l89 int32
	_ = l89
	var l90 int32
	_ = l90
	var l91 int64
	_ = l91
	var l92 int64
	_ = l92
	var l93 int64
	_ = l93
	var l94 int64
	_ = l94
	var l95 int64
	_ = l95
	var l96 int32
	_ = l96
	var l97 int32
	_ = l97
	var l98 int32
	_ = l98
	var l99 int32
	_ = l99
	var l100 int32
	_ = l100
	var l101 int64
	_ = l101
	var l102 int64
	_ = l102
	var l103 int64
	_ = l103
	var l104 int64
	_ = l104
	var l105 int64
	_ = l105
	var l106 int64
	_ = l106
	var l107 int64
	_ = l107
	var l108 int32
	_ = l108
	var l109 int32
	_ = l109
	var l110 int32
	_ = l110
	var l111 int32
	_ = l111
	var l112 int64
	_ = l112
	var l113 int64
	_ = l113
	var l114 int64
	_ = l114
	var l115 int64
	_ = l115
	var l116 int64
	_ = l116
	var l117 int64
	_ = l117
	var l118 int32
	_ = l118
	var l119 int32
	_ = l119
	var l120 int32
	_ = l120
	var l121 int32
	_ = l121
	var l122 int64
	_ = l122
	var l123 int64
	_ = l123
	var l124 int64
	_ = l124
	var l125 int64
	_ = l125
	var l126 int64
	_ = l126
	var l127 int64
	_ = l127
	var l128 int32
	_ = l128
	var l129 int32
	_ = l129
	var l130 int32
	_ = l130
	var l131 int32
	_ = l131
	var l132 int32
	_ = l132
	var l133 int32
	_ = l133
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l4 = s0i32
	// const
	s0i32 = 80
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+76):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+64):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l2
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l3
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+48):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+76):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+64):]))
	// set_local
	l10 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l11 = s0i64
	// get_local
	s0i64 = l10
	// get_local
	s1i64 = l11
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l12 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l13 = s0i64
	// get_local
	s0i64 = l12
	// set_local
	l14 = s0i64
	// get_local
	s0i64 = l13
	// set_local
	l15 = s0i64
	// get_local
	s0i64 = l14
	// get_local
	s1i64 = l15
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l16 = s0i32
	// const
	s0i32 = 1
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l16
	// get_local
	s1i32 = l17
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l18 = s0i32
	// block
	// get_local
	s0i32 = l18
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+64):]))
	// set_local
	l20 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l21 = s0i64
	// get_local
	s0i64 = l20
	// get_local
	s1i64 = l21
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l22 = s0i64
	// get_local
	s0i64 = l22
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l23
	// call
	f131(ctx, s0i32, s1i32)
	// end_block
lbl0:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l24 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l25 = s0i64
	// get_local
	s0i64 = l24
	// get_local
	s1i64 = l25
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l26 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l27 = s0i64
	// get_local
	s0i64 = l26
	// set_local
	l28 = s0i64
	// get_local
	s0i64 = l27
	// set_local
	l29 = s0i64
	// get_local
	s0i64 = l28
	// get_local
	s1i64 = l29
	// binary: i64.eq
	if s0i64 == s1i64 {
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
		goto lbl1
	}
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l34 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l35 = s0i64
	// get_local
	s0i64 = l34
	// get_local
	s1i64 = l35
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l36 = s0i64
	// get_local
	s0i64 = l36
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l33
	// get_local
	s1i32 = l37
	// call
	f131(ctx, s0i32, s1i32)
	// end_block
lbl1:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+48):]))
	// set_local
	l38 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l39 = s0i64
	// get_local
	s0i64 = l38
	// get_local
	s1i64 = l39
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l40 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l41 = s0i64
	// get_local
	s0i64 = l40
	// set_local
	l42 = s0i64
	// get_local
	s0i64 = l41
	// set_local
	l43 = s0i64
	// get_local
	s0i64 = l42
	// get_local
	s1i64 = l43
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l44 = s0i32
	// const
	s0i32 = 1
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l44
	// get_local
	s1i32 = l45
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l46 = s0i32
	// block
	// get_local
	s0i32 = l46
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
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+48):]))
	// set_local
	l48 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l49 = s0i64
	// get_local
	s0i64 = l48
	// get_local
	s1i64 = l49
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l50 = s0i64
	// get_local
	s0i64 = l50
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l51
	// call
	f131(ctx, s0i32, s1i32)
	// end_block
lbl2:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+76):]))
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l52
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+6188):]))
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+64):]))
	// set_local
	l54 = s0i64
	// get_local
	s0i32 = l53
	// get_local
	s1i64 = l54
	// call
	s0i64 = f344(ctx, s0i32, s1i64)
	// set_local
	l55 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l55
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+32):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+32):]))
	// set_local
	l56 = s0i64
	// const
	s0i64 = 9222246136947933188
	// set_local
	l57 = s0i64
	// get_local
	s0i64 = l56
	// set_local
	l58 = s0i64
	// get_local
	s0i64 = l57
	// set_local
	l59 = s0i64
	// get_local
	s0i64 = l58
	// get_local
	s1i64 = l59
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l60 = s0i32
	// const
	s0i32 = 1
	// set_local
	l61 = s0i32
	// get_local
	s0i32 = l60
	// get_local
	s1i32 = l61
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l62 = s0i32
	// block
	// get_local
	s0i32 = l62
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
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l63
	// call
	s0i32 = f111(ctx, s0i32)
	// set_local
	l64 = s0i32
	// get_local
	s0i32 = l64
	// call
	s0i64 = f122(ctx, s0i32)
	// set_local
	l65 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l65
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+32):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l66 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+76):]))
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l67
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+6188):]))
	// set_local
	l68 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+64):]))
	// set_local
	l69 = s0i64
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+32):]))
	// set_local
	l70 = s0i64
	// get_local
	s0i32 = l66
	// get_local
	s1i32 = l68
	// get_local
	s2i64 = l69
	// get_local
	s3i64 = l70
	// call
	f132(ctx, s0i32, s1i32, s2i64, s3i64)
	// end_block
lbl3:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+32):]))
	// set_local
	l71 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l72 = s0i64
	// get_local
	s0i64 = l71
	// get_local
	s1i64 = l72
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l73 = s0i64
	// get_local
	s0i64 = l73
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l74 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l74
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l75 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l76 = s0i64
	// get_local
	s0i32 = l75
	// get_local
	s1i64 = l76
	// call
	s0i64 = f344(ctx, s0i32, s1i64)
	// set_local
	l77 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l77
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l78 = s0i64
	// const
	s0i64 = 9222246136947933188
	// set_local
	l79 = s0i64
	// get_local
	s0i64 = l78
	// set_local
	l80 = s0i64
	// get_local
	s0i64 = l79
	// set_local
	l81 = s0i64
	// get_local
	s0i64 = l80
	// get_local
	s1i64 = l81
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l82 = s0i32
	// const
	s0i32 = 1
	// set_local
	l83 = s0i32
	// get_local
	s0i32 = l82
	// get_local
	s1i32 = l83
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l84 = s0i32
	// block
	// get_local
	s0i32 = l84
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
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l85 = s0i32
	// const
	s0i32 = 0
	// set_local
	l86 = s0i32
	// get_local
	s0i32 = l85
	// get_local
	s1i32 = l86
	// call
	s0i32 = f337(ctx, s0i32, s1i32)
	// set_local
	l87 = s0i32
	// get_local
	s0i32 = l87
	// call
	s0i64 = f122(ctx, s0i32)
	// set_local
	l88 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i64 = l88
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l89 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l90 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l91 = s0i64
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l92 = s0i64
	// get_local
	s0i32 = l89
	// get_local
	s1i32 = l90
	// get_local
	s2i64 = l91
	// get_local
	s3i64 = l92
	// call
	f132(ctx, s0i32, s1i32, s2i64, s3i64)
	// end_block
lbl4:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l93 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l94 = s0i64
	// get_local
	s0i64 = l93
	// get_local
	s1i64 = l94
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l95 = s0i64
	// get_local
	s0i64 = l95
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l96 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l96
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l97 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l98 = s0i32
	// const
	s0i32 = 16
	// set_local
	l99 = s0i32
	// get_local
	s0i32 = l98
	// get_local
	s1i32 = l99
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l100 = s0i32
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+48):]))
	// set_local
	l101 = s0i64
	// get_local
	s0i32 = l97
	// get_local
	s1i32 = l100
	// get_local
	s2i64 = l101
	// call
	f338(ctx, s0i32, s1i32, s2i64)
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+64):]))
	// set_local
	l102 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l103 = s0i64
	// get_local
	s0i64 = l102
	// get_local
	s1i64 = l103
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l104 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l105 = s0i64
	// get_local
	s0i64 = l104
	// set_local
	l106 = s0i64
	// get_local
	s0i64 = l105
	// set_local
	l107 = s0i64
	// get_local
	s0i64 = l106
	// get_local
	s1i64 = l107
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l108 = s0i32
	// const
	s0i32 = 1
	// set_local
	l109 = s0i32
	// get_local
	s0i32 = l108
	// get_local
	s1i32 = l109
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l110 = s0i32
	// block
	// get_local
	s0i32 = l110
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l111 = s0i32
	// get_local
	s0i32 = l111
	// call
	f133(ctx, s0i32)
	// end_block
lbl5:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l112 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l113 = s0i64
	// get_local
	s0i64 = l112
	// get_local
	s1i64 = l113
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l114 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l115 = s0i64
	// get_local
	s0i64 = l114
	// set_local
	l116 = s0i64
	// get_local
	s0i64 = l115
	// set_local
	l117 = s0i64
	// get_local
	s0i64 = l116
	// get_local
	s1i64 = l117
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l118 = s0i32
	// const
	s0i32 = 1
	// set_local
	l119 = s0i32
	// get_local
	s0i32 = l118
	// get_local
	s1i32 = l119
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l120 = s0i32
	// block
	// get_local
	s0i32 = l120
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
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l121 = s0i32
	// get_local
	s0i32 = l121
	// call
	f133(ctx, s0i32)
	// end_block
lbl6:
	// get_local
	s0i32 = l6
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+48):]))
	// set_local
	l122 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l123 = s0i64
	// get_local
	s0i64 = l122
	// get_local
	s1i64 = l123
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l124 = s0i64
	// const
	s0i64 = -1125899906842624
	// set_local
	l125 = s0i64
	// get_local
	s0i64 = l124
	// set_local
	l126 = s0i64
	// get_local
	s0i64 = l125
	// set_local
	l127 = s0i64
	// get_local
	s0i64 = l126
	// get_local
	s1i64 = l127
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l128 = s0i32
	// const
	s0i32 = 1
	// set_local
	l129 = s0i32
	// get_local
	s0i32 = l128
	// get_local
	s1i32 = l129
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l130 = s0i32
	// block
	// get_local
	s0i32 = l130
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
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l131 = s0i32
	// get_local
	s0i32 = l131
	// call
	f133(ctx, s0i32)
	// end_block
lbl7:
	// const
	s0i32 = 80
	// set_local
	l132 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l132
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l133 = s0i32
	// get_local
	s0i32 = l133
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
