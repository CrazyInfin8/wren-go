package internals

import (
	"encoding/binary"
)

func f129(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var l40 int64
	_ = l40
	var l41 int64
	_ = l41
	var l42 int64
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l10 = s0i32
	// const
	s0i32 = 76
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l10
	// get_local
	s1i32 = l11
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
	// block
	// get_local
	s0i32 = l10
	// br_table
	switch s0i32 {
	case 0:
		goto lbl4
	case 1:
		goto lbl6
	case 2:
		goto lbl6
	case 3:
		goto lbl6
	case 4:
		goto lbl6
	case 5:
		goto lbl6
	case 6:
		goto lbl6
	case 7:
		goto lbl6
	case 8:
		goto lbl6
	case 9:
		goto lbl6
	case 10:
		goto lbl6
	case 11:
		goto lbl6
	case 12:
		goto lbl6
	case 13:
		goto lbl5
	case 14:
		goto lbl5
	case 15:
		goto lbl5
	case 16:
		goto lbl5
	case 17:
		goto lbl4
	case 18:
		goto lbl4
	case 19:
		goto lbl5
	case 20:
		goto lbl5
	case 21:
		goto lbl5
	case 22:
		goto lbl5
	case 23:
		goto lbl6
	case 24:
		goto lbl4
	case 25:
		goto lbl4
	case 26:
		goto lbl4
	case 27:
		goto lbl4
	case 28:
		goto lbl4
	case 29:
		goto lbl4
	case 30:
		goto lbl4
	case 31:
		goto lbl4
	case 32:
		goto lbl4
	case 33:
		goto lbl4
	case 34:
		goto lbl4
	case 35:
		goto lbl4
	case 36:
		goto lbl4
	case 37:
		goto lbl4
	case 38:
		goto lbl4
	case 39:
		goto lbl4
	case 40:
		goto lbl4
	case 41:
		goto lbl3
	case 42:
		goto lbl3
	case 43:
		goto lbl3
	case 44:
		goto lbl3
	case 45:
		goto lbl3
	case 46:
		goto lbl3
	case 47:
		goto lbl3
	case 48:
		goto lbl3
	case 49:
		goto lbl3
	case 50:
		goto lbl3
	case 51:
		goto lbl3
	case 52:
		goto lbl3
	case 53:
		goto lbl3
	case 54:
		goto lbl3
	case 55:
		goto lbl3
	case 56:
		goto lbl3
	case 57:
		goto lbl3
	case 58:
		goto lbl4
	case 59:
		goto lbl4
	case 60:
		goto lbl4
	case 61:
		goto lbl4
	case 62:
		goto lbl4
	case 63:
		goto lbl6
	case 64:
		goto lbl6
	case 65:
		goto lbl2
	case 66:
		goto lbl6
	case 67:
		goto lbl6
	case 68:
		goto lbl5
	case 69:
		goto lbl6
	case 70:
		goto lbl6
	case 71:
		goto lbl4
	case 72:
		goto lbl4
	case 73:
		goto lbl6
	case 74:
		goto lbl4
	case 75:
		goto lbl4
	case 76:
		goto lbl6
	default:
		goto lbl1
	}
	// end_block
lbl6:
	// const
	s0i32 = 0
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l12
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl5:
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l13
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl4:
	// const
	s0i32 = 2
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l14
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl3:
	// const
	s0i32 = 4
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l15
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl2:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
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
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l16
	// get_local
	s1i32 = l19
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l20
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l21 = s0i32
	// const
	s0i32 = 255
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l22
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l23 = s0i32
	// const
	s0i32 = 8
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l24
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l27 = s0i32
	// const
	s0i32 = 2
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l27
	// get_local
	s1i32 = l28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l26
	// get_local
	s1i32 = l29
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l30
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// set_local
	l31 = s0i32
	// const
	s0i32 = 255
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
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l33
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l34
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l36 = s0i32
	// const
	s0i32 = 3
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
	s0i32 = l35
	// get_local
	s1i32 = l38
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l39
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l40 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l41 = s0i64
	// get_local
	s0i64 = l40
	// get_local
	s1i64 = l41
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l42 = s0i64
	// get_local
	s0i64 = l42
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l43
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l44 = s0i32
	// get_local
	s0i32 = l44
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+48):]))
	// set_local
	l45 = s0i32
	// const
	s0i32 = 1
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l45
	// get_local
	s1i32 = l46
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l47 = s0i32
	// const
	s0i32 = 2
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l49
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// const
	s0i32 = 0
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l50
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l51
	// return
	return s0i32
}
