package internals

import (
	"encoding/binary"
	"math/bits"
)

func f534(ctx *Context, l0 int32) int32 {
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
	var s1i64 int64
	_ = s1i64
	// get_global
	s0i32 = ctx.G0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l1 = s0i32
	// set_global
	ctx.G0 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41596):]))
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 0
	// call
	s0i32 = f542(ctx, s0i32)
	// const
	s1i32 = 108704
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 89
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42044):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// const
	s0i32 = 0
	// const
	s1i64 = -1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+42056):], uint64(s1i64))
	// const
	s0i32 = 0
	// const
	s1i64 = 281474976776192
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+42048):], uint64(s1i64))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -16
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 1431655768
	// binary: i32.xor
	s1i32 = s1i32 ^ s2i32
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42044):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42064):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42016):], uint32(s1i32))
	// end_block
lbl1:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42024):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 108704
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42020):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 108704
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41588):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41608):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41604):], uint32(s1i32))
	// loop
lbl2:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41632
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41620
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41612
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41640
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41628
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41648
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41636
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41644
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 256
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
	// end
	// const
	s0i32 = 108704
	// const
	s1i32 = -8
	// const
	s2i32 = 108704
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// const
	s3i32 = 108704
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l3 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = -56
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42060):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41600):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// const
	s0i32 = 108704
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 56
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// end_block
lbl0:
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 236
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl14
	}
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41572):]))
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 16
	// get_local
	s2i32 = l0
	// const
	s3i32 = 19
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = -16
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// get_local
	s3i32 = l0
	// const
	s4i32 = 11
	// binary: i32.lt_u
	if uint32(s3i32) < uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l2 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l4 = s1i32
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
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
		goto lbl15
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l4
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 1
	// binary: i32.xor
	s0i32 = s0i32 ^ s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 41620
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l0
	// const
	s2i32 = 41612
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl17
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// br
	goto lbl16
	// end_block
lbl17:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// end_block
lbl16:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l5 = s1i32
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l4
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl15:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41580):]))
	// tee_local
	l7 = s1i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// block
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
		goto lbl18
	}
	// block
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 2
	// get_local
	s2i32 = l4
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l3
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 12
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 16
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 5
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l3
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 41620
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l0
	// const
	s2i32 = 41612
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl20
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l6 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// br
	goto lbl19
	// end_block
lbl20:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// end_block
lbl19:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l5 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l7
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// get_local
	s0i32 = l7
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41612
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41592):]))
	// set_local
	l4 = s0i32
	// block
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 1
	// get_local
	s2i32 = l8
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l8 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l8
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l8 = s0i32
	// br
	goto lbl22
	// end_block
lbl23:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l8 = s0i32
	// end_block
lbl22:
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl21:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41592):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41580):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl18:
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41576):]))
	// tee_local
	l9 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l9
	// const
	s1i32 = 0
	// get_local
	s2i32 = l9
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 12
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 16
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 5
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l3
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l0
	// set_local
	l5 = s0i32
	// block
	// loop
lbl25:
	// block
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl26
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl24
	}
	// end_block
lbl26:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l5
	// get_local
	s3i32 = l4
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// tee_local
	l5 = s2i32
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l5
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l3
	// set_local
	l5 = s0i32
	// br
	goto lbl25
	// end
	// end_block
lbl24:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l10 = s0i32
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l0
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl27
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41588):]))
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
	// tee_local
	l3 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl4
	// end_block
lbl27:
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl28
	}
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl28:
	// loop
lbl29:
	// get_local
	s0i32 = l5
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l3
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl29
	}
	// get_local
	s0i32 = l8
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl29
	}
	// end
	// get_local
	s0i32 = l11
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl4
	// end_block
lbl14:
	// const
	s0i32 = -1
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -65
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 19
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = -16
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41576):]))
	// tee_local
	l7 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// const
	s0i32 = 0
	// set_local
	l11 = s0i32
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 256
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl30
	}
	// const
	s0i32 = 31
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16777215
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl30
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 8
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1048320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l4
	// const
	s2i32 = 520192
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 245760
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 15
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// get_local
	s2i32 = l5
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l3
	// const
	s3i32 = 21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// end_block
