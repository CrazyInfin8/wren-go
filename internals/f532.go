package internals

import (
	"encoding/binary"
)

func f532(ctx *Context, l0 int32, l1 int64) int64 {
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
	var l13 int64
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
	var l36 int64
	_ = l36
	var l37 int64
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
	var l53 int64
	_ = l53
	var l54 int32
	_ = l54
	var l55 int64
	_ = l55
	var l56 int32
	_ = l56
	var l57 int32
	_ = l57
	var l58 int64
	_ = l58
	var l59 int32
	_ = l59
	var l60 int64
	_ = l60
	var l61 int64
	_ = l61
	var l62 int64
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
	var l71 int64
	_ = l71
	var l72 int32
	_ = l72
	var l73 int32
	_ = l73
	var l74 int64
	_ = l74
	var l75 int32
	_ = l75
	var l76 int32
	_ = l76
	var l77 int32
	_ = l77
	var l78 int64
	_ = l78
	var l79 int64
	_ = l79
	var l80 int32
	_ = l80
	var l81 int32
	_ = l81
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
	s0i32 = 64
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+52):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+40):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+128):]))
	// set_local
	l6 = s0i32
	// const
	s0i32 = 0
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l7
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l9
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
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
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l13 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l13
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l14
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l15
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l18
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l19 = s0i32
	// const
	s0i32 = 1
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l19
	// get_local
	s1i32 = l20
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l21 = s0i32
	// const
	s0i32 = 12
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l21
	// get_local
	s1i32 = l22
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l17
	// get_local
	s1i32 = l23
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l24
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l25
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l26
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+32):]))
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l27
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l28
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l29
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l30
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+128):]))
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l33 = s0i32
	// const
	s0i32 = 24
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l33
	// get_local
	s1i32 = l34
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l36 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l37 = s0i64
	// get_local
	s0i64 = l36
	// get_local
	s1i64 = l37
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l38 = s0i64
	// get_local
	s0i64 = l38
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l39 = s0i32
	// const
	s0i32 = 24
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l39
	// get_local
	s1i32 = l40
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l32
	// get_local
	s1i32 = l35
	// get_local
	s2i32 = l41
	// get_local
	s3i32 = l31
	// call_indirect
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].f == nil {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = table[s3i32].f.(func(*Context, int32, int32, int32) int32)(ctx, s0i32, s1i32, s2i32)
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l42
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l43 = s0i32
	// const
	s0i32 = 0
	// set_local
	l44 = s0i32
	// get_local
	s0i32 = l43
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l44
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l45
	// get_local
	s1i32 = l46
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l47 = s0i32
	// const
	s0i32 = 1
	// set_local
	l48 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l48
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l49 = s0i32
	// block
	// get_local
	s0i32 = l49
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
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l51 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l52
	// call
	s0i64 = f127(ctx, s0i32)
	// set_local
	l53 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l53
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l51
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// const
	s0i32 = 5736
	// set_local
	l54 = s0i32
	// get_local
	s0i32 = l50
	// get_local
	s1i32 = l54
	// get_local
	s2i32 = l4
	// call
	s0i64 = f319(ctx, s0i32, s1i32, s2i32)
	// set_local
	l55 = s0i64
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l56
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l57
	// get_local
	s1i64 = l55
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+48):], uint64(s1i64))
	// const
	s0i64 = 9222246136947933185
	// set_local
	l58 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l58
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// br
	goto lbl0
	// end_block
lbl2:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l59 = s0i32
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l60 = s0i64
	// const
	s0i64 = 1125899906842623
	// set_local
	l61 = s0i64
	// get_local
	s0i64 = l60
	// get_local
	s1i64 = l61
	// binary: i64.and
	s0i64 = s0i64 & s1i64
	// set_local
	l62 = s0i64
	// get_local
	s0i64 = l62
	// unary: i32.wrap/i64
	s0i32 = int32(s0i64)
	// set_local
	l63 = s0i32
	// const
	s0i32 = 24
	// set_local
	l64 = s0i32
	// get_local
	s0i32 = l63
	// get_local
	s1i32 = l64
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l65 = s0i32
	// get_local
	s0i32 = l59
	// set_local
	l66 = s0i32
	// get_local
	s0i32 = l65
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l66
	// get_local
	s1i32 = l67
	// binary: i32.eq
	if s0i32 == s1i32 {
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
		goto lbl3
	}
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l71 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l71
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l72 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l73 = s0i32
	// get_local
	s0i32 = l72
	// get_local
	s1i32 = l73
	// call
	s0i64 = f309(ctx, s0i32, s1i32)
	// set_local
	l74 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l74
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+40):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l75 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l76 = s0i32
	// const
	s0i32 = 0
	// set_local
	l77 = s0i32
	// get_local
	s0i32 = l75
	// get_local
	s1i32 = l76
	// get_local
	s2i32 = l77
	// get_local
	s3i32 = l77
	// call
	s0i32 = f303(ctx, s0i32, s1i32, s2i32, s3i32)
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+40):]))
	// set_local
	l78 = s0i64
	// get_local
	s0i32 = l4
	// get_local
	s1i64 = l78
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+56):], uint64(s1i64))
	// end_block
lbl0:
	// get_local
	s0i32 = l4
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+56):]))
	// set_local
	l79 = s0i64
	// const
	s0i32 = 64
	// set_local
	l80 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l80
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l81 = s0i32
	// get_local
	s0i32 = l81
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i64 = l79
	// return
	return s0i64
}
