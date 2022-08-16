# wren-go

## Wren scripting language in pure Go
[![Go Reference](https://pkg.go.dev/badge/github.com/crazyinfin8/wren-go.svg)](https://pkg.go.dev/github.com/crazyinfin8/wren-go) [![GoReportCard](https://goreportcard.com/badge/github.com/crazyinfin8/wren-go)](https://goreportcard.com/report/github.com/crazyinfin8/wren-go) [![Wren](https://img.shields.io/badge/github-wren-hsl(200%2C%2060%25%2C%2050%25))](https://github.com/wren-lang/wren)

wren-go is a port of the [Wren] scripting language to pure Go using WebAssembly and [wagu]. Thanks to [wagu], wren-go needs almost no dependencies and does not use CGo, making it easier to target web and tinygo targets.

Wren-go also uses a modified version of wren similar to [wren-bindings-for-go] meant for embedding in non C projects. This means theres no need to [preallocate foreign functions](https://github.com/CrazyInfin8/WrenGo/blob/main/bindings.go) and the `VM.Exit` function is meant to stop the wren VM in the middle of execution.

[wren]:[https://wren.io/]
[wagu]:[https://github.com/chrsan/wagu]
[WrenGo]:[https://github.com/CrazyInfin8/WrenGo]
[wren-bindings-for-go]:[https://github.com/CrazyInfin8/wren-bindings-for-go]

## Installation

```bash
go get github.com/crazyinfin8/wren-go
```

## Usage

```go
func main() {
	cfg := wren.Config{
		WriteFn: func(vm wren.VM, message string) { fmt.Print(message) },
	}

	vm := wren.NewVM(cfg)

	vm.Interpret("main", `System.print("Hello from wren!")`)
}
```

## Building wren-go

To build, run:

``` bash
cd path/to/github.com/crazyinfin8/wren-go
go generate
```

this runs `build-wren.go` which:

1. first fetches wren (requires git)
2. generates the wren amalgamation file
3. downloads [WASI-libc](https://github.com/WebAssembly/wasi-libc) (it has been buggy to download and run wasi-libc in one go so you may need to run this script twice)
4. compiles the wren amalgamation with `src/shim.c` to a webassembly binary using wasi-libc
5. installs wagu
6. generates the IR for wagu
7. finally generates Go source code into the `internals` folder.