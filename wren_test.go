package wren_test

import (
	"errors"
	"testing"
	"time"

	"github.com/crazyinfin8/wren-go"
	"github.com/crazyinfin8/wren-go/internals"
)

type FnID struct {
	module    string
	class     string
	signature string
	static    bool
}

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

	time.Sleep(3 * time.Microsecond)
	vm.Exit()
}

func TestOOM(t *testing.T) {
	defer func() {
		v := recover()
		err, ok := v.(error)

		if !ok {
			println(v)
			t.Fatal("Expected Error")
			t.FailNow()
		}
		if !errors.Is(err, internals.ErrOutOfMemory) {
			t.Fatal("Expected out of memory error")
		}
		t.Log("Received expected error:", err.Error())
	}()

	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	vm := wren.NewVM(cfg)

	vm.Interpret("main", `
	var list = []
	var start = System.clock
	for (i in 0...1000000) list.add(i)

	var sum = 0
	for (i in list) sum = sum + i

	System.print(sum)
	System.print("elapsed: %(System.clock - start)")
	`)
}

func TestClock(t *testing.T) {
	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	vm := wren.NewVM(cfg)

	tick := time.NewTicker(time.Second)
	i := 0
	for i < 10 {
		<-tick.C
		vm.Interpret("main", `System.write("%(System.clock)\r")`)
		i++
	}
	tick.Stop()
}

func TestRand(t *testing.T) {
	cfg := wren.Config{}

	cfg.WriteFn = func(vm wren.VM, message string) { t.Log(message) }

	cfg.ErrorFn = func(vm wren.VM, err error) { t.Error(err.Error()) }

	fns := map[FnID]func(wren.VM){
		{"main", "Time", "wait(_)", true}: func(vm wren.VM) {
			time.Sleep(time.Duration(vm.GetNum(1)) * time.Second)
		},
	}

	cfg.BindForeignMethodFn = func(vm wren.VM, module, class, signature string,
		static bool) func(wren.VM) {
		return fns[FnID{module, class, signature, static}]
	}
	vm := wren.NewVM(cfg)
	vm.Interpret("main", `
	import "random" for Random

	foreign class Time {
		// sometimes it seems as if time doesn't increase when running.
		// so we will force it to wait some bit
		foreign static wait(t)
	}
	var rand

	var deepEqual = Fn.new{|a, b|
		if(a.count != b.count) return false
		for (i in 1...a.count) if(a[i] != b[i]) return false
		return true
	}

	/* --- random with different seeds --- */

	Time.wait(1)
	rand = Random.new()
	var a = []
	for(i in 0...100) a.add(rand.int())

	Time.wait(1)
	rand = Random.new()
	var b = []
	for(i in 0...100) b.add(rand.int())

	if (deepEqual.call(a, b)) Fiber.abort("Random is not random enough (or we are just really unlucky)")

	/* --- random with same seed --- */

	Time.wait(1)
	rand = Random.new(0)
	a = []
	for(i in 0...100) a.add(rand.int())
	
	Time.wait(1)
	rand = Random.new(0)
	b = []
	for(i in 0...100) b.add(rand.int())

	if (!deepEqual.call(a, b)) Fiber.abort("Random should have been equal but is not")
	`)
}
