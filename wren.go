package wren

//go:generate go run build-wren.go

import (
	"io"
	"strconv"
	"strings"

	"github.com/crazyinfin8/wren-go/internals"
)

type Config struct {
	// ResolveModuleFn allows the host to canonicalize the imported name beofre
	// it is passed to LoadModuleFn.
	//
	// It takes the name of the module doing the import and the target module to
	// import. It should return the new name to be passed to LoadModuleFn as
	// well as if the import is ok. If an import cannot be resolved, this
	// function should return false.
	ResolveModuleFn func(vm VM, importer, name string) (newName string, ok bool)

	// LoadModuleFn load a modules source code. It takes the name of the module
	// to import and returns the wren source code as well as whether the module
	// was successfully loaded.
	LoadModuleFn func(vm VM, name string) (src string, ok bool)

	// BindForeignMethodFn binds a foreign function to a class method in wren.
	//
	// It is passed the module, class name, and call signature of the function
	// as well as whether the function is static.
	//
	// Returning nil signifies that such function does not exist, raising an
	// error in wren
	//
	// The call signature's format varies between the type of function it is but
	// usually begins with the function name and ends with the amount of
	// parameters indicated as underscores.
	//
	//     - Regular functions enclose the parameters in parenthesis:
	//       "function(_,_,_)"
	//     - Getters do not have anything following the name: "function"
	//     - Setters have an equal sign before the parameters surrounded by
	//       parenthesis (and should only have one parameter): "function=(_)"
	//     - Indexes have no names but parameters surrounded by brackets:
	//       "[_,_,_]"
	//     - Index setters also have no name but parameters surrounded by
	//       brackets. It also has an equal sign and one parameter surrounded by
	//       parenthesis: "[_,_,_]=(_)"
	//     - Operators begin with the operation followed by one parameter
	//       surrounded by parenthesis: "+(_)", ">=(_)", "<<(_)"
	//
	BindForeignMethodFn func(vm VM, module, class,
		signature string, static bool) func(VM)

	// BindForeignClassMethodFn binds the allocator and finalizer of a class.
	//
	// It takes the name of the class and it's module and should return two
	// functions: one called during the creation of a foreign object (when a
	// constructor is called) and one called when the foreign object is garbage
	// collected. Setting the allocator to nil signifies an error to wren but
	// finalizer is optional and can be set to nil to indicate no finalizer
	// function.
	BindForeignClassMethodFn func(vm VM, module, class string) (allocator func(VM),
		finalizer func(VM, interface{}))

	// WriteFn is called when "System.write", "System.print", and related
	// functions have been called in wren.
	WriteFn func(vm VM, message string)

	// ErrorFn is called whenever acompilation or runtime error has occurred in
	// wren. The error value should always be of type WrenError.
	ErrorFn func(vm VM, err error)

	// InitialHeapSize is the number of bytes wren will allocate before
	// triggering the first garbage collection.
	//
	// Set this to zero to use wren's default value
	InitialHeapSize uint

	// MinHeapSize is the minimum size threshold after a garbage collection has
	// occured. This ensures that the heap does not get to small.
	//
	// Set this to zero to use wren's default value
	MinHeapSize uint

	// HeapGrowthPercent determines the amount of additional memory the heap
	// will grow to after garbage collection. This is specified as a percent of
	// the current heap size with a value like 50 increasing the heap size by
	// 50%
	//
	// The smaller this value is means the less memory wasted but the garbage
	// collector may be triggered more often
	//
	// Set this to zero to use wren's default value
	HeapGrowthPercent int
}

// WrenError represents a compile or runtime error from wren. Depending on the
// ErrorType, not all field may be populated
type WrenError struct {
	Type    ErrorType
	Module  string
	Line    int
	Message string
}

func (err WrenError) Error() string {
	switch err.Type {
	case ErrorCompile, ErrorStackTrace:
		// Don't use fmt if we don't need to.
		buf := strings.Builder{}
		// Based on: fprintf(stderr, "[%s line %d] %s\n", module, line, message)
		buf.WriteByte('[')
		buf.WriteString(err.Module)
		buf.WriteString(" line ")
		buf.WriteString(strconv.FormatInt(int64(err.Line), 10))
		buf.WriteString("] ")
		buf.WriteString(err.Message)
		return buf.String()
	default:
		return err.Message
	}
}

