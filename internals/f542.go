package internals

import (
	"encoding/binary"
)

func f542(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// block
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// current_mem
	s0i32 = int32(len(ctx.Mem) / 65536)
	// const
	s1i32 = 16
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// return
	return s0i32
	// end_block
lbl0:
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 65535
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = -1
	// binary: i32.le_s
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// grow_mem
	s0i32 = int32(len(ctx.Mem) / 65536)
	ctx.growMem(int(s0i32) * 65536)
	// tee_local
	l0 = s0i32
	// const
	s1i32 = -1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl2
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// const
	s0i32 = -1
	// return
	return s0i32
	// end_block
lbl2:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// return
	return s0i32
	// end_block
lbl1:
	// call
	f541(ctx)
	// unreachable
	panic("unreachable executed")
}
