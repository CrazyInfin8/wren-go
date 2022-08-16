package wren_test

import (
	"testing"
	"time"

	"github.com/crazyinfin8/wren-go"
)

func Test(t *testing.T) {
	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	vm := wren.NewVM(cfg)

	res := vm.Interpret("main", `System.print("Hello world")`)

	if res != wren.ResultSuccess {
		t.Error("Unexpected error")
	}
}

func TestForeignFunctions(t *testing.T) {
	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	cfg.ResolveModuleFn = func(vm wren.VM, importer,
		name string) (newName string, ok bool) {
		if importer == "main" && name == "a.wren" {
			return "a", true
		}
		return
	}

	cfg.LoadModuleFn = func(vm wren.VM, name string) (src string, ok bool) {
		if name == "a" {
			return `
			foreign class A {
				construct new() {}
				foreign static sayHello()
				foreign add(a, b)
				foreign static a()
			}
			`, true
		}
		return
	}

	type FnID struct {
		module    string
		class     string
		signature string
		static    bool
	}
	fns := map[FnID]func(wren.VM){
		{"a", "A", "sayHello()", true}: func(v wren.VM) {
			t.Log("Hello from foreign function")
			v.EnsureSlots(1)
			v.SetNull(0)
		},
		{"a", "A", "add(_,_)", false}: func(v wren.VM) {
			if v.SlotCount() != 3 {
				t.Error("Wrong number of parameters")
			}
			a, b := v.GetNum(1), v.GetNum(2)
			v.EnsureSlots(1)
			v.SetNum(0, a+b)
		},
		{"a", "A", "a()", true}: func(v wren.VM) {
			v.EnsureSlots(1)
			v.SetNewForeign(0, 0, "instance of class A from a method")
		},
	}
	cfg.BindForeignMethodFn = func(vm wren.VM, module, class, signature string,
		static bool) func(wren.VM) {
		return fns[FnID{module, class, signature, static}]
	}

	classes := map[FnID]struct {
		allocator func(wren.VM)
		finalizer func(wren.VM, interface{})
	}{
		{module: "a", class: "A"}: {
			func(v wren.VM) {
				t.Log("constructor called")
				v.EnsureSlots(1)
				v.SetNewForeign(0, 0, "instance of class A")
			},
			func(v wren.VM, a interface{}) { t.Log("finalizer called with:", a) },
		},
	}

	cfg.BindForeignClassMethodFn = func(vm wren.VM, module,
		class string) (allocator func(wren.VM), finalizer func(wren.VM, interface{})) {
		fns := classes[FnID{module: module, class: class}]
		return fns.allocator, fns.finalizer
	}

	vm := wren.NewVM(cfg)

	res := vm.Interpret("main", `
	import "a.wren" for A
	A.sayHello()
	System.print("A: %(A); type: %(A.type)")

	var a = A.new()
	System.print("5 + 7 = %(a.add(5, 7))")

	System.print("A.a: %(A.a())")
	`)

	vm.CollectGarbage()

	t.Log("freeing VM")

	vm.Free()

	if res != wren.ResultSuccess {
		t.Error("Unexpected error")
	}
}

func TestVersion(t *testing.T) {
	vm := wren.NewVM(wren.Config{})

	t.Log(vm.VersionTuple())
	t.Log(vm.VersionString())
}

func TestExit(t *testing.T) {
	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	vm := wren.NewVM(cfg)

	go func() {
		result := vm.Interpret("main", `
		var i = 0
		while (true) {
			i = i+1
			System.write("The loop has ran %(i) times")
		}`)
		if result == wren.ResultRuntimeError {
			t.Log("Got expected runtime error")
		} else {
			t.Error("Incorect result:", result)
		}
	}()

	time.Sleep(15 * time.Microsecond)
	vm.Exit()
}