lbl30:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l11
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// br
	goto lbl33
	// end_block
lbl34:
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l11
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l11
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// loop
lbl35:
	// block
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// get_local
	s0i32 = l6
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// const
	s0i32 = 0
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l5
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// set_local
	l3 = s0i32
	// br
	goto lbl32
	// end_block
lbl36:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l5
	// const
	s2i32 = 20
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l6 = s1i32
	// get_local
	s2i32 = l6
	// get_local
	s3i32 = l5
	// get_local
	s4i32 = l0
	// const
	s5i32 = 29
	// binary: i32.shr_u
	s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
	// const
	s5i32 = 4
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 16
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+0):]))
	// tee_local
	l5 = s3i32
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l6
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl35
	}
	// end
	// end_block
lbl33:
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl37
	}
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// const
	s0i32 = 2
	// get_local
	s1i32 = l11
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 0
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 12
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 16
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 5
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l3
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l0
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l5
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l3 = s0i32
	// end_block
lbl37:
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
		goto lbl31
	}
	// end_block
lbl32:
	// loop
lbl38:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l0 = s0i32
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl39
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l5 = s0i32
	// end_block
lbl39:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l0
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l0
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l5
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l5
	// br_if
	if s0i32 != 0 {
		goto lbl38
	}
	// end
	// end_block
lbl31:
	// get_local
	s0i32 = l8
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41580):]))
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l11 = s0i32
	// block
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l8
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl40
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41588):]))
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
	// tee_local
	l3 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl5
	// end_block
lbl40:
	// block
	// get_local
	s0i32 = l8
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl41
	}
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl11
	}
	// get_local
	s0i32 = l8
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// end_block
lbl41:
	// loop
lbl42:
	// get_local
	s0i32 = l5
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l3
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl42
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl42
	}
	// end
	// get_local
	s0i32 = l6
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl5
	// end_block
lbl13:
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41580):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl43
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41592):]))
	// set_local
	l4 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 16
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl45
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41580):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41592):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl44
	// end_block
lbl45:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41592):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41580):], uint32(s1i32))
	// end_block
lbl44:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// br
	goto lbl3
	// end_block
lbl43:
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41584):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl46
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41596):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l5 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// br
	goto lbl3
	// end_block
lbl46:
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42044):]))
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl48
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42052):]))
	// set_local
	l4 = s0i32
	// br
	goto lbl47
	// end_block
lbl48:
	// const
	s0i32 = 0
	// const
	s1i64 = -1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+42056):], uint64(s1i64))
	// const
	s0i32 = 0
	// const
	s1i64 = 281474976776192
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+42048):], uint64(s1i64))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 12
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -16
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 1431655768
	// binary: i32.xor
	s1i32 = s1i32 ^ s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42044):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42064):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42016):], uint32(s1i32))
	// const
	s0i32 = 65536
	// set_local
	l4 = s0i32
	// end_block
lbl47:
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// const
	s2i32 = 71
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l7 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l4
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l11 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl49
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl49:
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42012):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl50
	}
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42004):]))
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl51
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl50
	}
	// end_block
lbl51:
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 0
	// const
	s1i32 = 48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl50:
	// const
	s0i32 = 0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+42016)])
	// const
	s1i32 = 4
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// block
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41596):]))
	// tee_local
	l4 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl54
	}
	// const
	s0i32 = 42020
	// set_local
	l3 = s0i32
	// loop
lbl55:
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl53
	}
	// end_block
lbl56:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl55
	}
	// end
	// end_block
lbl54:
	// const
	s0i32 = 0
	// call
	s0i32 = f542(ctx, s0i32)
	// tee_local
	l0 = s0i32
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0i32 = l8
	// set_local
	l6 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42048):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l0
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
		goto lbl57
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l0
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l3
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// end_block
lbl57:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l2
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42012):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl58
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+42004):]))
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// end_block
lbl58:
	// get_local
	s0i32 = l6
	// call
	s0i32 = f542(ctx, s0i32)
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l0
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl52
	}
	// br
	goto lbl7
	// end_block
