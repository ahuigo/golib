package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	rvField := structValue.FieldByName(name)
    fmt.Printf("rvField:%T\n", rvField)

	if !rvField.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !rvField.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	val := reflect.ValueOf(value)
	if rvField.Type() != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	rvField.Set(val)
	return nil
}

type MyStruct struct {
	Name string
	Age  int64
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = int64(23)

	result1 := MyStruct{}
	result :=&result1
	err := result.FillStruct(myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
