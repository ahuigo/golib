package main

import (
	"fmt"
	"os"
)

func main() {
    // content, err := ioutil.ReadFile(filename)
    f, err := os.Open("file-read.go")

    if err!=nil{
        fmt.Println(err)
        return
    }
    defer f.Close()

    buf := make([]byte, 50)

    //
    _, err = f.Read(buf)
    fmt.Println(string(buf))
}

