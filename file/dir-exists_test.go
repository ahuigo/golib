package main

import (
	"fmt"
	"os"
    "testing"
)

func TestFileExists(t *testing.T) {
    if _, err := os.Stat("./dir-exists_test.go"); !os.IsNotExist(err) {
        fmt.Println("file exists!")
    }else{
        t.Fatal("file not fould")
    }
}



