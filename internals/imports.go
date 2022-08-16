package internals

type ImportedFuncs interface {
	// module: wren, field: resolveModuleFn
	F0(ctx *Context, p0 int32, p1 int32, p2 int32) int32
	// module: wren, field: writeFn
	F1(ctx *Context, p0 int32, p1 int32)
	// module: wren, field: errorFn
	F2(ctx *Context, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32)
	// module: wren, field: heapSettings
	F3(ctx *Context, p0 int32) int32
	// module: wren, field: loadModuleFn
	F4(ctx *Context, p0 int32, p1 int32) int32
	// module: wren, field: bindForeignMethodFn
	F5(ctx *Context, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) int32
	// module: wren, field: dispatchForeignMethodFn
	F6(ctx *Context, p0 int32, p1 int32)
	// module: wren, field: bindForeignClassFn
	F7(ctx *Context, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32)
	// module: wren, field: dispatchFinalizerFn
	F8(ctx *Context, p0 int32, p1 int32, p2 int32)
	// module: wasi_snapshot_preview1, field: clock_time_get
	F9(ctx *Context, p0 int32, p1 int64, p2 int32) int32
	// module: wasi_snapshot_preview1, field: fd_close
	F10(ctx *Context, p0 int32) int32
	// module: wasi_snapshot_preview1, field: fd_fdstat_get
	F11(ctx *Context, p0 int32, p1 int32) int32
	// module: wasi_snapshot_preview1, field: fd_seek
	F12(ctx *Context, p0 int32, p1 int64, p2 int32, p3 int32) int32
	// module: wasi_snapshot_preview1, field: fd_write
	F13(ctx *Context, p0 int32, p1 int32, p2 int32, p3 int32) int32
	// module: wasi_snapshot_preview1, field: proc_exit
	F14(ctx *Context, p0 int32)
}
