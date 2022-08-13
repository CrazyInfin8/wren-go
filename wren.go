package wren

//go:generate go run build-wren.go

import (
	"io"
	"strconv"
	"strings"

	"github.com/crazyinfin8/wren-go/internals"
)

type Config interface {
	// ResolveModuleFn allows the host to canonicalize the imported name beofre
	// it is passed to LoadModuleFn.
	//
	// It takes the name of the module doing the import and the target module to
	// import. It should return the new name to be passed to LoadModuleFn as
	// well as if the import is ok. If an import cannot be resolved, this
	// function should return false.
	ResolveModuleFn(vm VM, importer, name string) (newName string, ok bool)

	// LoadModuleFn load a modules source code. It takes the name of the module
	// to import and returns the wren source code as well as whether the module
	// was successfully loaded.
	LoadModuleFn(vm VM, name string) (string, bool)

	// BindForeignMethodFn binds a foreign function to a class method in wren.
	//
	// It is passed the module, class name, and call signature of the function
	// as well as whether the function is static.
	//
	// It should return an int32 key which will be passed to DispatchForeignFn
	// when this foreign method is to be called. Returning 0 signifies an error
	// to wren when binding this function.
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
	BindForeignMethodFn(vm VM, module, className string, isStatic bool,
		signature string) int32

	// BindForeignClassMethodFn binds the allocator and finalizer of a class.
	//
	// It takes the name of the class and it's module and should return an int32
	// each for the allocator and finalizer. Setting the allocator to 0
	// signifies an error to wren but finalizer is optional and can be set to 0
	// to indicate no finalizer function.
	//
	// The value of allocator will be passed to DispatchForeignFn when the
	// constructor of the foreign class is called. It should make a call to
	// SetSlotNewForeign once. The value of finalizer is passed to
	// DispatchFinalizerFn as well as the value passed to SetSlotNewForeign when
	// the foreign object is garbage collected by wren.
	//
	// The allocator is called when the construtor of the foreign class is
	// called and should call SetSlotNewForeign once. The finalizer is called
	// when the foreign object is garbage collected from wren.
	BindForeignClassMethodFn(vm VM, module, className string) (allocator int32,
		finalizer int32)

	// WriteFn is called when "System.write", "System.print", and related
	// functions have been called in wren.
	WriteFn(vm VM, message string)

	// ErrorFn is called whenever acompilation or runtime error has occurred in
	// wren. The error value should always be of type WrenError.
	ErrorFn(vm VM, err error)

	// DispatchForeignFn is called by wren when a foreign function is
	// dispatched. It takes the value returned by BindForeignMethodFn or the
	// allocator of BindForeignClassMethodFn.
	DispatchForeignFn(vm VM, fnID int32)

	// DispatchFinalizerFn is called by wren when it has garbage collected a
	// foreign object. The value of fnID will be the value of finalizer that was
	// returned from BindForeignClassMethodFn. The value of data is the value
	// passed by SetSlotNewForeign.
	DispatchFinalizerFn(vm VM, data int32, fnID int32)
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
		var buf strings.Builder
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

// CollectGarbage runs the garbage collector for wren.
func (vm VM) CollectGarbage() {
	vm.ctx().WrenCollectGarbage(vm.env.vmPtr())
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
func (vm VM) CallHandle(handle Handle) InterpretResult {
	ctx := vm.ctx()
	if ctx != handle.vm.ctx() {
		panic("handle's context does not match VM")
	}
	return InterpretResult(ctx.WrenCall(vm.ptr(), handle.ptr))
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

// GetSlotType gets the value type located in the slot provided.
func (vm VM) GetSlotType(slot int) Type {
	return Type(vm.ctx().WrenGetSlotType(vm.ptr(), int32(slot)))
}

// GetSlotBool retreives a bool in the slot provided.
func (vm VM) GetSlotBool(slot int) bool {
	return vm.ctx().WrenGetSlotBool(vm.ptr(), int32(slot)) != 0
}

// GetSlotBytes retrieves string data from wren and writes it into writer. Error
// values from writer are returned.
func (vm VM) GetSlotBytes(slot int, writer io.Writer) error {
	ctx := vm.ctx()

	sizePtr := ctx.Malloc(4)
	defer ctx.Free(sizePtr)

	dataPtr := ctx.WrenGetSlotBytes(vm.ptr(), int32(slot), sizePtr)
	endPtr := dataPtr + readInt32(ctx, sizePtr)

	for dataPtr < endPtr {
		n, err := writer.Write(ctx.Mem[dataPtr : endPtr-dataPtr])
		if err != nil {
			return err
		}
		dataPtr += int32(n)
	}
	return nil
}

// GetSlotDouble returns a number in the slot provided.
func (vm VM) GetSlotDouble(slot int) float64 {
	return vm.ctx().WrenGetSlotDouble(vm.ptr(), int32(slot))
}

// GetSlotForeign returns and ID for the current foreign object. It's value
// reflects what was passed from SetSlotNewForeign.
func (vm VM) GetSlotForeign(slot int) int32 {
	return vm.ctx().WrenGetSlotForeign(vm.ptr(), int32(slot))
}

// GetSlotString copies a string from the slot provided and returns it.
func (vm VM) GetSlotString(slot int) string {
	buf := strings.Builder{}
	vm.GetSlotBytes(slot, &buf)

	return buf.String()
}

// GetSlotHandle creates a new handle from the value in the slot provided.
// Handles prevent a value from being garbage collected.
func (vm VM) GetSlotHandle(slot int) Handle {
	return Handle{vm, vm.ctx().WrenGetSlotHandle(vm.ptr(), int32(slot))}
}

// SetSlotBool sets the provided slot to a boolean value.
func (vm VM) SetSlotBool(slot int, value bool) {
	vm.ctx().WrenSetSlotBool(vm.ptr(), int32(slot), boolToInt(value))
}

// SetSlotBytes sets the provided slot to a string value from the given bytes.
func (vm VM) SetSlotBytes(slot int, data []byte) {
	ctx := vm.ctx()

	cbuf := cbytes(ctx, data)
	defer ctx.Free(cbuf)

	ctx.WrenSetSlotBytes(vm.ptr(), int32(slot), cbuf, int32(len(data)))
}

// SetSlotDouble sets the provided slot to a number value from the given
// float64.
func (vm VM) SetSlotDouble(slot int, value float64) {
	vm.ctx().WrenSetSlotDouble(vm.ptr(), int32(slot), value)
}

// SetSlotnewForeign sets the slot to the value of a new foreign object. The
// objects type should be of the class type found in classSlot. The ID will be
// the same id provided from GetSlotForeign and DispatchFinalizerFn.
func (vm VM) SetSlotNewForeign(slot, classSlot int, ID int32) {
	ctx := vm.ctx()
	ctx.WrenSetSlotNewForeign(vm.ptr(), int32(slot), int32(classSlot), ID)
}

// SetSlotNewList creates an empty list and sets it in the slot provided.
func (vm VM) SetSlotNewList(slot int) {
	vm.ctx().WrenSetSlotNewList(vm.ptr(), int32(slot))
}

// SetSlotNewMap creates an empty map and sets it in the slot provided.
func (vm VM) SetSlotNewMap(slot int) {
	vm.ctx().WrenSetSlotNewMap(vm.ptr(), int32(slot))
}

// SetSlotNull sets the provided slot to a null value.
func (vm VM) SetSlotNull(slot int) {
	vm.ctx().WrenSetSlotNull(vm.ptr(), int32(slot))
}

// SetSlotString sets the provided slot to the given string value.
func (vm VM) SetSlotString(slot int, value string) {
	ctx := vm.ctx()

	cstr := cstring(ctx, value)
	defer ctx.Free(cstr)

	ctx.WrenSetSlotBytes(vm.ptr(), int32(slot), cstr, int32(len(value)))
}

// SetSlotHandle sets the provided slot to the value held by the given handle.
func (vm VM) SetSlotHandle(slot int, handle Handle) {
	ctx := vm.ctx()
	if ctx != handle.vm.ctx() {
		panic("handle's context does not match VM")
	}
	ctx.WrenSetSlotHandle(vm.ptr(), int32(slot), handle.ptr)
}

// GetListCount returns the list of the list in the given slot.
func (vm VM) GetListCount(slot int) int {
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

// GetMapCount returns the amount of elements in the map at the specified slot.
func (vm VM) GetMapCount(slot int) int {
	return int(vm.ctx().WrenGetMapCount(vm.ptr(), int32(slot)))
}

// GetMapContainsKey returns true if the value in keySlot is a valid key in this
// map
func (vm VM) GetMapContainsKey(mapSlot, keySlot int) bool {
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
	vm.SetSlotString(0, err.Error())
	vm.AbortFiber(0)
}

// GetUserData gets the userdata of this VM.
func (vm VM) GetUserData() int32 {
	return vm.ctx().WrenGetUserData(vm.ptr())
}

// SetUserData sets the userdata of this VM
func (vm VM) SetUserData(data int32) {
	vm.ctx().WrenSetUserData(vm.ptr(), data)
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
func New(cfg Config) VM {
	e := new(env)
	e.cfg = cfg
	e.ctx = internals.NewContext(e)
	e.ctx.Start()
	return VM{e}
}
