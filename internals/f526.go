package internals

import (
	"encoding/binary"
	"math"
)

func f526(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
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
	var l13 float64
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 float64
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 float64
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 float64
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f64 float64
	_ = s0f64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l1 = s0i32
	// const
	s0i32 = 16
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l3
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l5 = s0i32
	// const
	s0i32 = 10
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
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
	s0i32 = l5
	// br_table
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl1
	case 2:
		goto lbl1
	case 3:
		goto lbl4
	case 4:
		goto lbl1
	case 5:
		goto lbl1
	case 6:
		goto lbl1
	case 7:
		goto lbl1
	case 8:
		goto lbl1
	case 9:
		goto lbl3
	case 10:
		goto lbl2
	default:
		goto lbl1
	}
	// end_block
lbl5:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+36):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l8
	// call
	s0i32 = f526(ctx, s0i32)
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl4:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l11
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+52):]))
	// set_local
	l12 = s0i32
	// get_local
	s0i32 = l12
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l13 = s0f64
	// get_local
	s0f64 = l13
	// call
	s0i32 = f528(ctx, s0f64)
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l15
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// unary: f64.convert_s/i32
	s0f64 = float64(s0i32)
	// set_local
	l17 = s0f64
	// get_local
	s0f64 = l17
	// call
	s0i32 = f528(ctx, s0f64)
	// set_local
	l18 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l18
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l19 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l19
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl3:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l20 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l20
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l21 = s0i32
	// get_local
	s0i32 = l21
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+16):]))
	// set_local
	l22 = s0f64
	// get_local
	s0f64 = l22
	// call
	s0i32 = f528(ctx, s0f64)
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l24 = s0i32
	// get_local
	s0i32 = l24
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+24):]))
	// set_local
	l25 = s0f64
	// get_local
	s0f64 = l25
	// call
	s0i32 = f528(ctx, s0f64)
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l26
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l27
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl2:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l28 = s0i32
	// get_local
	s0i32 = l28
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l29 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l29
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl0
	// end_block
lbl1:
	// const
	s0i32 = 0
	// set_local
	l30 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l30
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l31 = s0i32
	// const
	s0i32 = 16
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l33
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l31
	// return
	return s0i32
}