lbl53:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l11
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// get_local
	s0i32 = l6
	// call
	s0i32 = f542(ctx, s0i32)
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l3
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l0
	// set_local
	l3 = s0i32
	// end_block
lbl52:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl59
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 72
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl59
	}
	// block
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l6
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42052):]))
	// tee_local
	l4 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l4
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 2147483646
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl60
	}
	// get_local
	s0i32 = l3
	// set_local
	l0 = s0i32
	// br
	goto lbl7
	// end_block
lbl60:
	// block
	// get_local
	s0i32 = l4
	// call
	s0i32 = f542(ctx, s0i32)
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl61
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l3
	// set_local
	l0 = s0i32
	// br
	goto lbl7
	// end_block
lbl61:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// call
	s0i32 = f542(ctx, s0i32)
	// br
	goto lbl9
	// end_block
lbl59:
	// get_local
	s0i32 = l3
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l3
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
		goto lbl7
	}
	// br
	goto lbl9
	// end_block
lbl12:
	// const
	s0i32 = 0
	// set_local
	l8 = s0i32
	// br
	goto lbl4
	// end_block
lbl11:
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// br
	goto lbl5
	// end_block
lbl10:
	// get_local
	s0i32 = l0
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
		goto lbl7
	}
	// end_block
lbl9:
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42016):]))
	// const
	s2i32 = 4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42016):], uint32(s1i32))
	// end_block
lbl8:
	// get_local
	s0i32 = l8
	// const
	s1i32 = 2147483646
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
	// get_local
	s0i32 = l8
	// call
	s0i32 = f542(ctx, s0i32)
	// set_local
	l0 = s0i32
	// const
	s0i32 = 0
	// call
	s0i32 = f542(ctx, s0i32)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 56
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
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
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42004):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42004):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42008):]))
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl62
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42008):], uint32(s1i32))
	// end_block
lbl62:
	// block
	// block
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41596):]))
	// tee_local
	l4 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl66
	}
	// const
	s0i32 = 42020
	// set_local
	l3 = s0i32
	// loop
lbl67:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l5 = s1i32
	// get_local
	s2i32 = l3
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// tee_local
	l8 = s2i32
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl65
	}
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl67
	}
	// br
	goto lbl64
	// end
	// end_block
lbl66:
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41588):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl69
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl68
	}
	// end_block
lbl69:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41588):], uint32(s1i32))
	// end_block
lbl68:
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42024):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42020):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41604):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42044):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41608):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42032):], uint32(s1i32))
	// loop
lbl70:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41632
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41620
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41612
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41624
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41640
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41628
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41648
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 41636
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l4 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 41644
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 256
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl70
	}
	// end
	// get_local
	s0i32 = l0
	// const
	s1i32 = -8
	// get_local
	s2i32 = l0
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l0
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l3 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l6
	// const
	s2i32 = -56
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l5 = s1i32
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42060):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41600):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 56
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl63
	// end_block
lbl65:
	// get_local
	s0i32 = l3
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl64
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl64
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl64
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = -8
	// get_local
	s2i32 = l4
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l4
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l5 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41584):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l11 = s1i32
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l5 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42060):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41600):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l11
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 56
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl63
	// end_block
lbl64:
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41588):]))
	// tee_local
	l8 = s1i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl71
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41588):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// set_local
	l8 = s0i32
	// end_block
lbl71:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// const
	s0i32 = 42020
	// set_local
	l3 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// loop
lbl79:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l5
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl78
	}
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl79
	}
	// br
	goto lbl77
	// end
	// end_block
lbl78:
	// get_local
	s0i32 = l3
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	// const
	s1i32 = 8
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
		goto lbl76
	}
	// end_block
lbl77:
	// const
	s0i32 = 42020
	// set_local
	l3 = s0i32
	// loop
lbl80:
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl81
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl75
	}
	// end_block
lbl81:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l3 = s0i32
	// br
	goto lbl80
	// end
	// end_block
