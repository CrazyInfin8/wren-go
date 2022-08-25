package internals

import (
	"encoding/binary"
)

func f582(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l3 = s0i32
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
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
	s0i32 = l2
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
	// block
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l1
	// const
	s2i32 = 255
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l0
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l2
	// set_local
	l5 = s0i32
	// br
	goto lbl1
	// end_block
lbl4:
	// get_local
	s0i32 = l2
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
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
	s0i32 = l5
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
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l1
	// const
	s2i32 = 255
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = -2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
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
	s0i32 = l5
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
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l1
	// const
	s2i32 = 255
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = -3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
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
	s0i32 = l5
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
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l1
	// const
	s2i32 = 255
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l3 = s0i32
	// br
	goto lbl2
	// end_block
lbl3:
	// get_local
	s0i32 = l2
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l4 = s0i32
	// end_block
lbl2:
	// get_local
	s0i32 = l3
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
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l4
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l1
	// const
	s2i32 = 255
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 4
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 16843009
	// binary: i32.mul
	s0i32 = s0i32 * s1i32
	// set_local
	l0 = s0i32
	// loop
lbl6:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l0
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = -1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = -16843009
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = -2139062144
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = -4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// end
	// end_block
lbl5:
	// get_local
	s0i32 = l5
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
	s0i32 = l1
	// const
	s1i32 = 255
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l2 = s0i32
	// loop
lbl7:
	// block
	// get_local
	s0i32 = l4
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// get_local
	s1i32 = l2
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	// return
	return s0i32
	// end_block
lbl8:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// end
	// end_block
lbl0:
	// const
	s0i32 = 0
	// return
	return s0i32
}
