package internals

import (
	"encoding/binary"
)

func f510(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_global
	s0i32 = ctx.G0
	// set_local
	l4 = s0i32
	// const
	s0i32 = 48
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
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+32):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l8
	// call
	s0i32 = f27(ctx, s0i32, s1i32)
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l10
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+6168):]))
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l12
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+49)])
	// set_local
	l13 = s0i32
	// const
	s0i32 = 1
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l13
	// get_local
	s1i32 = l14
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l15 = s0i32
	// block
	// block
	// get_local
	s0i32 = l15
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
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l16 = s0i32
	// const
	s0i32 = 36
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l16
	// get_local
	s1i32 = l17
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l18
	// set_local
	l19 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l20 = s0i32
	// const
	s0i32 = 24
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
	s0i32 = l22
	// set_local
	l19 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l19
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l23
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l24
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// block
	// loop
lbl3:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l26
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l25
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
	// binary: i32.lt_s
	if s0i32 < s1i32 {
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
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l33
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l35 = s0i32
	// const
	s0i32 = 2
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l35
	// get_local
	s1i32 = l36
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l34
	// get_local
	s1i32 = l37
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l38
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l39
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l40
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l41
	// get_local
	s1i32 = l42
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l43 = s0i32
	// const
	s0i32 = 1
	// set_local
	l44 = s0i32
	// get_local
	s0i32 = l43
	// get_local
	s1i32 = l44
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l45 = s0i32
	// block
	// get_local
	s0i32 = l45
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
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l46
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+49)])
	// set_local
	l47 = s0i32
	// const
	s0i32 = 7408
	// set_local
	l48 = s0i32
	// const
	s0i32 = 20774
	// set_local
	l49 = s0i32
	// const
	s0i32 = 1
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l50
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l48
	// get_local
	s1i32 = l49
	// get_local
	s2i32 = l51
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l52
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l54 = s0i32
	// get_local
	s0i32 = l54
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+6168):]))
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l55
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l56 = s0i32
	// const
	s0i32 = 24
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l56
	// get_local
	s1i32 = l57
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l58 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l59 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l60 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l60
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l59
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l58
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 5417
	// set_local
	l61 = s0i32
	// get_local
	s0i32 = l53
	// get_local
	s1i32 = l61
	// get_local
	s2i32 = l6
	// call
	f26(ctx, s0i32, s1i32, s2i32)
	// br
	goto lbl2
	// end_block
lbl4:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// set_local
	l62 = s0i32
	// const
	s0i32 = 1
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l62
	// get_local
	s1i32 = l63
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l64 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l64
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// br
	goto lbl3
	// end
	// end_block
lbl2:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l65 = s0i32
	// get_local
	s0i32 = l65
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l66 = s0i32
	// get_local
	s0i32 = l66
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l68 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l69 = s0i32
	// get_local
	s0i32 = l67
	// get_local
	s1i32 = l68
	// get_local
	s2i32 = l69
	// call
	f363(ctx, s0i32, s1i32, s2i32)
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l70 = s0i32
	// const
	s0i32 = 48
	// set_local
	l71 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l71
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l72 = s0i32
	// get_local
	s0i32 = l72
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l70
	// return
	return s0i32
}