lbl76:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = -8
	// get_local
	s2i32 = l0
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l0
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l11 = s0i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// const
	s1i32 = -8
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l5
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l11
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l5 = s0i32
	// block
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l6
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl82
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41584):]))
	// get_local
	s2i32 = l5
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl73
	// end_block
lbl82:
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41592):]))
	// get_local
	s1i32 = l6
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl83
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41592):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41580):]))
	// get_local
	s2i32 = l5
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41580):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl73
	// end_block
lbl83:
	// block
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl84
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl86
	}
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 3
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l8 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 41612
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// block
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l4
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl87
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41572):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l8
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// br
	goto lbl85
	// end_block
lbl87:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl85
	// end_block
lbl86:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l9 = s0i32
	// block
	// block
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l6
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl89
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l6
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
	// tee_local
	l3 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl88
	// end_block
lbl89:
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl90
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl90
	}
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// br
	goto lbl88
	// end_block
lbl90:
	// loop
lbl91:
	// get_local
	s0i32 = l3
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l4
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl91
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl91
	}
	// end
	// get_local
	s0i32 = l8
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl88:
	// get_local
	s0i32 = l9
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl85
	}
	// block
	// block
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l6
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl93
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl92
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41576):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l4
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// br
	goto lbl85
	// end_block
lbl93:
	// get_local
	s0i32 = l9
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l9
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l6
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
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
		goto lbl85
	}
	// end_block
lbl92:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl94
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl94:
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl85
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl85:
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l6 = s0i32
	// end_block
lbl84:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l6
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l5
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl95
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41612
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41572):]))
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l4
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l4 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl97
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// set_local
	l4 = s0i32
	// br
	goto lbl96
	// end_block
lbl97:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l4 = s0i32
	// end_block
lbl96:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl73
	// end_block
lbl95:
	// const
	s0i32 = 31
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 16777215
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl98
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 8
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1048320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l4
	// const
	s2i32 = 520192
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l4 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l0
	// const
	s2i32 = 245760
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 15
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// get_local
	s2i32 = l0
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l3
	// const
	s3i32 = 21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// end_block
lbl98:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41576):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l8 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl99
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l8
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl73
	// end_block
lbl99:
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l3
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0i32
	// loop
lbl100:
	// get_local
	s0i32 = l0
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l5
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl74
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l8 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl100
	}
	// end
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl73
	// end_block
lbl75:
	// get_local
	s0i32 = l0
	// const
	s1i32 = -8
	// get_local
	s2i32 = l0
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 15
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l0
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 15
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// tee_local
	l3 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l11 = s0i32
	// get_local
	s1i32 = l6
	// const
	s2i32 = -56
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l8 = s1i32
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 56
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// const
	s2i32 = 55
	// get_local
	s3i32 = l5
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 15
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 0
	// get_local
	s4i32 = l5
	// const
	s5i32 = -55
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// const
	s5i32 = 15
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -63
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l8 = s1i32
	// get_local
	s2i32 = l8
	// get_local
	s3i32 = l4
	// const
	s4i32 = 16
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 35
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+42060):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41600):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+42028):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// get_local
	s0i32 = l8
	// const
	s1i32 = 0
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+42020):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l8
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42028):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42024):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42020):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42032):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// const
	s1i32 = 36
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// loop
lbl101:
	// get_local
	s0i32 = l3
	// const
	s1i32 = 7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l3
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl101
	}
	// end
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l4
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl63
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l4
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l6 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l6
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl102
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41612
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41572):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l5
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l5 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl104
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l5
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// set_local
	l5 = s0i32
	// br
	goto lbl103
	// end_block
lbl104:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l5 = s0i32
	// end_block
lbl103:
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl63
	// end_block
lbl102:
	// const
	s0i32 = 31
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16777215
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl105
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 8
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1048320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 520192
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l0
	// const
	s2i32 = 245760
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 15
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l5
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// get_local
	s2i32 = l0
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l3
	// const
	s3i32 = 21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// end_block
lbl105:
	// get_local
	s0i32 = l4
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41576):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l8 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l8
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl63
	// end_block