// VM interacts with the wren virtual machine.
type VM struct {
	env *env
}

func (vm VM) ctx() *internals.Context {
	return vm.env.ctx
}

func (vm VM) ptr() int32 { return vm.env.vmPtr() }

func (vm VM) VersionNumber() int {
	return int(vm.ctx().WrenGetVersionNumber())
}

func (vm VM) VersionTuple() [3]int {
	ctx := vm.ctx()
	ptr := ctx.G1
	return [3]int{
		int(readInt32(ctx, ptr)),
		int(readInt32(ctx, ptr+4)),
		int(readInt32(ctx, ptr+8)),
	}
}

func (vm VM) VersionString() string {
	ctx := vm.ctx()
	return gostring(ctx, ctx.G2)
}

// Free disposes all data used by the VM. This isn't always necessary and won't
// immediately free it's data because the VM's memory is still managed by Go and
// must be garbage collected by Go. However, this function garbage collects any
// remaining foreign objects existing in the VM to ensure their finalizers are
// called.
//
// This VM should no longer be used after calling this function.
func (vm VM) Free() {
	vm.ctx().WrenFreeVM(vm.ptr())
}

// CollectGarbage runs the garbage collector for wren.
func (vm VM) CollectGarbage() {
	vm.ctx().WrenCollectGarbage(vm.ptr())
}

// Exit attempts to interrupt wren while it is running and exit. The VM has not
// been tested to verify that it's state is usable afterwards so it is
// recommended to free it, though your mileage may vary.
func (vm VM) Exit() {
	vm.ctx().WrenEarlyExit(vm.ptr())
}

// Allocated returns the amount of bytes allocated by the wren VM. This does not
// include foreign objects or functions as those are still handled by Go and it
// doesn't reflect the entire amount of memory needed to run this instance of VM
// but it is useful to compare against the original version of wren.
func (vm VM) Allocated() int {
	return int(vm.ctx().WrenGetAllocated(vm.ptr()))
}

// Interpret runs the source code in a new fiber in the resolved module.
//
// Note: that wren is not reentrant. You should not call this function from a
// foreign method or while the VM is currently running.
func (vm VM) Interpret(module, source string) InterpretResult {
	ctx := vm.ctx()

	cmod := cstring(ctx, module)
	defer ctx.Free(cmod)

	csrc := cstring(ctx, source)
	defer ctx.Free(csrc)

	return InterpretResult(ctx.WrenInterpret(vm.ptr(), cmod, csrc))
}

// NewCallHandle creates a handle used to invoke methods in wren. When the
// handle is no longer needed, you should free the handle using ReleaseHandle.
//
// The Handle is reusable and is not locked to a specific receiver. Instead the
// receiver must be added to the callstack by setting slot 0.
//
// Note: that wren is not reentrant. You should not call this handle from a
// foreign method or while the VM is currently running.
//
// The call signature's format varies between the type of function it is but
// usually begins with the function name and ends with the amount of parameters
// indicated as underscores.
//
//     - Regular functions enclose the parameters in parenthesis:
//       "function(_,_,_)"
//     - Getters do not have anything following the name: "function"
//     - Setters have an equal sign before the parameters surrounded by
//       parenthesis (and should only have one parameter): "function=(_)"
//     - Indexes have no names but parameters surrounded by brackets: "[_,_,_]"
//     - Index setters also have no name but parameters surrounded by brackets.
//       It also has an equal sign and one parameter surrounded by parenthesis:
//       "[_,_,_]=(_)"
//     - Operators begin with the operation followed by one parameter surrounded
//       by parenthesis: "+(_)", ">=(_)", "<<(_)"
//
func (vm VM) NewCallHandle(signature string) Handle {
	ctx := vm.ctx()

	csig := cstring(ctx, signature)
	defer ctx.Free(csig)

	hPtr := ctx.WrenMakeCallHandle(vm.ptr(), csig)

	return Handle{vm, hPtr}
}

