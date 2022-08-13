package wren

import (
	"encoding/binary"

	"github.com/crazyinfin8/wren-go/internals"
)

func writeUInt64(ctx *internals.Context, ptr int32, val uint64) {
	binary.LittleEndian.PutUint64(ctx.Mem[ptr:], val)
}

func writeUInt32(ctx *internals.Context, ptr int32, val uint32) {
	binary.LittleEndian.PutUint32(ctx.Mem[ptr:], val)
}

func readUInt32(ctx *internals.Context, ptr int32) uint32 {
	return binary.LittleEndian.Uint32(ctx.Mem[ptr:])
}

func writeInt32(ctx *internals.Context, ptr int32, val int32) {
	binary.LittleEndian.PutUint32(ctx.Mem[ptr:], uint32(val))
}

func readInt32(ctx *internals.Context, ptr int32) int32 {
	return int32(readUInt32(ctx, ptr))
}

func gostring(ctx *internals.Context, ptr int32) string {
	end := ptr
	c := ctx.Mem[end]
	for c != '\000' {
		end++
		c = ctx.Mem[end]
	}
	if end == ptr {
		return ""
	}
	return string(ctx.Mem[ptr:end])
}

func cstring(ctx *internals.Context, text string) int32 {
	ptr := ctx.Malloc(int32(len(text)) + 1)
	if ptr == 0 {
		panic("Out of memory (in wagu)")
	}
	copy(ctx.Mem[ptr:], []byte(text))
	ctx.Mem[int(ptr)+len(text)] = '\000'
	return ptr
}

func cbytes(ctx *internals.Context, data []byte) int32 {
	ptr := ctx.Malloc(int32(len(data)) + 1)
	if ptr == 0 {
		panic("Out of memory (in wagu)")
	}
	copy(ctx.Mem[ptr:], data)
	ctx.Mem[int(ptr)+len(data)] = '\000'
	return ptr
}

func boolToInt(value bool) int32 {
	if value {
		return 1
	}
	return 0
}