lbl106:
	// get_local
	s0i32 = l6
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l3
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l0 = s0i32
	// loop
lbl107:
	// get_local
	s0i32 = l0
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l6
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl72
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l8 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl107
	}
	// end
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl63
	// end_block
lbl74:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl73:
	// get_local
	s0i32 = l11
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// br
	goto lbl3
	// end_block
lbl72:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl63:
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41584):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41596):]))
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41584):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41596):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// br
	goto lbl3
	// end_block
lbl6:
	// const
	s0i32 = 0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 0
	// const
	s1i32 = 48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+42068):], uint32(s1i32))
	// br
	goto lbl3
	// end_block
lbl5:
	// block
	// get_local
	s0i32 = l11
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl108
	}
	// block
	// block
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l8
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+28):]))
	// tee_local
	l5 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 41876
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl110
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// br_if
	if s0i32 != 0 {
		goto lbl109
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l7
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l7 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// br
	goto lbl108
	// end_block
lbl110:
	// get_local
	s0i32 = l11
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l11
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l8
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
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
		goto lbl108
	}
	// end_block
lbl109:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l11
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l8
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl111
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl111:
	// get_local
	s0i32 = l8
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl108
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl108:
	// block
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 15
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl113
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl112
	// end_block
lbl113:
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l4
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 255
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl114
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41612
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41572):]))
	// tee_local
	l5 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l4
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l4 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl116
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// set_local
	l4 = s0i32
	// br
	goto lbl115
	// end_block
lbl116:
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l4 = s0i32
	// end_block
lbl115:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl112
	// end_block
lbl114:
	// const
	s0i32 = 31
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 16777215
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl117
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1048320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 520192
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l5 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 245760
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l2 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 15
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// get_local
	s1i32 = l3
	// get_local
	s2i32 = l5
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// get_local
	s2i32 = l2
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l3
	// const
	s3i32 = 21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// end_block
lbl117:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41876
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l5 = s0i32
	// block
	// get_local
	s0i32 = l7
	// const
	s1i32 = 1
	// get_local
	s2i32 = l3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l2 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl118
	}
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l7
	// get_local
	s2i32 = l2
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl112
	// end_block
lbl118:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l3
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l3
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l2 = s0i32
	// block
	// loop
lbl120:
	// get_local
	s0i32 = l2
	// tee_local
	l5 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l4
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl119
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l2
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l2 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// end
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl112
	// end_block
lbl119:
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl112:
	// get_local
	s0i32 = l8
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// br
	goto lbl3
	// end_block
lbl4:
	// block
	// get_local
	s0i32 = l10
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl121
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+28):]))
	// tee_local
	l5 = s1i32
	// const
	s2i32 = 2
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 41876
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl123
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// br_if
	if s0i32 != 0 {
		goto lbl122
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l9
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41576):], uint32(s1i32))
	// br
	goto lbl121
	// end_block
lbl123:
	// get_local
	s0i32 = l10
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l10
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l0
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl121
	}
	// end_block
lbl122:
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l10
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl124
	}
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl124:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl121
	}
	// get_local
	s0i32 = l8
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl121:
	// block
	// block
	// get_local
	s0i32 = l4
	// const
	s1i32 = 15
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl126
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// br
	goto lbl125
	// end_block
lbl126:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l5 = s0i32
	// get_local
	s1i32 = l4
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l2
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l7
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl127
	}
	// get_local
	s0i32 = l7
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41612
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41592):]))
	// set_local
	l3 = s0i32
	// block
	// block
	// const
	s0i32 = 1
	// get_local
	s1i32 = l8
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l8 = s0i32
	// get_local
	s1i32 = l6
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl129
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l8
	// get_local
	s2i32 = l6
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41572):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// set_local
	l8 = s0i32
	// br
	goto lbl128
	// end_block
lbl129:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l8 = s0i32
	// end_block
lbl128:
	// get_local
	s0i32 = l8
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl127:
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41592):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41580):], uint32(s1i32))
	// end_block
lbl125:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// end_block
lbl3:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l3
	// return
	return s0i32
}
