package wren

// basicConfig implements Config on top of BasicConfig. This is a seperate
// unexported type since it automatically handles interactions between the VM
// and BasicConfig.
type basicConfig[ObjType any] BasicConfig[ObjType]

func (cfg *basicConfig[_]) ResolveModuleFn(vm VM, importer,
	name string) (string, bool) {
	if cfg.Fns == nil {
		return "", false
	}
	return cfg.Fns.ResolveModuleFn(vm, importer, name)
}

func (cfg *basicConfig[_]) LoadModuleFn(vm VM, name string) (string, bool) {
	if cfg.Fns == nil {
		return "", false
	}
	return cfg.Fns.LoadModuleFn(vm, name)
}

func (cfg *basicConfig[_]) bindForeignFn(fnIndex FnIndex) int32 {
	fn, ok := cfg.foreignFns[fnIndex]
	if ok {
		delete(cfg.foreignFns, fnIndex)
		cfg.foreignSeed = nextEmptyIndex(cfg.importedForeigns, cfg.foreignSeed)
		cfg.importedForeigns[cfg.foreignSeed] = fn
		return cfg.foreignSeed
	}
	return 0
}

func (cfg *basicConfig[_]) BindForeignMethodFn(vm VM, module, className string,
	isStatic bool, signature string) int32 {
	return cfg.bindForeignFn(FnIndex{module, className, isStatic, signature})
}

func (cfg *basicConfig[_]) BindForeignClassMethodFn(vm VM, module,
	className string) (int32, int32) {
	fnIndex := FnIndex{Module: module, ClassName: className}
	allocator := cfg.bindForeignFn(fnIndex)
	// Only search finalizers if a constructor exists.
	if allocator == 0 {
		return 0, 0
	}
	var finalizer int32
	fn, ok := cfg.finalizerFns[fnIndex]
	if ok {
		delete(cfg.finalizerFns, fnIndex)
		cfg.finalizerSeed = nextEmptyIndex(cfg.importedFinalizers,
			cfg.finalizerSeed)
		cfg.importedFinalizers[cfg.finalizerSeed] = fn
	}
	return allocator, finalizer
}

func (cfg *basicConfig[_]) WriteFn(vm VM, message string) {
	if cfg.Fns != nil {
		cfg.Fns.WriteFn(vm, message)
	}
}

func (cfg *basicConfig[_]) ErrorFn(vm VM, err error) {
	if cfg.Fns != nil {
		cfg.Fns.ErrorFn(vm, err)
	}
}

func (cfg *basicConfig[_]) DispatchForeignFn(vm VM, ID int32) {
	cfg.importedForeigns[ID](vm)
}

// DispatchFinalizerFn implements Config for BasicConfig. It automatically
// removes the foreign object from its list of foreigns so it can be garbage
// collected by Go.
func (cfg *basicConfig[ObjType]) DispatchFinalizerFn(vm VM, dataID int32,
	fnID int32) {
	// Remove item so it can be garbage collected by Go.
	delete(cfg.objects, dataID)
	cfg.importedFinalizers[fnID](vm, cfg.objects[dataID])
}

type FnIndex struct {
	Module, ClassName string
	Static            bool
	Signature         string
}

// BasicConfig makes Config much more ergonomic to use since Config doesn't
// allow the host to directly pass functions and foreign objects to wren.
// Instead, this stores foreign methods, finalizers, and objects so it can
// easily pass index keys to wren instead.
//
// ObjType specifies the type of foreign objects passed to wren. This might help
// accessing foreign objects without having to cast.
type BasicConfig[ObjType any] struct {
	env *env

	foreignFns   map[FnIndex]func(VM)
	finalizerFns map[FnIndex]func(VM, ObjType)

	foreignSeed        int32
	importedForeigns   map[int32]func(VM)
	finalizerSeed      int32
	importedFinalizers map[int32]func(VM, ObjType)

	objectSeed int32
	objects    map[int32]ObjType

	Fns BasicConfigFns
}

// BasicConfigFns are a set of functions to optionally add to BasicConfig.
type BasicConfigFns interface {
	ResolveModuleFn(vm VM, importer, name string) (string, bool)
	LoadModuleFn(vm VM, name string) (string, bool)
	WriteFn(vm VM, message string)
	ErrorFn(vm VM, err error)
}

