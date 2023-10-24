package demo

import (
	"testing"
)

func TestMapNil(t *testing.T) {
	var m map[string]string //nil
	m["a"] = "1"            // error
}
