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

func TestCaller(t *testing.T) {
	if !isInTest() {
		return
	}
	fmt.Println("in test mode")
	// 仅在编译时生效
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), ".")
	fmt.Println(dir)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func isInTest() bool {
	// strings.HasSuffix(os.Args[0], ".test")
	//Or strings.Contains(os.Args[0], "/_test/")
	return flag.Lookup("test.v") != nil
}
