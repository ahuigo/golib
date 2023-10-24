package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func slice2interfce[T any](a []T) []interface{} {
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	return b
}

func interface2Slice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func TestSlice2Interface(t *testing.T) {
	fmt.Println("slice2interfce:", slice2interfce([]int{1, 2, 3}))
}
