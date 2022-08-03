package main

import (
	"fmt"
	"testing"
)

func TestTypeAlias(t *testing.T) {
	type fileHandler struct {
		root string
	}
	buildHandler := func() interface{} {
		return &fileHandler{"/root"}
	}
	type FileHandler = fileHandler

	fh := buildHandler()
	_ = fh
	_ = []byte("abc")
	_ = string("abc")
	fh1 := fh.(*fileHandler)
	fh2 := fh.(*FileHandler)
	fmt.Printf("fh1:%v, fh2:%v\n", fh1, fh2)

}
