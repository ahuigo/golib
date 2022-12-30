package demo

import (
	"fmt"
	"testing"
)

func TestInterfaceNil(t *testing.T) {
	var i interface{} = nil
	fmt.Printf("obj=%v,%T\n", i, i)
	fmt.Printf("obj=%v,%T\n", i == nil, i == nil)
}
