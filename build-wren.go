//go:build ignore
// +build ignore

package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)


// These are used to build the go sources:
//     `go env GOPATH`/bin/cmd ir wren.wasm -o wren.ir
//     `go env GOPATH`/bin/cmd gen wren.ir -c -e=false -C -p internals -u -d internals

// TODO: Wagu doesn't like wasms from wasi-sdk-16, should use 15 for now. Also fetching and extracting wasi-sdk from this script is broken at the moment. Some executable files are do not get copied for some unknown reason. For now it is best to manually download wasi-sdk-15 and store it in ./wasi-sdk.

func main() {
	// Fetch wren
	if _, err := os.Stat("wren"); os.IsNotExist(err) {
		Must(Command("git", "clone", "-b", "made-for-embedding", "https://github.com/crazyinfin8/wren").Run())
	} else {
		Must(err)
		fmt.Println("Wren Exists already")
	}
	// Generate amalgamation
	if _, err := os.Stat("wren/wren.c"); os.IsNotExist(err) {
		f, err := os.Create("wren/wren.c")
		Must(err)
		defer f.Close()
		cmd := exec.Command("python", "wren/util/generate_amalgamation.py")
		cmd.Stdout = f
		Must(cmd.Run())
	} else {
		Must(err)
		fmt.Println("Amalgamation generated already")
	}
	// Fetch Wasi SDK
	if _, err := os.Stat("wasi-sdk"); os.IsNotExist(err) {
		resp, err := http.Get("https://github.com/WebAssembly/wasi-sdk/releases/latest")
		Must(err)
		var search *regexp.Regexp
		switch runtime.GOOS {
		case "windows":
			search = regexp.MustCompile(`\/WebAssembly\/wasi-sdk\/releases\/download\/wasi-sdk-\d+\/wasi-sdk-\d+\.0-linux\.tar\.gz`)
		case "linux":
			search = regexp.MustCompile(`\/WebAssembly\/wasi-sdk\/releases\/download\/wasi-sdk-\d+\/wasi-sdk-\d+\.0-mingw\.tar\.gz`)
		case "darwin":
			search = regexp.MustCompile(`\/WebAssembly\/wasi-sdk\/releases\/download\/wasi-sdk-\d+\/wasi-sdk-\d+\.0-macos\.tar\.gz`)
		default:
			panic("Wasi-sdk is not available to download for system: \"" + runtime.GOOS + "\". Please download into path ./wasi-sdk")
		}
		data, err := io.ReadAll(resp.Body)
		Must(err)
		WASI_SDK_URL := search.Find(data)
		if WASI_SDK_URL == nil {
			fmt.Println("Cannot find latest wasi-sdk download")
		}
		println("latest wasi-sdk is available at: github.com"+string(WASI_SDK_URL), "\nfetching...")
		resp, err = http.Get("https://github.com" + string(WASI_SDK_URL))
		Must(err)
		gzReader, err := gzip.NewReader(resp.Body)
		Must(err)
		tarReader := tar.NewReader(gzReader)
		header, eof := tarReader.Next()
		for eof == nil {
			fullPath := strings.Split(header.Name, "/")
			fullPath[0] = "wasi-sdk"
			currentPath := filepath.Join(fullPath...)
			if header.FileInfo().IsDir() {
				fmt.Println("Creating directory: " + currentPath)
				Must(os.Mkdir(currentPath, header.FileInfo().Mode()))
			} else {
				fmt.Println("Writing file: " + currentPath)
				f, err := os.Create(currentPath)
				Must(err)
				defer f.Close()
				_, err = io.Copy(f, tarReader)
				Must(err)
			}
			header, eof = tarReader.Next()
		}
	} else {
		Must(err)
		fmt.Println("Wasi-sdk exists already")
	}
	// Create wren.wasm
	Must(Command("./wasi-sdk/bin/clang", append([]string{"-Iwren/src/include", "-Iwren/src/optional", "-Iwren/src/vm", "-D_WASI_EMULATED_PROCESS_CLOCKS", "-lwasi-emulated-process-clocks", "-Wdeprecated-register", /*"-Os",*/ "wren/wren.c", "src/*.c", "-o", "wren.wasm"}, GetExports()...)...).Run())
	// Install wagu
	Must(Command("go", "install", "github.com/crazyinfin8/wagu/cmd@latest").Run())
	// Get path of Wagu
	cmd := Command("go", "env", "GOPATH")
	buf := strings.Builder{}
	cmd.Stdout = &buf
	Must(cmd.Run())
	GOPATH := buf.String()
	GOPATH = GOPATH[:len(GOPATH)-1]
	println("GOPATH:", GOPATH)
	WAGU_PATH := filepath.Join(GOPATH, "bin/cmd")
	println("WAGU_PATH:", WAGU_PATH)
	// Create wren.ir
	Must(Command(WAGU_PATH, "ir", "wren.wasm", "-o", "wren.ir").Run())
	// Generate go source from wren.ir
	Must(Command(WAGU_PATH, "gen", "wren.ir", "-c", "-C", "-p", "internals", "-d", "internals").Run())
}

func LookPath(file string) string {
	if cmd, err := exec.LookPath("git"); err == nil {
		return cmd
	}
	return file
}

func Command(path string, args ...string) *exec.Cmd {
	fmt.Println(path, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func GetExports() (exports []string) {
	regex := regexp.MustCompile(`WREN_API .+ (.+)\(`)
	str, err := ioutil.ReadFile("wren/src/include/wren.h")
	Must(err)
	matches := regex.FindAllSubmatch(str, -1)
	buf := strings.Builder{}
	for _, match := range matches {
		buf.Reset()
		buf.WriteString("-Wl,--export=")
		buf.Write(match[1])
		str := buf.String()
		exports = append(exports, str)
	}
	exports = append(exports,
		"-Wl,--export=free",
		"-Wl,--export=malloc",
		"-Wl,--export=calloc",
		"-Wl,--export=realloc",
		"-Wl,--export=stdin",
		"-Wl,--export=stdout",
		"-Wl,--export=stderr",
		"-Wl,--export=getVM",
	)
	return
}