// CallHandle calls the handle, using the receiver in slot 0 and arguments from
// slot 1 or greater.
//
// The handle should have been created using NewCallHandle.
//
// Note: that wren is not reentrant. You should not call this function from a
// foreign method or while the VM is currently running.
func (vm VM) Call(handle Handle) InterpretResult {
	return handle.Call()
}

// ReleaseHandle releases and frees up this handle. Handles prevent values from
// being garbage collected so releasing a handle removes this referance to that
// value, allowing it to be garbage collected when all references are gone.
func (vm VM) ReleaseHandle(handle Handle) {
	ctx := vm.ctx()
	if ctx != handle.vm.ctx() {
		panic("handle's context does not match VM")
	}
	ctx.WrenReleaseHandle(vm.ptr(), handle.ptr)
}

// SlotCount returns the numbers of slots currently available to the current
// foreign method.
func (vm VM) SlotCount() int {
	return int(vm.ctx().WrenGetSlotCount(vm.ptr()))
}

// EnsureSlots ensures that the callstack has enough slots to accomadate the
// value passed.
//
// Each slot coorosponds to one receiver or parameter
func (vm VM) EnsureSlots(numSlots int) {
	vm.ctx().WrenEnsureSlots(vm.ptr(), int32(numSlots))
}

// SlotType gets the value type located in the slot provided.
func (vm VM) SlotType(slot int) Type {
	return Type(vm.ctx().WrenGetSlotType(vm.ptr(), int32(slot)))
}

// GetBool retreives a bool in the slot provided.
func (vm VM) GetBool(slot int) bool {
	return vm.ctx().WrenGetSlotBool(vm.ptr(), int32(slot)) != 0
}

// GetBytes retrieves string data from wren and writes it into writer. Error
// values from writer are returned.
func (vm VM) GetBytes(slot int, writer io.Writer) error {
	ctx := vm.ctx()

	sizePtr := ctx.Malloc(4)
	defer ctx.Free(sizePtr)

	dataPtr := ctx.WrenGetSlotBytes(vm.ptr(), int32(slot), sizePtr)
	endPtr := dataPtr + readInt32(ctx, sizePtr)

	for dataPtr < endPtr {
		n, err := writer.Write(ctx.Mem[dataPtr:endPtr])
		if err != nil {
			return err
		}
		dataPtr += int32(n)
	}
	return nil
}

// GetNum returns a number in the slot provided.
func (vm VM) GetNum(slot int) float64 {
	return vm.ctx().WrenGetSlotDouble(vm.ptr(), int32(slot))
}

// GetForeign returns the foreign object referenced in the current slot. It's
// value reflects what was passed from SetNewForeign.
func (vm VM) GetForeign(slot int) any {
	ptr := vm.ctx().WrenGetSlotForeign(vm.ptr(), int32(slot))
	return vm.env.objs[ptr]
}

// GetString copies a string from the slot provided and returns it.
func (vm VM) GetString(slot int) string {
	buf := strings.Builder{}
	vm.GetBytes(slot, &buf)

	return buf.String()
}

// GetSlotHandle creates a new handle from the value in the slot provided.
// Handles prevent a value from being garbage collected.
func (vm VM) GetHandle(slot int) Handle {
	return Handle{vm, vm.ctx().WrenGetSlotHandle(vm.ptr(), int32(slot))}
}

// SetSlotBool sets the provided slot to a boolean value.
func (vm VM) SetBool(slot int, value bool) {
	vm.ctx().WrenSetSlotBool(vm.ptr(), int32(slot), boolToInt(value))
}

// SetSlotBytes sets the provided slot to a string value from the given bytes.
func (vm VM) SetBytes(slot int, data []byte) {
	ctx := vm.ctx()

	cbuf := cbytes(ctx, data)
	defer ctx.Free(cbuf)

	ctx.WrenSetSlotBytes(vm.ptr(), int32(slot), cbuf, int32(len(data)))
}

// SetNum sets the provided slot to a number value from the given
// float64.
func (vm VM) SetNum(slot int, value float64) {
	vm.ctx().WrenSetSlotDouble(vm.ptr(), int32(slot), value)
}

