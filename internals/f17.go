package internals

import (
	"encoding/binary"
)

func f17(ctx *Context) int32 {
	var l0 int32
	_ = l0
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+41568):]))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l1
	// return
	return s0i32
}
