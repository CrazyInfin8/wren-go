package internals

import (
	"encoding/binary"
	"math"
)

func f448(ctx *Context, l0 int32, l1 int32, l2 float64) {
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
	var l8 float64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l3 = s0i32
	// const
	s0i32 = 16
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l5
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l5
	// get_local
	s1f64 = l2
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+0):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l5
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+0):]))
	// set_local
	l8 = s0f64
	// get_local
	s0f64 = l8
	// call
	s0i64 = f321(ctx, s0f64)
	// set_local
	l9 = s0i64
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l7
	// get_local
	s2i64 = l9
	// call
	f446(ctx, s0i32, s1i32, s2i64)
	// const
	s0i32 = 16
	// set_local
	l10 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l10
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l11 = s0i32
	// get_local
	s0i32 = l11
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