// SetNewForeign creates a new foreign object in wren.
func (vm VM) SetNewForeign(slot, classSlot int, obj interface{}) {
	ID := vm.ctx().WrenSetSlotNewForeign(vm.ptr(), int32(slot), int32(classSlot), 1)
	vm.env.objs[ID] = obj
}

// SetNewList creates an empty list and sets it in the slot provided.
func (vm VM) SetNewList(slot int) {
	vm.ctx().WrenSetSlotNewList(vm.ptr(), int32(slot))
}

// SetNewMap creates an empty map and sets it in the slot provided.
func (vm VM) SetNewMap(slot int) {
	vm.ctx().WrenSetSlotNewMap(vm.ptr(), int32(slot))
}

// SetNull sets the provided slot to a null value.
func (vm VM) SetNull(slot int) {
	vm.ctx().WrenSetSlotNull(vm.ptr(), int32(slot))
}

// SetString sets the provided slot to the given string value.
func (vm VM) SetString(slot int, value string) {
	ctx := vm.ctx()

	cstr := cstring(ctx, value)
	defer ctx.Free(cstr)

	ctx.WrenSetSlotBytes(vm.ptr(), int32(slot), cstr, int32(len(value)))
}

// SetHandle sets the provided slot to the value held by the given handle.
func (vm VM) SetHandle(slot int, handle Handle) {
	ctx := vm.ctx()
	if ctx != handle.vm.ctx() {
		panic("handle's context does not match VM")
	}
	ctx.WrenSetSlotHandle(vm.ptr(), int32(slot), handle.ptr)
}

// ListCount returns the list of the list in the given slot.
func (vm VM) ListCount(slot int) int {
	return int(vm.ctx().WrenGetListCount(vm.ptr(), int32(slot)))
}

// GetListElement places an element at the specified index from a list located
// in listSlot into elementSlot.
func (vm VM) GetListElement(listSlot, index, elementSlot int) {
	vm.ctx().WrenGetListElement(vm.ptr(), int32(listSlot), int32(index),
		int32(elementSlot))
}

// SetListElement sets a lists element at the specified index from a list
// locates in listSLot to the value locates at elementSlot.
func (vm VM) SetListElement(listSlot, index, elementSlot int) {
	vm.ctx().WrenSetListElement(vm.ptr(), int32(listSlot), int32(index),
		int32(elementSlot))
}

// InsertInList inserts an element into the list at the specified index.
//
// Negative values are supported to index an array from the end.
func (vm VM) InsertInList(listSlot, index, elementSlot int) {
	vm.ctx().WrenInsertInList(vm.ptr(), int32(listSlot), int32(index),
		int32(elementSlot))
}

// MapCount returns the amount of elements in the map at the specified slot.
func (vm VM) MapCount(slot int) int {
	return int(vm.ctx().WrenGetMapCount(vm.ptr(), int32(slot)))
}

// MapHasKey returns true if the value in keySlot is a valid key in this
// map
func (vm VM) MapHasKey(mapSlot, keySlot int) bool {
	return vm.ctx().WrenGetMapContainsKey(vm.ptr(), int32(mapSlot),
		int32(keySlot)) != 0
}

// GetMapValue puts the element at index of keySlot from the map in mapSlot into
// valueSlot.
func (vm VM) GetMapValue(mapSlot, keySlot, valueSlot int) {
	vm.ctx().WrenGetMapValue(vm.ptr(), int32(mapSlot), int32(keySlot),
		int32(valueSlot))
}

// SetMapValue sets the element at index of keySlot from the map in mapSlot into
// valueSlot.
func (vm VM) SetMapValue(mapSlot, keySlot, valueSlot int) {
	vm.ctx().WrenSetMapValue(vm.ptr(), int32(mapSlot), int32(keySlot),
		int32(valueSlot))
}

// RemoveMapValue removes an element from the map. If the element was found, it
// is put into valueSlot. If not, then valueSlot will be set to null.
func (vm VM) RemoveMapValue(mapSlot, keySlot, valueSlot int) {
	vm.ctx().WrenRemoveMapValue(vm.ptr(), int32(mapSlot), int32(keySlot),
		int32(valueSlot))
}

