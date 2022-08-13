package internals

import (
	"encoding/binary"
)

func f93(ctx *Context, l0 int64) float64 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int64
	_ = l4
	var l5 float64
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
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
	s0f64 = f120(ctx, s0i64)
	// set_local
	l5 = s0f64
	// const
	s0i32 = 16
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l7
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0f64 = l5
	// return
	return s0f64
}
