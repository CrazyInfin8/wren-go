package wren

import (
	"io"
	"time"

	"github.com/crazyinfin8/wren-go/internals"
)

type env struct {
	wasi_snapshot_preview1
	// Rather than have multiple VMs per wasm instance, use one wasm instance
	// per VM to make things simpler.
	ctx *internals.Context
	cfg Config

	// Keep Go foreign data and functions in Go.
	objs       map[int32]any
	fns        map[int32]func(VM)
	finalizers map[int32]func(VM, interface{})

	startTime time.Time
	UserData any
}

func (e *env) vmPtr() int32 {
	return e.ctx.GetVM()
}

// module: wren, field: resolveModuleFn
func (e *env) F0(ctx *internals.Context, vm int32, importer int32,
	name int32) int32 {
	if e.cfg.ResolveModuleFn == nil {
		return name
	}
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
	if e.cfg.WriteFn == nil {
		return
	}
	e.cfg.WriteFn(VM{e}, gostring(ctx, msg))
}

// module: wren, field: errorFn
func (e *env) F2(ctx *internals.Context, vm int32, errType int32, mod int32,
	line int32, msg int32) {
	if e.cfg.ErrorFn == nil {
		return
	}
	e.cfg.ErrorFn(VM{e}, WrenError{ErrorType(errType), gostring(ctx, mod),
		int(line), gostring(ctx, msg)})
}

// module: wren, field: heapSettings
func (e *env) F3(ctx *internals.Context, p0 int32) int32 {
	if e.cfg.InitialHeapSize != 0 {
		writeUInt32(ctx, p0, uint32(e.cfg.InitialHeapSize))
	}
	if e.cfg.MinHeapSize != 0 {
		writeUInt32(ctx, p0, uint32(e.cfg.MinHeapSize))
	}
	if e.cfg.HeapGrowthPercent != 0 {
		writeInt32(ctx, p0, int32(e.cfg.HeapGrowthPercent))
	}
	return 0
}

// module: wren, field: loadModuleFn
func (e *env) F4(ctx *internals.Context, vm int32, name int32) int32 {
	if e.cfg.LoadModuleFn == nil {
		return 0
	}
	src, ok := e.cfg.LoadModuleFn(VM{e}, gostring(ctx, name))
	if !ok {
		return 0
	}

	// Freed by shim.c
	return cstring(ctx, src)
}

// module: wren, field: bindForeignMethodFn
func (e *env) F5(ctx *internals.Context, vm int32, module int32,
	className int32, isStatic int32, signature int32) int32 {
	if e.cfg.BindForeignMethodFn == nil {
		return 0
	}

	fn := e.cfg.BindForeignMethodFn(VM{e}, gostring(ctx, module),
		gostring(ctx, className), gostring(ctx, signature), isStatic != 0)
	if fn == nil {
		return 0
	}
	// We don't really need to malloc, but it is an easy way to get a unique
	// int32. We don't need to free it since foreign functions don't get
	// discarded, and each environment has it's own set of memory.
	ptr := ctx.Malloc(1)
	e.fns[ptr] = fn
	return ptr
}

// module: wren, field: dispatchForeignMethodFn
func (e *env) F6(ctx *internals.Context, vm int32, userData int32) {
	e.fns[userData](VM{e})
}

// module: wren, field: bindForeignClassFn
func (e *env) F7(ctx *internals.Context, vm int32, module int32,
	className int32, allocatePtr int32, finalizePtr int32) {
	writeInt32(ctx, allocatePtr, 0)
	writeInt32(ctx, finalizePtr, 0)
	if e.cfg.BindForeignClassMethodFn == nil {
		return
	}
	allocate, finalize := e.cfg.BindForeignClassMethodFn(VM{e},
		gostring(ctx, module), gostring(ctx, className))
	if allocate == nil {
		return
	}
	// We don't really need to malloc, but it is an easy way to get a unique
	// int32. We don't need to free it since foreign functions don't get
	// discarded, and each environment has it's own set of memory.
	ptr := ctx.Malloc(1)
	e.fns[ptr] = allocate
	writeInt32(ctx, allocatePtr, ptr)
	if finalize != nil {
		e.finalizers[ptr] = finalize
		writeInt32(ctx, finalizePtr, ptr)
	}
}

// module: wren, field: dispatchFinalizerFn
func (e *env) F8(ctx *internals.Context, vm int32, data int32, userData int32) {
	e.finalizers[userData](VM{e}, e.objs[data])
	// Remove reference to foreign object so it can be garbage collected by Go
	delete(e.objs, data)
}

// module: wren, field: clock
func (e *env)F9(ctx *internals.Context) float64 {
	return time.Since(e.startTime).Seconds()
}

// module: wren, field: seed
func (e *env) F10(ctx *internals.Context) int32 {
	return int32(time.Since(e.startTime).Nanoseconds())
}

// module: wasi_snapshot_preview1, field: fd_close
func (e *env) F11(ctx *internals.Context, p0 int32) int32 {
	return e.fd_close(ctx, p0)
}

// module: wasi_snapshot_preview1, field: fd_fdstat_get
func (e *env) F12(ctx *internals.Context, p0 int32, p1 int32) int32 {
	return e.fd_fdstat_get(ctx, p0, p1)
}

// module: wasi_snapshot_preview1, field: fd_seek
func (e *env) F13(ctx *internals.Context, p0 int32, p1 int64, p2 int32,
	p3 int32) int32 {
	return e.fd_seek(ctx, p0, p1, p2, p3)
}

// module: wasi_snapshot_preview1, field: fd_write
func (e *env) F14(ctx *internals.Context, p0 int32, p1 int32, p2 int32,
	p3 int32) int32 {
	return e.fd_write(ctx, p0, p1, p2, p3)
}

// module: wasi_snapshot_preview1, field: proc_exit
func (e *env) F15(ctx *internals.Context, p0 int32) {
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

// TODO: Enable this function if exporting and importing runtime memory is
// feasible
//
// func NewFromEnvironment(cfg Config, in io.Reader) (VM, error) {
// 	type Context struct {
// 		Mem []byte
// 		f   internals.ImportedFuncs
// 		G0  int32
// 		G1  int32
// 		G2  int32
// 		G3  int32
// 	}

// 	mem, err := io.ReadAll(in)
// 	if err != nil {
// 		return VM{}, err
// 	}

// 	env := new(env)
// 	ctx := (*internals.Context)((unsafe.Pointer)(&Context{mem, env,
// 		109840,
// 		22844,
// 		23328,
// 		39392,
// 	}))

// 	env.ctx = ctx
// 	env.cfg = cfg

// 	return VM{env}, nil
// }