// GetVariable retreives a variable from the VM and stores it in the specified
// slot.
func (vm VM) GetVariable(module, name string, slot int) {
	ctx := vm.ctx()

	cmod := cstring(ctx, module)
	defer ctx.Free(cmod)

	cname := cstring(ctx, name)
	defer ctx.Free(cname)

	vm.ctx().WrenGetVariable(vm.ptr(), cmod, cname, int32(slot))
}

// HasVariable returns true if the value exists in this module.
func (vm VM) HasVariable(module, name string) bool {
	ctx := vm.ctx()

	cmod := cstring(ctx, module)
	defer ctx.Free(cmod)

	cname := cstring(ctx, name)
	defer ctx.Free(cname)

	return ctx.WrenHasVariable(vm.ptr(), cmod, cname) != 0
}

// HasModule returns true if the module has previously been imported or
// interpreted.
func (vm VM) HasModule(module string) bool {
	ctx := vm.ctx()

	cmod := cstring(ctx, module)
	defer ctx.Free(cmod)

	return ctx.WrenHasModule(vm.ptr(), cmod) != 0
}

// AbortFiber stops the current fiber using the value in slot as the error
// value.
func (vm VM) AbortFiber(slot int) {
	vm.ctx().WrenAbortFiber(vm.ptr(), int32(slot))
}

// AbortError is a helper function that takes a Go error and uses it to abort
// wren.
func (vm VM) AbortError(err error) {
	vm.EnsureSlots(1)
	vm.SetString(0, err.Error())
	vm.AbortFiber(0)
}

// UserData gets the user data of this VM.
func (vm VM) UserData() any {
	return vm.env.UserData
}

// SetUserdata sets the user data for this VM
func (vm VM) SetUserdata(data interface{}) {
	vm.env.UserData = data
}

// Handle is used to keep a reference to a value in the VM as to prevent it from
// being garbage collected and to refer to it later in foreign functions, but
// can also be used to call methods from the VM.
type Handle struct {
	vm  VM
	ptr int32
}

// Call calls this handle, using the receiver in slot 0 and arguments from slot
// 1 or greater.
//
// The handle should have been created using NewCallHandle.
//
// Note: that wren is not reentrant. You should not call this function from a
// foreign method or while the VM is currently running.
func (h *Handle) Call() InterpretResult {
	vm := h.vm
	return InterpretResult(vm.ctx().WrenCall(vm.ptr(), h.ptr))
}

// Free releases and frees up this handle. Handles prevent values from being
// garbage collected so releasing a handle removes this referance to that value,
// allowing it to be garbage collected when all references are gone.
func (h *Handle) Free() {
	vm := h.vm
	vm.ctx().WrenReleaseHandle(vm.ptr(), h.ptr)
}

// ErrorType represents what type of error occured in wren.
type ErrorType int32

const (
	// ErrorCompile represents an error during compilation.
	ErrorCompile ErrorType = iota
	// ErrorRuntime represunts an error during runtime.
	ErrorRuntime
	// ErrorStackTrace represents previous callstacks that lead to a runtime
	// error.
	ErrorStackTrace
)

// InterpretResult represents the result of running code in the VM.
type InterpretResult int32

const (
	// ResultSuccess represents no errors
	ResultSuccess InterpretResult = iota
	// ResultCompileError represents an error during the compilation of a module
	ResultCompileError
	// ResultRuntimeError represents an error during the running of a module.
	ResultRuntimeError
)

// Type represents the value type of a slot in wren.
type Type int32

const (
	TypeBool Type = iota
	TypeNum
	// TypeForeign represents a foreign object.
	TypeForeign
	TypeList
	TypeMap
	TypeNull
	TypeString
	// TypeUnknown represents a type that cannot currently be accessed by the
	// API.
	TypeUnknown
)

// New creates and initializes a new VM.
func NewVM(cfg Config) VM {
	e := new(env)
	e.objs = make(map[int32]interface{})
	e.fns = make(map[int32]func(VM))
	e.finalizers = make(map[int32]func(VM, interface{}))
	e.cfg = cfg
	e.ctx = internals.NewContext(e)
	e.ctx.MaxSize = 0x7FFFFFFF
	e.ctx.Start()
	return VM{e}
}
