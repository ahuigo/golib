package demo

import (
	"fmt"
	"testing"
)

func TestInterfaceAssert(t *testing.T) {
	var i interface{} = "string"

	// method 1: switch
	switch v := i.(type) {
	case string:
		bytes := []byte(v)
		fmt.Printf("bytes: %v, type of v: %T\n", bytes, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	// method 2
	s := i.(string)
	fmt.Println("str:", s)

}
