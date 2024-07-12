package demo

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCurrentDir(t *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}

func isInTest() bool {
	// strings.HasSuffix(os.Args[0], ".test")
	//Or strings.Contains(os.Args[0], "/_test/")
	return flag.Lookup("test.v") != nil
}

func TestCaller(t *testing.T) {
	if !isInTest() {
		return
	}
	fmt.Println("in test mode")
	// 仅在编译时生效: pc uintptr 是调用栈的在程序中的pc counter(cpu)地址
	pc, filename, lineno, _ := runtime.Caller(0)
	fmt.Printf("funcname=%s, file=%s:%d\n", runtime.FuncForPC(pc).Name(), filename, lineno)
	dir := path.Join(path.Dir(filename), ".")
	fmt.Println(dir)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestCallers(t *testing.T) {
	var pcs = make([]uintptr, 64)
	n := runtime.Callers(0, pcs)
	for _, pc := range pcs[:n] {
		fmt.Println(pc, runtime.FuncForPC(pc).Name())
	}
}
func TestCallersFrames(t *testing.T) {
	var pcs = make([]uintptr, 64)
	var count = 0
	n := runtime.Callers(0, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	for {
		count++
		frame, more := frames.Next()
		fmt.Println("funcName", frame.Function, frame.File)
		if !more || count > 100 {
			break
		}
	}
}
