package internals

import (
	"encoding/binary"
)

func f352(ctx *Context, l0 int64) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int64
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
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 int64
	_ = l18
	var l19 int64
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
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
	s1i64 = l0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l4 = s0i64
	// get_local
	s0i64 = l4
	// call
	s0i32 = f353(ctx, s0i64)
	// set_local
	l5 = s0i32
	// const
	s0i32 = 1
	// set_local
	l6 = s0i32
	// const
	s0i32 = 1
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l6
	// set_local
	l9 = s0i32
	// block
	// get_local
	s0i32 = l8
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l10 = s0i64
	// const
	s0i32 = 0
	// set_local
	l11 = s0i32
	// get_local
	s0i64 = l10
	// get_local
	s1i32 = l11
	// call
	s0i32 = f311(ctx, s0i64, s1i32)
	// set_local
	l12 = s0i32
	// const
	s0i32 = 1
	// set_local
	l13 = s0i32
	// const
	s0i32 = 1
	// set_local
	l14 = s0i32
	// get_local
	s0i32 = l12
	// get_local
	s1i32 = l14
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l13
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l15
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l16 = s0i64
	// const
	s0i64 = 9222246136947933185
	// set_local
	l17 = s0i64
	// get_local
	s0i64 = l16
	// set_local
	l18 = s0i64
	// get_local
	s0i64 = l17
	// set_local
	l19 = s0i64
	// get_local
	s0i64 = l18
	// get_local
	s1i64 = l19
	// binary: i64.eq
	if s0i64 == s1i64 {
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
	// const
	s0i32 = 1
	// set_local
	l22 = s0i32
	// get_local
	s0i32 = l20
	// get_local
	s1i32 = l22
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l21
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l23
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l24 = s0i64
	// const
	s0i64 = 9222246136947933184
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
	s0i64 = 9222246136947933184
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
	// binary: i64.ne
	if s0i64 != s1i64 {
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
	// const
	s0i32 = 1
	// set_local
	l32 = s0i32
	// get_local
	s0i32 = l30
	// get_local
	s1i32 = l32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l33 = s0i32
	// get_local
	s0i32 = l31
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l33
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l34 = s0i64
	// const
	s0i32 = 9
	// set_local
	l35 = s0i32
	// get_local
	s0i64 = l34
	// get_local
	s1i32 = l35
	// call
	s0i32 = f311(ctx, s0i64, s1i32)
	// set_local
	l36 = s0i32
	// const
	s0i32 = 1
	// set_local
	l37 = s0i32
	// const
	s0i32 = 1
	// set_local
	l38 = s0i32
	// get_local
	s0i32 = l36
	// get_local
	s1i32 = l38
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l39 = s0i32
	// get_local
	s0i32 = l37
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l39
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// get_local
	s0i32 = l3
	// load: i64.load
	s0i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l40 = s0i64
	// const
	s0i32 = 10
	// set_local
	l41 = s0i32
	// get_local
	s0i64 = l40
	// get_local
	s1i32 = l41
	// call
	s0i32 = f311(ctx, s0i64, s1i32)
	// set_local
	l42 = s0i32
	// get_local
	s0i32 = l42
	// set_local
	l9 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l9
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
	// const
	s0i32 = 16
	// set_local
	l46 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l46
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l47 = s0i32
	// get_local
	s0i32 = l47
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l45
	// return
	return s0i32
}
