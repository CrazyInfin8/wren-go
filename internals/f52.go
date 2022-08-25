package internals

import (
	"encoding/binary"
)

func f52(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	// get_global
	s0i32 = ctx.G0
	// set_local
	l5 = s0i32
	// const
	s0i32 = 208
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+204):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+200):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+196):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+192):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+188):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
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
	// store: i32.store8
	ctx.Mem[int(s0i32+133)] = uint8(s1i32)
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l10
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+132)])
	// set_local
	l11 = s0i32
	// const
	s0i32 = 1
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l11
	// get_local
	s1i32 = l12
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l13 = s0i32
	// block
	// block
	// get_local
	s0i32 = l13
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// br
	goto lbl0
	// end_block
lbl1:
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l14
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l15
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+148):]))
	// set_local
	l16 = s0i32
	// const
	s0i32 = 0
	// set_local
	l17 = s0i32
	// get_local
	s0i32 = l16
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l17
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l18
	// get_local
	s1i32 = l19
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
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
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l22 = s0i32
	// block
	// get_local
	s0i32 = l22
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
	// br
	goto lbl0
	// end_block
lbl2:
	// const
	s0i32 = 16
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l23
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l24
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+196):]))
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l26
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 7434
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l25
	// get_local
	s1i32 = l27
	// get_local
	s2i32 = l7
	// call
	s0i32 = f552(ctx, s0i32, s1i32, s2i32)
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l28
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// const
	s0i32 = 16
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l29
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l30
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l31
	// get_local
	s1i32 = l32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+192):]))
	// set_local
	l34 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+188):]))
	// set_local
	l35 = s0i32
	// get_local
	s0i32 = l33
	// get_local
	s1i32 = l34
	// get_local
	s2i32 = l35
	// call
	s0i32 = f575(ctx, s0i32, s1i32, s2i32)
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l37
	// get_local
	s1i32 = l36
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l38
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l39
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l40
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l41 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l41
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l42 = s0i32
	// const
	s0i32 = 0
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l42
	// set_local
	l44 = s0i32
	// get_local
	s0i32 = l43
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l44
	// get_local
	s1i32 = l45
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l46 = s0i32
	// const
	s0i32 = 1
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l46
	// get_local
	s1i32 = l47
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l48 = s0i32
	// block
	// block
	// get_local
	s0i32 = l48
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
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l49 = s0i32
	// const
	s0i32 = 24
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l49
	// get_local
	s1i32 = l50
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l51 = s0i32
	// get_local
	s0i32 = l51
	// set_local
	l52 = s0i32
	// br
	goto lbl3
	// end_block
lbl4:
	// const
	s0i32 = 1968
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l53
	// set_local
	l52 = s0i32
	// end_block
lbl3:
	// get_local
	s0i32 = l52
	// set_local
	l54 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l54
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l55
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l56
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+148):]))
	// set_local
	l57 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+204):]))
	// set_local
	l58 = s0i32
	// get_local
	s0i32 = l58
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l59 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l60 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+200):]))
	// set_local
	l61 = s0i32
	// const
	s0i32 = 16
	// set_local
	l62 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l62
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l63
	// set_local
	l64 = s0i32
	// const
	s0i32 = 0
	// set_local
	l65 = s0i32
	// get_local
	s0i32 = l59
	// get_local
	s1i32 = l65
	// get_local
	s2i32 = l60
	// get_local
	s3i32 = l61
	// get_local
	s4i32 = l64
	// get_local
	s5i32 = l57
	// call_indirect
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].f == nil {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	table[s5i32].f.(func(*Context, int32, int32, int32, int32, int32))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// end_block
lbl0:
	// const
	s0i32 = 208
	// set_local
	l66 = s0i32
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l66
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l67
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
