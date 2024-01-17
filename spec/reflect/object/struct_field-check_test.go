package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructFieldCheck(t *testing.T) {
	type StructB struct {
		Age  int
		Name string
	}

	pf := fmt.Printf
	b := StructB{Age: 100}
	val := reflect.ValueOf(b)

	// Check if the struct has a field named 'Transport'
	fieldVal := val.FieldByName("Age")
	if !fieldVal.IsValid() {
		fmt.Println("hdtiRaw does not have a field named 'Transport'")
		return
	}
	pf("age:%d\n", fieldVal.Interface().(int))

}
