package internals

import (
	"encoding/binary"
	"math"
)

func f479(ctx *Context, l0 int32, l1 int32) {
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
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 float64
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
	var l19 float64
	_ = l19
	var l20 float64
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 int32
	_ = l23
	var l24 float64
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	// get_global
	s0i32 = ctx.G0
	// set_local
	l2 = s0i32
	// const
	s0i32 = 32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l4
	// set_global
	ctx.G0 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l5 = s0i32
	// const
	s0i32 = 0
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l6
	// call
	s0i32 = f440(ctx, s0i32, s1i32)
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l8
	// call
	s0i32 = f481(ctx, s0i32)
	// set_local
	l9 = s0i32
	// get_local
	s0i32 = l9
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l10 = s0f64
	// const
	s0f64 = 2.097152e+06
	// set_local
	l11 = s0f64
	// get_local
	s0f64 = l10
	// get_local
	s1f64 = l11
	// binary: f64.mul
	s0f64 = s0f64 * s1f64
	// set_local
	l12 = s0f64
	// get_local
	s0i32 = l4
	// get_local
	s1f64 = l12
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// set_local
	l13 = s0i32
	// get_local
	s0i32 = l13
	// call
	s0i32 = f481(ctx, s0i32)
	// set_local
	l14 = s0i32
	// const
	s0i32 = 2097151
	// set_local
	l15 = s0i32
	// get_local
	s0i32 = l14
	// get_local
	s1i32 = l15
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l16 = s0i32
	// get_local
	s0i32 = l16
	// unary: f64.convert_u/i32
	s0f64 = float64(uint32(s0i32))
	// set_local
	l17 = s0f64
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l18 = s0f64
	// get_local
	s0f64 = l18
	// get_local
	s1f64 = l17
	// binary: f64.add
	s0f64 = s0f64 + s1f64
	// set_local
	l19 = s0f64
	// get_local
	s0i32 = l4
	// get_local
	s1f64 = l19
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l20 = s0f64
	// const
	s0f64 = 9.007199254740992e+15
	// set_local
	l21 = s0f64
	// get_local
	s0f64 = l20
	// get_local
	s1f64 = l21
	// binary: f64.div
	s0f64 = s0f64 / s1f64
	// set_local
	l22 = s0f64
	// get_local
	s0i32 = l4
	// get_local
	s1f64 = l22
	// store: f64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], math.Float64bits(s1f64))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// set_local
	l23 = s0i32
	// get_local
	s0i32 = l4
	// load: f64.load
	s0f64 = math.Float64frombits(binary.LittleEndian.Uint64(ctx.Mem[int(s0i32+8):]))
	// set_local
	l24 = s0f64
	// const
	s0i32 = 0
	// set_local
	l25 = s0i32
	// get_local
	s0i32 = l23
	// get_local
	s1i32 = l25
	// get_local
	s2f64 = l24
	// call
	f446(ctx, s0i32, s1i32, s2f64)
	// const
	s0i32 = 32
	// set_local
	l26 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l26
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l27 = s0i32
	// get_local
	s0i32 = l27
	// set_global
	ctx.G0 = s0i32
	// return
	return
}
