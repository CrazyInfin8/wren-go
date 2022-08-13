package internals

import (
	"encoding/binary"
	"math/bits"
)

func f538(ctx *Context, l0 int32, l1 int32) {
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
	var s1i64 int64
	_ = s1i64
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl1
	}
	// get_local
	s0i32 = l3
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
		goto lbl0
	}
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41680):]))
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.eq
	if s0i32 == s1i32 {
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
		goto lbl4
	}
	// get_local
	s0i32 = l0
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
	l5 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 41700
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l6 = s1i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// get_local
	s0i32 = l0
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
		goto lbl2
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41660):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41660):], uint32(s1i32))
	// br
	goto lbl1
	// end_block
lbl4:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l6 = s0i32
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
		goto lbl6
	}
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41676):]))
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
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl5
	// end_block
lbl6:
	// block
	// get_local
	s0i32 = l0
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
		goto lbl7
	}
	// get_local
	s0i32 = l0
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
		goto lbl7
	}
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// br
	goto lbl5
	// end_block
lbl7:
	// loop
lbl8:
	// get_local
	s0i32 = l3
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l4
	// tee_local
	l6 = s0i32
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
		goto lbl8
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l4 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl8
	}
	// end
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl5:
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
		goto lbl1
	}
	// block
	// block
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41964
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
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
		goto lbl10
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl9
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41664):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l4
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41664):], uint32(s1i32))
	// br
	goto lbl1
	// end_block
lbl10:
	// get_local
	s0i32 = l7
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l7
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
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l6
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
	// end_block
lbl9:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
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
		goto lbl11
	}
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl11:
	// get_local
	s0i32 = l0
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
		goto lbl1
	}
	// get_local
	s0i32 = l6
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
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// br
	goto lbl1
	// end_block
lbl3:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 3
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	// get_local
	s1i32 = l3
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41668):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// return
	return
	// end_block
lbl2:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
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
	// end_block
lbl1:
	// block
	// block
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 2
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl13
	}
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41684):]))
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
		goto lbl14
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41684):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41672):]))
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41672):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41680):]))
	// binary: i32.ne
	if s0i32 != s1i32 {
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
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41668):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41680):], uint32(s1i32))
	// return
	return
	// end_block
lbl14:
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41680):]))
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
		goto lbl15
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41680):], uint32(s1i32))
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41668):]))
	// get_local
	s2i32 = l1
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41668):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// return
	return
	// end_block
lbl15:
	// get_local
	s0i32 = l3
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
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
		goto lbl17
	}
	// get_local
	s0i32 = l2
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
	l5 = s1i32
	// const
	s2i32 = 3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 41700
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l6 = s1i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// block
	// get_local
	s0i32 = l2
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
		goto lbl18
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41660):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l5
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41660):], uint32(s1i32))
	// br
	goto lbl16
	// end_block
lbl18:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
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
	goto lbl16
	// end_block
lbl17:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l7 = s0i32
	// block
	// block
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.eq
	if s0i32 == s1i32 {
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
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41676):]))
	// get_local
	s1i32 = l2
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
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// br
	goto lbl19
	// end_block
lbl20:
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl21
	}
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// br
	goto lbl19
	// end_block
lbl21:
	// loop
lbl22:
	// get_local
	s0i32 = l4
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l3
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl22
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l6
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l3 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl22
	}
	// end
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl19:
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
		goto lbl16
	}
	// block
	// block
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41964
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
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
		goto lbl24
	}
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// br_if
	if s0i32 != 0 {
		goto lbl23
	}
	// const
	s0i32 = 0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41664):]))
	// const
	s2i32 = -2
	// get_local
	s3i32 = l4
	// binary: i32.rotl
	s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41664):], uint32(s1i32))
	// br
	goto lbl16
	// end_block
lbl24:
	// get_local
	s0i32 = l7
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l7
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l2
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
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l6
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl16
	}
	// end_block
lbl23:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// block
	// get_local
	s0i32 = l2
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
		goto lbl25
	}
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl25:
	// get_local
	s0i32 = l2
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
		goto lbl16
	}
	// get_local
	s0i32 = l6
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
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl16:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+41680):]))
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl12
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41668):], uint32(s1i32))
	// return
	return
	// end_block
lbl13:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// end_block
lbl12:
	// block
	// get_local
	s0i32 = l1
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
		goto lbl26
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 3
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 41700
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41660):]))
	// tee_local
	l4 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l3
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l3 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl28
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41660):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// set_local
	l3 = s0i32
	// br
	goto lbl27
	// end_block
lbl28:
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l3 = s0i32
	// end_block
lbl27:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// return
	return
	// end_block
lbl26:
	// const
	s0i32 = 31
	// set_local
	l3 = s0i32
	// block
	// get_local
	s0i32 = l1
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
		goto lbl29
	}
	// get_local
	s0i32 = l1
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
	l6 = s0i32
	// get_local
	s1i32 = l6
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
	l6 = s1i32
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
	s2i32 = l6
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
	s1i32 = l1
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
lbl29:
	// get_local
	s0i32 = l0
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l0
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
	s1i32 = 41964
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l4 = s0i32
	// block
	// const
	s0i32 = 0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41664):]))
	// tee_local
	l6 = s0i32
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
		goto lbl30
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// get_local
	s2i32 = l2
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+41664):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
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
	// return
	return
	// end_block
lbl30:
	// get_local
	s0i32 = l1
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
	l6 = s0i32
	// block
	// loop
lbl32:
	// get_local
	s0i32 = l6
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl31
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l6 = s0i32
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
	s1i32 = l6
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
	l2 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l6 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl32
	}
	// end
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
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
	// return
	return
	// end_block
lbl31:
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl0:
}
