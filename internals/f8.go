package internals

// module: wasi_snapshot_preview1, field: clock_time_get
func f8(ctx *Context, l0 int32, l1 int64, l2 int32) int32 {
	return ctx.f.F8(ctx, l0, l1, l2)
}