// NewBasicConfig creates a new BasicConfig and initializes a new VM.
func NewBasicConfig[ObjType any](fns BasicConfigFns) *BasicConfig[ObjType] {
	cfg := new(BasicConfig[ObjType])

	cfg.Fns = fns

	cfg.foreignFns = make(map[FnIndex]func(VM))
	cfg.finalizerFns = make(map[FnIndex]func(VM, ObjType))

	cfg.importedForeigns = make(map[int32]func(VM))
	cfg.importedFinalizers = make(map[int32]func(VM, ObjType))
	cfg.objects = make(map[int32]ObjType)

	cfg.env = New((*basicConfig[ObjType])(cfg)).env

	return cfg
}

// VM returns the current VM tied to this BasicConfig.
func (cfg *BasicConfig[_]) VM() VM {
	return VM{cfg.env}
}

// nextEmptyIndex attempts to find the next unpopulated index in m. This
// function uses a linear congruential generator to create pseudorandom indexes
// though this isn't entirely necessary. The values for the LCG were picked to
// ensure that the LCGs period is equal to the size of int32.
//
// See: https://en.wikipedia.org/wiki/Linear_congruential_generator#c_%E2%89%A0_0
func nextEmptyIndex[T any](m map[int32]T, seed int32) int32 {
	zero := false
	for {
		seed = int32((1664525*int64(seed) + 0x100000000) % 0xFFFFFFFF)
		// Do not use seed value of 0.
		if seed == 0 {
			// Unlikely that this loop will hit every value before running out
			// of memory but if zero is hit twice, we have exausted all possible
			// keys available.
			if zero {
				panic("No keys available")
			}
			zero = true
			continue
		}
		if _, ok := m[seed]; !ok {
			return seed
		}
	}
}

// SetForeignFn sets a foreign method for this config. This method can be
// changed up until the VM actually imports this module. Setting fn to nil will
// delete the target method if it exists.
func (cfg *BasicConfig[_]) SetForeignFn(module, className string, isStatic bool,
	signature string, fn func(VM)) {
	index := FnIndex{module, className, isStatic, signature}
	if fn == nil {
		delete(cfg.foreignFns, index)
	} else {
		cfg.foreignFns[index] = fn
	}
}

// SetClassMethods sets a foreign class's allocator and finalizer. These can be
// changed until the VM actually imports this module. If finalizer is nil, it
// will delete the target finalizer if it exists. if allocator is nil, it will
// delete both of the target allocator and finalizer because an allocator is
// needed in order to call the finalizer.
func (cfg *BasicConfig[ObjType]) SetClassMethods(module, className string,
	allocator func(VM), finalizer func(VM, ObjType)) {
	index := FnIndex{Module: module, ClassName: className}
	if allocator == nil {
		// Class must have an allocator in order to have a finalizer.
		delete(cfg.foreignFns, index)
		delete(cfg.finalizerFns, index)
		return
	}
	cfg.foreignFns[index] = allocator
	if finalizer == nil {
		delete(cfg.finalizerFns, index)
	} else {
		cfg.finalizerFns[index] = finalizer
	}
}

// AddNewObject adds a new foreign object and returns an ID to be passed to
// SetSlotNewForeign.
func (cfg *BasicConfig[ObjType]) AddObject(obj ObjType) int32 {
	cfg.objectSeed = nextEmptyIndex(cfg.objects, cfg.objectSeed)
	cfg.objects[cfg.objectSeed] = obj
	return cfg.objectSeed
}

// GetObject retreives a foreign object with the given ID. The value ok is true
// if an element exists.
func (cfg *BasicConfig[ObjType]) GetObject(ID int32) (val ObjType, ok bool) {
	val, ok = cfg.objects[ID]
	return
}

// SetObject sets the value of an existing foreign object with the given ID. If
// the ID doesn't exist, this object doesn't do anything and returns false. To
// add new foreign objects, call AddObject instead.
func (cfg *BasicConfig[ObjType]) SetObject(obj ObjType, ID int32) (ok bool) {
	_, ok = cfg.objects[ID]
	if ok {
		cfg.objects[ID] = obj
	}
	return
}
