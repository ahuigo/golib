package demo

import (
	"fmt"
	"testing"
)

func TestMapNil(t *testing.T) {
	var m map[string]string
	m["a"] = "1"
	fmt.Println(m) //bad
}
