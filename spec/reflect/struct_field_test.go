package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructField(t *testing.T) {
	type StructB struct {
		Age  int
		Name string
	}

	pf := fmt.Printf
	b := StructB{Age: 100}
	s := reflect.ValueOf(b)

	pf("KindType:%T, KindValue: %#v \n", s.Kind(), s.Kind())
	pf("TypeType:%T, TypeValue: %#v\n", s.Type(), s.Type())
	pf("NumFiled:%v \n", s.NumField())
	pf("NumMethods:%v \n", s.NumMethod())
	f0 := s.Field(0)
	pf("Field0:%v, T:%T \n", f0, f0)
	pf("Field0.Type().Name():name=%#v, T=%T \n", f0.Type().Name(), f0.Type().Name())
	pf("FieldByName(`age`):value=%v, T:%T \n", s.FieldByName("Age"), s.FieldByName("Age"))

	pf("-------set struct.Age---------------------\n")
	rv := reflect.ValueOf(12)
	refpb := reflect.Indirect(reflect.ValueOf(&b))
	refpb = reflect.ValueOf(&b).Elem()
	refpb.FieldByName("Age").Set(rv)
	pf("structS:%+v\n", s)
	pf("originB:%+v\n", b)
}
