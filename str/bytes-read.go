package main

import (
	"fmt"
    "bytes"
)

func main() {
    // read to bytes via buffer
    b:=make([]byte, 1, 10)
    buf := bytes.NewBuffer([]byte("hi"))
    buf.Read(b)
    fmt.Println("buf:",string(b))

    // read to bytes from bytes
    // bp:  0123llo   [0 48 49 50 51 108 108 111 0 0]
    p := make([]byte, 10)
    copy(p[3:], []byte("hello"))
    copy(p[1:], []byte("0123"))
    fmt.Println("bp:",string(p), p)
}

