package internals

import (
	"encoding/binary"
)

func f520(ctx *Context, l0 int32, l1 int32) {
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
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int64
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
	var l36 int64
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
	var l44 int64
	_ = l44
	var l45 int64
	_ = l45
	var l46 int64
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
	var l56 int32
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
	var l63 int32
	_ = l63
	var l64 int64
	_ = l64
	var l65 int64
	_ = l65
	var l66 int64
	_ = l66
	var l67 int64
	_ = l67
	var l68 int32
	_ = l68
	var l69 int32
	_ = l69
	var l70 int32
	_ = l70
	var l71 int32
	_ = l71
	var l72 int32
	_ = l72
	var l73 int64
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
	var l81 int64
	_ = l81
	var l82 int64
	_ = l82
	var l83 int64
	_ = l83
	var l84 int32
	_ = l84
	var l85 int32
	_ = l85
	var l86 int32
	_ = l86
	var l87 int32
	_ = l87
	var l88 int32
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
	var l109 int32
	_ = l109
	var l110 int32
	_ = l110
	var l111 int32
	_ = l111
	var l112 int32
	_ = l112
	var l113 int32
	_ = l113
	var l114 int32
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
	var l123 int32
	_ = l123
	var l124 int32
	_ = l124
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
	// get_global
	s0i32 = ctx.G0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 48
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l5 = s0i32
	// const
	s0i32 = 1437
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// call
	f56(ctx, s0i32, s1i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l7 = s0i32
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// const
	s0i32 = 7273
	// set_local
	l9 = s0i32
	// const
	s0i32 = 5
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l9
	// get_local
	s3i32 = l10
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// const
	s0i32 = 0
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// block
	// loop
lbl1:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l13
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l14
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l15
	// get_local
	s1i32 = l16
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
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
		goto lbl0
	}
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l20
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l22 = s0i32
	// const
	s0i32 = 4
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l22
	// get_local
	s1i32 = l23
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l25
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l26
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l27 = s0i64
	// const
	s0i64 = 9222246136947933188
	// set_local
	l28 = s0i64
	// get_local
	s0i64 = l27
	// set_local
	l29 = s0i64
	// get_local
	s0i64 = l28
	// set_local
	l30 = s0i64
	// get_local
	s0i64 = l29
	// get_local
	s1i64 = l30
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l31 = s0i32
	// const
	s0i32 = 1
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l31
	// get_local
	s1i32 = l32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l33 = s0i32
	// block
	// block
	// get_local
	s0i32 = l33
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
	// br
	goto lbl2
	// end_block
lbl3:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l35
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l36 = s0i64
	// get_local
	s0i32 = l34
	// get_local
	s1i64 = l36
	// call
	f93(ctx, s0i32, s1i64)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l37 = s0i32
	// const
	s0i32 = 1437
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l37
	// get_local
	s1i32 = l38
	// call
	f56(ctx, s0i32, s1i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l39 = s0i32
	// const
	s0i32 = 0
	// set_local
	l40 = s0i32
	// const
	s0i32 = 7273
	// set_local
	l41 = s0i32
	// const
	s0i32 = 5
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l39
	// get_local
	s1i32 = l40
	// get_local
	s2i32 = l41
	// get_local
	s3i32 = l42
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l43
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l44 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l45 = s0i64
	// get_local
	s0i64 = l44
	// get_local
	s1i64 = l45
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l46 = s0i64
	// get_local
	s0i64 = l46
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l47
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// block
	// loop
lbl5:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l50
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l49
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l51
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l52
	// get_local
	s1i32 = l53
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l54 = s0i32
	// const
	s0i32 = 1
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l54
	// get_local
	s1i32 = l55
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l56
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
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l57
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l58 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l59 = s0i32
	// const
	s0i32 = 4
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
	s0i32 = l4
	// get_local
	s1i32 = l62
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l63
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l64 = s0i64
	// const
	s0i64 = 9222246136947933188
	// set_local
	l65 = s0i64
	// get_local
	s0i64 = l64
	// set_local
	l66 = s0i64
	// get_local
	s0i64 = l65
	// set_local
	l67 = s0i64
	// get_local
	s0i64 = l66
	// get_local
	s1i64 = l67
	// binary: i64.eq
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l68 = s0i32
	// const
	s0i32 = 1
	// set_local
	l69 = s0i32
	// get_local
	s0i32 = l68
	// get_local
	s1i32 = l69
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l70 = s0i32
	// block
	// block
	// get_local
	s0i32 = l70
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
	// br
	goto lbl6
	// end_block
lbl7:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l71 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l72 = s0i32
	// get_local
	s0i32 = l72
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l73 = s0i64
	// get_local
	s0i32 = l71
	// get_local
	s1i64 = l73
	// call
	f93(ctx, s0i32, s1i64)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l74 = s0i32
	// const
	s0i32 = 1124
	// set_local
	l75 = s0i32
	// get_local
	s0i32 = l74
	// get_local
	s1i32 = l75
	// call
	f56(ctx, s0i32, s1i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l76 = s0i32
	// const
	s0i32 = 0
	// set_local
	l77 = s0i32
	// const
	s0i32 = 7273
	// set_local
	l78 = s0i32
	// const
	s0i32 = 5
	// set_local
	l79 = s0i32
	// get_local
	s0i32 = l76
	// get_local
	s1i32 = l77
	// get_local
	s2i32 = l78
	// get_local
	s3i32 = l79
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l80 = s0i32
	// get_local
	s0i32 = l80
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l81 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l82 = s0i64
	// get_local
	s0i64 = l81
	// get_local
	s1i64 = l82
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l83 = s0i64
	// get_local
	s0i64 = l83
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l84 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l84
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l85 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l85
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// block
	// loop
lbl9:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l86 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l87 = s0i32
	// get_local
	s0i32 = l87
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l88 = s0i32
	// get_local
	s0i32 = l86
	// set_local
	l89 = s0i32
	// get_local
	s0i32 = l88
	// set_local
	l90 = s0i32
	// get_local
	s0i32 = l89
	// get_local
	s1i32 = l90
	// binary: i32.lt_s
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l91 = s0i32
	// const
	s0i32 = 1
	// set_local
	l92 = s0i32
	// get_local
	s0i32 = l91
	// get_local
	s1i32 = l92
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l93 = s0i32
	// get_local
	s0i32 = l93
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
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l94 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l95 = s0i32
	// get_local
	s0i32 = l95
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l96 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l97 = s0i32
	// const
	s0i32 = 3
	// set_local
	l98 = s0i32
	// get_local
	s0i32 = l97
	// get_local
	s1i32 = l98
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l99 = s0i32
	// get_local
	s0i32 = l96
	// get_local
	s1i32 = l99
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l100 = s0i32
	// get_local
	s0i32 = l100
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l101 = s0i64
	// get_local
	s0i32 = l94
	// get_local
	s1i64 = l101
	// call
	f93(ctx, s0i32, s1i64)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l102 = s0i32
	// const
	s0i32 = 1
	// set_local
	l103 = s0i32
	// const
	s0i32 = 7160
	// set_local
	l104 = s0i32
	// const
	s0i32 = 11
	// set_local
	l105 = s0i32
	// get_local
	s0i32 = l102
	// get_local
	s1i32 = l103
	// get_local
	s2i32 = l104
	// get_local
	s3i32 = l105
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l106 = s0i32
	// const
	s0i32 = 1
	// set_local
	l107 = s0i32
	// get_local
	s0i32 = l106
	// get_local
	s1i32 = l107
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l108 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l108
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl9
	// end
	// end_block
lbl8:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l109 = s0i32
	// const
	s0i32 = 2
	// set_local
	l110 = s0i32
	// const
	s0i32 = 6747
	// set_local
	l111 = s0i32
	// const
	s0i32 = 13
	// set_local
	l112 = s0i32
	// get_local
	s0i32 = l109
	// get_local
	s1i32 = l110
	// get_local
	s2i32 = l111
	// get_local
	s3i32 = l112
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// end_block
lbl6:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l113 = s0i32
	// const
	s0i32 = 1
	// set_local
	l114 = s0i32
	// get_local
	s0i32 = l113
	// get_local
	s1i32 = l114
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l115 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l115
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// br
	goto lbl5
	// end
	// end_block
lbl4:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l116 = s0i32
	// const
	s0i32 = 2
	// set_local
	l117 = s0i32
	// const
	s0i32 = 6747
	// set_local
	l118 = s0i32
	// const
	s0i32 = 13
	// set_local
	l119 = s0i32
	// get_local
	s0i32 = l116
	// get_local
	s1i32 = l117
	// get_local
	s2i32 = l118
	// get_local
	s3i32 = l119
	// call
	f57(ctx, s0i32, s1i32, s2i32, s3i32)
	// end_block
lbl2:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l120 = s0i32
	// const
	s0i32 = 1
	// set_local
	l121 = s0i32
	// get_local
	s0i32 = l120
	// get_local
	s1i32 = l121
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l122 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l122
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// br
	goto lbl1
	// end
	// end_block
lbl0:
	// const
	s0i32 = 48
	// set_local
	l123 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l123
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l124 = s0i32
	// get_local
	s0i32 = l124
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
