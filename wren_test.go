package wren_test

import (
	"testing"

	"github.com/crazyinfin8/wren-go"
)

type Fns struct{ *testing.T }

func (t *Fns) ResolveModuleFn(vm wren.VM, importer,
	name string) (string, bool) {
	return "", false
}

func (t *Fns) LoadModuleFn(vm wren.VM, name string) (string, bool) {
	return "", false
}

func (t *Fns) WriteFn(vm wren.VM, message string) { t.Log(message) }

func (t *Fns) ErrorFn(vm wren.VM, err error) { t.Error(err.Error()) }

func Test(t *testing.T) {
	fns := new(Fns)
	fns.T = t

	cfg := wren.NewBasicConfig[any](fns)
	vm := cfg.VM()

	res := vm.Interpret("main", `System.print("Hello world")`)

	if res != wren.ResultSuccess {
		t.Error("Unexpected error")
	}
}

func TestExportEnvironment(t *testing.T) {
}
