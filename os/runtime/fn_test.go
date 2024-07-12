package demo

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func _add(a, b int) int {
	return a + b

}
func TestFuncName(t *testing.T) {
	ptr := reflect.ValueOf(_add).Pointer()
	fnName := runtime.FuncForPC(ptr).Name()
	if fnName == "demo._add" {
		fmt.Println("yes")
	}
}
