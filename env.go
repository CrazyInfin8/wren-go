package wren

import (
	"io"
	"unsafe"

	"github.com/crazyinfin8/wren-go/internals"
)

type env struct {
	wasi_snapshot_preview1
	// Rather than have multiple VMs per wasm instance, use one wasm instance
	// per VM to make things simpler.
	ctx *internals.Context
	cfg Config
}

func (e *env) vmPtr() int32 {
	return e.ctx.GetVM()
}

// module: wren, field: resolveModuleFn
func (e *env) F0(ctx *internals.Context, vm int32, importer int32,
	name int32) int32 {
	newName, ok := e.cfg.ResolveModuleFn(VM{e}, gostring(ctx, importer),
		gostring(ctx, name))

	if !ok {
		return 0
	}

	// TODO: is this freed?
	return cstring(ctx, newName)
}

// module: wren, field: writeFn
func (e *env) F1(ctx *internals.Context, vm int32, msg int32) {
	e.cfg.WriteFn(VM{e}, gostring(ctx, msg))
}

// module: wren, field: errorFn
func (e *env) F2(ctx *internals.Context, vm int32, errType int32, mod int32,
	line int32, msg int32) {
	e.cfg.ErrorFn(VM{e}, WrenError{ErrorType(errType), gostring(ctx, mod),
		int(line), gostring(ctx, msg)})
}

// module: wren, field: loadModuleFn
func (e *env) F3(ctx *internals.Context, vm int32, name int32) int32 {
	src, ok := e.cfg.LoadModuleFn(VM{e}, gostring(ctx, name))
	if !ok {
		return 0
	}

	// Freed by shim.c
	return cstring(ctx, src)
}

// module: wren, field: bindForeignMethodFn
func (e *env) F4(ctx *internals.Context, vm int32, module int32,
	className int32, isStatic int32, signature int32) int32 {
	return e.cfg.BindForeignMethodFn(VM{e}, gostring(ctx, module),
		gostring(ctx, className), isStatic != 0, gostring(ctx, signature))
}

// module: wren, field: dispatchForeignMethodFn
func (e *env) F5(ctx *internals.Context, vm int32, userData int32) {
	e.cfg.DispatchForeignFn(VM{e}, userData)
}

// module: wren, field: bindForeignClassFn
func (e *env) F6(ctx *internals.Context, vm int32, module int32,
	className int32, allocatePtr int32, finalizePtr int32) {
	allocate, finalize := e.cfg.BindForeignClassMethodFn(VM{e},
		gostring(ctx, module), gostring(ctx, className))
	writeInt32(ctx, allocatePtr, allocate)
	writeInt32(ctx, finalizePtr, finalize)
}

// module: wren, field: dispatchFinalizerFn
func (e *env) F7(ctx *internals.Context, vm int32, data int32, userData int32) {
	e.cfg.DispatchFinalizerFn(VM{e}, data, userData)
}

// module: wasi_snapshot_preview1, field: clock_time_get
func (e *env) F8(ctx *internals.Context, p0 int32, p1 int64, p2 int32) int32 {
	return e.clock_time_get(ctx, p0, p1, p2)
}

// module: wasi_snapshot_preview1, field: fd_close
func (e *env) F9(ctx *internals.Context, p0 int32) int32 {
	return e.fd_close(ctx, p0)
}

// module: wasi_snapshot_preview1, field: fd_fdstat_get
func (e *env) F10(ctx *internals.Context, p0 int32, p1 int32) int32 {
	return e.fd_fdstat_get(ctx, p0, p1)
}

// module: wasi_snapshot_preview1, field: fd_read
func (e *env) F11(ctx *internals.Context, p0 int32, p1 int32, p2 int32,
	p3 int32) int32 {
	return e.fd_read(ctx, p0, p1, p2, p3)
}

// module: wasi_snapshot_preview1, field: fd_seek
func (e *env) F12(ctx *internals.Context, p0 int32, p1 int64, p2 int32,
	p3 int32) int32 {
	return e.fd_seek(ctx, p0, p1, p2, p3)
}

// module: wasi_snapshot_preview1, field: fd_write
func (e *env) F13(ctx *internals.Context, p0 int32, p1 int32, p2 int32,
	p3 int32) int32 {
	return e.fd_write(ctx, p0, p1, p2, p3)
}

// module: wasi_snapshot_preview1, field: proc_exit
func (e *env) F14(ctx *internals.Context, p0 int32) {
	e.proc_exit(ctx, p0)
}

// ExportEnvironment exports the memory tied to this wasm runtime to the given
// writer. In theory, this memory could be used inside of another runtime
// instance to essentially clone or restore the previous session
//
// TODO: Specify a binary format that will contain the memory as well as a list
// of foreign functions and/or objects so a wren VM can be cloned.
func (e *env) ExportEnvironment(out io.Writer) error {
	offset := 0
	for offset < len(e.ctx.Mem) {
		n, err := out.Write(e.ctx.Mem[offset:])
		if err != nil {
			return err
		}
		offset += n
	}
	return nil
}

func NewFromEnvironment(cfg Config, in io.Reader) (VM, error) {
	type Context struct {
		Mem []byte
		f   internals.ImportedFuncs
		G0  int32
		G1  int32
		G2  int32
		G3  int32
	}

	mem, err := io.ReadAll(in)
	if err != nil {
		return VM{}, err
	}

	env := new(env)
	ctx := (*internals.Context)((unsafe.Pointer)(&Context{mem, env,
		109840,
		22844,
		23328,
		39392,
	}))

	env.ctx = ctx
	env.cfg = cfg

	return VM{env}, nil
}
