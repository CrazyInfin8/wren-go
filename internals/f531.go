package internals

import (
	"encoding/binary"
)

func f531(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var l34 int64
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
	var l57 int64
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
	var l69 int64
	_ = l69
	var l70 int32
	_ = l70
	var l71 int32
	_ = l71
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
	var s6i32 int32
	_ = s6i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
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
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+44):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+40):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+36):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l9
	// store: i32.store8
	ctx.Mem[int(s0i32+35)] = uint8(s1i32)
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// const
	s0i64 = 0
	// set_local
	l10 = s0i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l10
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l11
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+136):]))
	// set_local
	l12 = s0i32
	// const
	s0i32 = 0
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l12
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l13
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l15
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l19
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+136):]))
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+35)])
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l25 = s0i32
	// const
	s0i32 = 16
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l26
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l27
	// set_local
	l28 = s0i32
	// const
	s0i32 = 1
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l24
	// get_local
	s1i32 = l29
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l28
	// get_local
	s1i32 = l21
	// get_local
	s2i32 = l22
	// get_local
	s3i32 = l23
	// get_local
	s4i32 = l30
	// get_local
	s5i32 = l25
	// get_local
	s6i32 = l20
	// call_indirect
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].f == nil {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	table[s6i32].f.(func(*Context, int32, int32, int32, int32, int32, int32))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	// const
	s0i32 = 16
	// set_local
	l31 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l31
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l32
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l33
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l34 = s0i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l34
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// end_block
lbl0:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l35 = s0i32
	// const
	s0i32 = 0
	// set_local
	l36 = s0i32
	// get_local
	s0i32 = l35
	// set_local
	l37 = s0i32
	// get_local
	s0i32 = l36
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l37
	// get_local
	s1i32 = l38
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l39 = s0i32
	// const
	s0i32 = 1
	// set_local
	l40 = s0i32
	// get_local
	s0i32 = l39
	// get_local
	s1i32 = l40
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l41 = s0i32
	// block
	// get_local
	s0i32 = l41
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
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l42 = s0i32
	// const
	s0i32 = 1902
	// set_local
	l43 = s0i32
	// get_local
	s0i32 = l42
	// get_local
	s1i32 = l43
	// call
	s0i32 = f579(ctx, s0i32, s1i32)
	// set_local
	l44 = s0i32
	// block
	// get_local
	s0i32 = l44
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l45 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+35)])
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l48 = s0i32
	// const
	s0i32 = 8
	// set_local
	l49 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l49
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l50 = s0i32
	// get_local
	s0i32 = l50
	// set_local
	l51 = s0i32
	// const
	s0i32 = 1
	// set_local
	l52 = s0i32
	// get_local
	s0i32 = l47
	// get_local
	s1i32 = l52
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l53 = s0i32
	// get_local
	s0i32 = l51
	// get_local
	s1i32 = l45
	// get_local
	s2i32 = l46
	// get_local
	s3i32 = l53
	// get_local
	s4i32 = l48
	// call
	f478(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// const
	s0i32 = 8
	// set_local
	l54 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l54
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l55 = s0i32
	// get_local
	s0i32 = l55
	// set_local
	l56 = s0i32
	// get_local
	s0i32 = l56
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l57 = s0i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l57
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// end_block
lbl2:
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+40):]))
	// set_local
	l58 = s0i32
	// const
	s0i32 = 1518
	// set_local
	l59 = s0i32
	// get_local
	s0i32 = l58
	// get_local
	s1i32 = l59
	// call
	s0i32 = f579(ctx, s0i32, s1i32)
	// set_local
	l60 = s0i32
	// block
	// get_local
	s0i32 = l60
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+44):]))
	// set_local
	l61 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l62 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+35)])
	// set_local
	l63 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l64 = s0i32
	// get_local
	s0i32 = l8
	// set_local
	l65 = s0i32
	// const
	s0i32 = 1
	// set_local
	l66 = s0i32
	// get_local
	s0i32 = l63
	// get_local
	s1i32 = l66
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l67 = s0i32
	// get_local
	s0i32 = l65
	// get_local
	s1i32 = l61
	// get_local
	s2i32 = l62
	// get_local
	s3i32 = l67
	// get_local
	s4i32 = l64
	// call
	f482(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	// get_local
	s0i32 = l8
	// set_local
	l68 = s0i32
	// get_local
	s0i32 = l68
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l69 = s0i64
	// get_local
	s0i32 = l0
	// get_local
	s1i64 = l69
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// end_block
lbl3:
	// end_block
lbl1:
	// const
	s0i32 = 48
	// set_local
	l70 = s0i32
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l70
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l71 = s0i32
	// get_local
	s0i32 = l71
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
