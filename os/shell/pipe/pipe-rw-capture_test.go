package demo

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPipeStdout(t *testing.T) {
	fmt.Println("stdout==stderr", os.Stderr==os.Stdout)
	getOutput, restore := captureStdout()
	getErr := captureStderr()

	// wrtie to: system err(not use os.Stderr)
	println("some err1")
	// wrtie to: os.Stderr
	fmt.Fprintln(os.Stderr, "some err2")
	// wrtie to: os.Stdout
	fmt.Println("some stdout")

	out := getOutput()
	err := getErr()
	defer func() {
		restore()
		fmt.Printf("out:%d:%s\n", len(out), out)
		fmt.Printf("err:%d:%s\n", len(err), err)
	}()
}

func captureStderr() (getOutput func() string) {
	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = w
	getOutput = func() string {
		os.Stderr = old
		w.Close()
		buf := make([]byte, 2048)
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		return string(buf[:n])
	}
	return getOutput
}

func captureStdout() (getOutput func() string, restore func()) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	getOutput = func() string {
		w.Close()
		buf := make([]byte, 2048)
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		return string(buf[:n])
	}
	restore = func() {
		os.Stdout = old
		w.Close()
	}
	return getOutput, restore
}
