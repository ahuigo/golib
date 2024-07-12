package main

import (
	"os"
	"path/filepath"
	"testing"
	//"path/filepath"
)


func TestCreateTmpFile(t *testing.T){
    // os.MkdirAll(path, os.ModePerm) //recursive
    dir := "tmp/a"
    os.MkdirAll(dir, 0700)
    path, err := os.MkdirTemp(dir, "pre_*_suffix")
    if err!=nil{
        panic(err)
    }

    file, err := os.CreateTemp(path, "tmp_*.go")
    defer os.Remove(file.Name())
    if err!=nil{
        panic(err)
    }
	// how to get file's absolute path?
    println("Succssfully create file: ", file.Name())
	absPath,_ := filepath.Abs(file.Name())
    println("abs path ", absPath)

}
