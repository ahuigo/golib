package main

import (
	"fmt"
)

func main() {
	//src:=[]byte{'a',0,'b'}
	src := []byte("abc")
	p := make([]byte, 10)
	p[4] = 'c'
	copy(p, src)
	fmt.Println((p), len(p))
}
