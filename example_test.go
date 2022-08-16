package wren_test

import (
	"errors"
	"fmt"

	"github.com/crazyinfin8/wren-go"
)

func ExampleVM() {
	cfg := wren.Config{
		WriteFn: func(vm wren.VM, message string) { fmt.Print(message) },
	}

	vm := wren.NewVM(cfg)

	vm.Interpret("main", `System.print("Hello from wren!")`)
	// Output:
	// Hello from wren!
}

func ExampleConfig_foreign() {
	type ID struct {
		module, class, signature string
		bool
	}
	cfg := wren.Config{
		WriteFn: func(vm wren.VM, message string) { println(message) },
		ErrorFn: func(vm wren.VM, err error) { println(err.Error()) },
		LoadModuleFn: func(vm wren.VM, name string) (src string, ok bool) {
			src, ok = map[ID]string{
				{module: "foreign"}: `
				foreign class Foreign {
					foreign static greet(name)
					construct new() {
						System.print("Hello world")
					}
				}
				`,
			}[ID{module: name}]
			return
		},
		BindForeignMethodFn: func(vm wren.VM, module, class, signature string, static bool) func(wren.VM) {
			return map[ID]func(wren.VM){
				{"foreign", "Foreign", "greet(_)", true}: func(v wren.VM) {
					if v.SlotCount() != 2 { // slot 0 is the "Foreign" class
						v.AbortError(errors.New("Expected 1 parameter"))
					}
					if v.SlotType(1) != wren.TypeString {
						v.AbortError(errors.New("Expected parameter 1 to be string"))
					}
					println("Hello", v.GetString(1))
				},
			}[ID{module, class, signature, static}]
		},
	}

	vm := wren.NewVM(cfg)

	vm.Interpret("main", `
	import "foreign" for Foreign

	Foreign.greet("wren")
	`)
}

func init() {
	// ExampleConfig_foreign()
}
