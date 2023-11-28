package main

import (
	"fmt"
	"testing"
)

func TestArgsPointer(t *testing.T) {
	f := func(ids *[]byte) {
		*ids = (*ids)[1:]
	}
	ids := []byte{1, 2, 3}
	fmt.Printf("ids:%v, len(ids)=%d\n", ids, len(ids))
	f(&ids)
	fmt.Printf("ids:%v, len(ids)=%d\n", ids, len(ids))

}

func TestArgsSlice(t *testing.T) {
	f := func(ids []byte) {
		ids[0] = 100
		ids = []byte{2, 3} //unsed
		_ = ids
	}
	ids := []byte{1}
	f(ids)
	fmt.Printf("ids:%v, len(ids)=%d\n", ids, len(ids)) //ids:[100], len(ids)=1

}
