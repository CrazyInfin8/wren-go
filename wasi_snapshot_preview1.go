package wren

import (
	"io"

	"github.com/crazyinfin8/wren-go/internals"
)

// wasi_snapshot_preview1 provide the wasi functions necessary to run wren.wasm.
//
// It currently provides only the bare minimum necessary to run wren but is not
// meant to be correct or feature complete.
type wasi_snapshot_preview1 struct {
}

type clock_id int32

const (
	clockIDRealtime clock_id = iota
	clockIDMonotonic
	clockIDProcessCputime
	clockIDThreadCputime
)

// StdOut is where output from any context is printed to.
var StdOut io.Writer

// StdErr is where error messages from any context is printed to. This may
// contain debug information from the incorrect use of wren.
var StdErr io.Writer

func (wasi_snapshot_preview1) clock_time_get(ctx *internals.Context, id int32,
	precision int64, resultTimestamp int32) (errno int32) {
	if id == int32(clockIDMonotonic) {
		writeUInt64(ctx, resultTimestamp, 0)
	}
	return 0
}

const (
	fdStdin = iota
	fdStdout
	fdStderr
)

func (wasi_snapshot_preview1) fd_close(ctx *internals.Context,
	fd int32) (errno int32) {
	panic("Unimplemented")
}

func (wasi_snapshot_preview1) fd_fdstat_get(ctx *internals.Context, fd int32,
	results int32) (errno int32) {
	switch fd {
	case fdStdin, fdStdout, fdStderr:
		return 0
	default:
		panic("Unimplemented")
	}
}

func (wasi_snapshot_preview1) fd_seek(ctx *internals.Context, fd int32,
	offset int64, whence int32, newOffset int32) (errno int32) {
	panic("Unimplemented")
}

func (wasi_snapshot_preview1) fd_write(ctx *internals.Context, fd int32,
	iovs int32, count int32, sizePtr int32) (errno int32) {
	var out io.Writer
	switch fd {
	case fdStdout:
		out = StdOut
	case fdStderr:
		out = StdErr
	default:
		panic("Unimplemented")
	}
	if out == nil {
		return 0
	}
	size := uint32(0)
	for i := int32(0); i < count; i++ {
		offset := readUInt32(ctx, iovs)
		length := readUInt32(ctx, iovs+4)
		n, err := out.Write(ctx.Mem[offset : offset+length])
		if err != nil {
			writeUInt32(ctx, sizePtr, size)
			return 5
		}
		iovs += 4 * 2
		size += uint32(n)
	}
	writeUInt32(ctx, sizePtr, size)
	return 0
}

func (wasi_snapshot_preview1) proc_exit(ctx *internals.Context, code int32) {
	panic("Unimplemented")
}
