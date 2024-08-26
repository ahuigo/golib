package demo

import (
	"fmt"
	"testing"
)

func TestMapNil(t *testing.T) {
	// test nil[k] exists
	var m1 map[string]int
	println("get mapnil(key):", m1["key113"])

	// fatal: 不过slice 可以 append(nil, 1)
	// m1["kk"] = 1 // 写不行,读可以
    
    // 1. read: range
	for _, v := range m1 {
		fmt.Println(v)
	}

    // 2. read: key
	println("nil[k]=0:", m1["k"])
	if _, ok := m1["a"]; !ok {
		fmt.Println("nil[a] not exists. ")
	}
    // 3. all is nil
    fmt.Printf("nil?=%v\n", m1==nil)
    fmt.Printf("nil?=%v\n", map[string]string(nil)==nil)
    fmt.Printf("nil?=%v\n", []byte(nil)==nil)
    // 3.2 this is not nil
    fmt.Printf("[]uint8{} !=nil, %v\n", []uint8{}==nil)

}
