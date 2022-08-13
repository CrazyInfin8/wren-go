package internals

import (
	"encoding/binary"
	"math"
)

func f573(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s1f64 float64
	_ = s1f64
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
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l1
	// const
	s1i32 = -9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// br_table
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl18
	case 2:
		goto lbl17
	case 3:
		goto lbl14
	case 4:
		goto lbl16
	case 5:
		goto lbl15
	case 6:
		goto lbl13
	case 7:
		goto lbl12
	case 8:
		goto lbl11
	case 9:
		goto lbl10
	case 10:
		goto lbl9
	case 11:
		goto lbl8
	case 12:
		goto lbl7
	case 13:
		goto lbl6
	case 14:
		goto lbl5
	case 15:
		goto lbl4
	case 16:
		goto lbl3
	case 17:
		goto lbl2
	default:
		goto lbl0
	}
	// end_block
lbl18:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_s
	s1i64 = int64(int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):])))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl17:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_u
	s1i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl16:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_s
	s1i64 = int64(int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):])))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl15:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_u
	s1i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl14:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 7
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl13:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load16_s
	s1i64 = int64(int16(binary.LittleEndian.Uint16(ctx.Mem[int(s1i32+0):])))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl12:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load16_u
	s1i64 = int64(binary.LittleEndian.Uint16(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl11:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load8_s
	s1i64 = int64(int8(ctx.Mem[int(s1i32+0)]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl10:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load8_u
	s1i64 = int64(ctx.Mem[int(s1i32+0)])
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl9:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 7
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl8:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_u
	s1i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl7:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 7
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl6:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 7
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl5:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_s
	s1i64 = int64(int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):])))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl4:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i64.load32_u
	s1i64 = int64(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], uint64(s1i64))
	// return
	return
	// end_block
lbl3:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 7
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: f64.load
	s1f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// return
	return
	// end_block
lbl2:
	// call
	f575(ctx)
	// unreachable
	panic("unreachable executed")
	// end_block
lbl1:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 4
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl0:
}
