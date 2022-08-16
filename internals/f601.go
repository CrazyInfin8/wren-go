package internals

import (
	"encoding/binary"
)

func f601(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// const
	s0i32 = 1
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l0
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
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 127
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
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
	// get_local
	s1i32 = l1
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// const
	s0i32 = 1
	// return
	return s0i32
	// end_block
lbl1:
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+43120):]))
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -128
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 57216
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 25
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42076):], uint32(s1i32))
	// br
	goto lbl2
	// end_block
lbl4:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// const
	s0i32 = 1
	// return
	return s0i32
	// end_block
lbl3:
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2047
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 6
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 192
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// const
	s0i32 = 2
	// return
	return s0i32
	// end_block
lbl5:
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = 55296
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -8192
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 57344
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// end_block
lbl7:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 12
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 224
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 6
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// const
	s0i32 = 3
	// return
	return s0i32
	// end_block
lbl6:
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -65536
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 1048575
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 18
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 240
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 6
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 12
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 63
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 128
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store8
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	// const
	s0i32 = 4
	// return
	return s0i32
	// end_block
lbl8:
	// const
	s0i32 = 0
	// const
	s1i32 = 25
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42076):], uint32(s1i32))
	// end_block
lbl2:
	// const
	s0i32 = -1
	// set_local
	l3 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l3
	// return
	return s0i32
}
