package demo

import (
	"fmt"
	"testing"
)

func TestMapNil(t *testing.T) {
	// test nil[k] exists
	var m1 map[string]int

	// fatal: 不过slice 可以 append(nil, 1)
	// m1["kk"] = 1

	println("nil[k]=0:", m1["k"])
	if _, ok := m1["a"]; !ok {
		fmt.Println("nil[a] not exists. ")
	}

}
