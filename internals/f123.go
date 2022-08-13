package internals

import (
	"encoding/binary"
)

func f123(ctx *Context, l0 int32, l1 int32) {
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
	var l51 int64
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
	var l64 int32
	_ = l64
	var l65 int32
	_ = l65
	var l66 int32
	_ = l66
	var l67 int32
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
	var l85 int32
	_ = l85
	var l86 int64
	_ = l86
	var l87 int64
	_ = l87
	var l88 int64
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
	var l101 int32
	_ = l101
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// loop
lbl0:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l9
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l11 = s0i32
	// const
	s0i32 = -19
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l12
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l13 = s0i32
	// const
	s0i32 = 57
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l13
	// get_local
	s1i32 = l14
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l13
	// br_table
	switch s0i32 {
	case 0:
		goto lbl6
	case 1:
		goto lbl6
	case 2:
		goto lbl6
	case 3:
		goto lbl6
	case 4:
		goto lbl2
	case 5:
		goto lbl2
	case 6:
		goto lbl2
	case 7:
		goto lbl2
	case 8:
		goto lbl2
	case 9:
		goto lbl2
	case 10:
		goto lbl2
	case 11:
		goto lbl2
	case 12:
		goto lbl2
	case 13:
		goto lbl2
	case 14:
		goto lbl2
	case 15:
		goto lbl2
	case 16:
		goto lbl2
	case 17:
		goto lbl2
	case 18:
		goto lbl2
	case 19:
		goto lbl2
	case 20:
		goto lbl2
	case 21:
		goto lbl2
	case 22:
		goto lbl5
	case 23:
		goto lbl5
	case 24:
		goto lbl5
	case 25:
		goto lbl5
	case 26:
		goto lbl5
	case 27:
		goto lbl5
	case 28:
		goto lbl5
	case 29:
		goto lbl5
	case 30:
		goto lbl5
	case 31:
		goto lbl5
	case 32:
		goto lbl5
	case 33:
		goto lbl5
	case 34:
		goto lbl5
	case 35:
		goto lbl5
	case 36:
		goto lbl5
	case 37:
		goto lbl5
	case 38:
		goto lbl5
	case 39:
		goto lbl2
	case 40:
		goto lbl2
	case 41:
		goto lbl2
	case 42:
		goto lbl2
	case 43:
		goto lbl2
	case 44:
		goto lbl2
	case 45:
		goto lbl2
	case 46:
		goto lbl4
	case 47:
		goto lbl2
	case 48:
		goto lbl2
	case 49:
		goto lbl2
	case 50:
		goto lbl2
	case 51:
		goto lbl2
	case 52:
		goto lbl2
	case 53:
		goto lbl2
	case 54:
		goto lbl2
	case 55:
		goto lbl2
	case 56:
		goto lbl2
	case 57:
		goto lbl3
	default:
		goto lbl2
	}
	// end_block
lbl6:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l15
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l18
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l20 = s0i32
	// const
	s0i32 = 1
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l21
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l22
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l23
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l24 = s0i32
	// const
	s0i32 = 255
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
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l17
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l27
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// br
	goto lbl1
	// end_block
lbl5:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l28
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l30 = s0i32
	// const
	s0i32 = 3
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l30
	// get_local
	s1i32 = l31
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l29
	// get_local
	s1i32 = l32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l33
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l34 = s0i32
	// const
	s0i32 = 255
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l34
	// get_local
	s1i32 = l35
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l36 = s0i32
	// const
	s0i32 = 8
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l36
	// get_local
	s1i32 = l37
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l39
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l41 = s0i32
	// const
	s0i32 = 4
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l41
	// get_local
	s1i32 = l42
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l40
	// get_local
	s1i32 = l43
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l44 = s0i32
	// get_local
	s0i32 = l44
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l45 = s0i32
	// const
	s0i32 = 255
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l45
	// get_local
	s1i32 = l46
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l38
	// get_local
	s1i32 = l47
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l49
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l50
	// call
	s0i64 = f122(ctx, s0i32)
	// set_local
	l51 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l52
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l54 = s0i32
	// const
	s0i32 = 3
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l54
	// get_local
	s1i32 = l55
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l53
	// get_local
	s1i32 = l56
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l57
	// get_local
	s1i64 = l51
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// br
	goto lbl1
	// end_block
lbl4:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l58 = s0i32
	// get_local
	s0i32 = l58
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l59 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
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
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l62 = s0i32
	// get_local
	s0i32 = l59
	// get_local
	s1i32 = l62
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l63
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l64 = s0i32
	// const
	s0i32 = 255
	// set_local
	l65 = s0i32
	// get_local
	s0i32 = l64
	// get_local
	s1i32 = l65
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l66 = s0i32
	// const
	s0i32 = 8
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l66
	// get_local
	s1i32 = l67
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l68 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l69 = s0i32
	// get_local
	s0i32 = l69
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l70 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l71 = s0i32
	// const
	s0i32 = 2
	// set_local
	l72 = s0i32
	// get_local
	s0i32 = l71
	// get_local
	s1i32 = l72
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l73 = s0i32
	// get_local
	s0i32 = l70
	// get_local
	s1i32 = l73
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l74 = s0i32
	// get_local
	s0i32 = l74
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l75 = s0i32
	// const
	s0i32 = 255
	// set_local
	l76 = s0i32
	// get_local
	s0i32 = l75
	// get_local
	s1i32 = l76
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l77 = s0i32
	// get_local
	s0i32 = l68
	// get_local
	s1i32 = l77
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l78 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l78
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l79 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l80 = s0i32
	// get_local
	s0i32 = l80
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l81 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l82 = s0i32
	// const
	s0i32 = 3
	// set_local
	l83 = s0i32
	// get_local
	s0i32 = l82
	// get_local
	s1i32 = l83
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l84 = s0i32
	// get_local
	s0i32 = l81
	// get_local
	s1i32 = l84
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l85 = s0i32
	// get_local
	s0i32 = l85
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l86 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l87 = s0i64
	// get_local
	s0i64 = l86
	// get_local
	s1i64 = l87
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l88 = s0i64
	// get_local
	s0i64 = l88
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l89 = s0i32
	// get_local
	s0i32 = l79
	// get_local
	s1i32 = l89
	// call
	f123(ctx, s0i32, s1i32)
	// br
	goto lbl1
	// end_block
lbl3:
	// const
	s0i32 = 32
	// set_local
	l90 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l90
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l91 = s0i32
	// get_local
	s0i32 = l91
	// set_global
	ctx.G0 = s0i32
	// return
	return
	// end_block
lbl2:
	// end_block
lbl1:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l92 = s0i32
	// get_local
	s0i32 = l92
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l93 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l94 = s0i32
	// get_local
	s0i32 = l94
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l95 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l96 = s0i32
	// get_local
	s0i32 = l93
	// get_local
	s1i32 = l95
	// get_local
	s2i32 = l96
	// call
	s0i32 = f124(ctx, s0i32, s1i32, s2i32)
	// set_local
	l97 = s0i32
	// const
	s0i32 = 1
	// set_local
	l98 = s0i32
	// get_local
	s0i32 = l97
	// get_local
	s1i32 = l98
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l99 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l100 = s0i32
	// get_local
	s0i32 = l100
	// get_local
	s1i32 = l99
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l101 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l101
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// br
	goto lbl0
	// end
}
