package demo

import (
	"fmt"
	"reflect"
	"strings"
)

// Refer to github.com/thoas/go-funk library
func contains(in interface{}, elem interface{}) bool {
	inValue := reflect.ValueOf(in)
	elemValue := reflect.ValueOf(elem)
	inType := inValue.Type()

	switch inType.Kind() {
	case reflect.String:
		return strings.Contains(inValue.String(), elemValue.String())
	case reflect.Map:
		equalTo := equal(elem, true)
		for _, key := range inValue.MapKeys() {
			if equalTo(key, inValue.MapIndex(key)) {
				return true
			}
		}
	case reflect.Slice, reflect.Array:
		equalTo := equal(elem)
		for i := 0; i < inValue.Len(); i++ {
			if equalTo(reflect.Value{}, inValue.Index(i)) {
				return true
			}
		}
	default:
		panic(fmt.Sprintf("Type %s is not supported by Contains, supported types are String, Map, Slice, Array", inType.String()))
	}

	return false
}

func equal(expectedOrPredicate interface{}, optionalIsMap ...bool) func(keyValueIfMap, actualValue reflect.Value) bool {
	isMap := append(optionalIsMap, false)[0]

	if IsFunction(expectedOrPredicate) {
		inTypes := []reflect.Type{nil}
		if isMap {
			inTypes = append(inTypes, nil)
		}

		if !IsPredicate(expectedOrPredicate, inTypes...) {
			panic(fmt.Sprintf("Predicate function must have %d parameter and must return boolean", len(inTypes)))
		}

		predicateValue := reflect.ValueOf(expectedOrPredicate)

		return func(keyValueIfMap, actualValue reflect.Value) bool {

			if isMap && !keyValueIfMap.Type().ConvertibleTo(predicateValue.Type().In(0)) {
				panic("Given key is not compatible with type of parameter for the predicate.")
			}

			if (isMap && !actualValue.Type().ConvertibleTo(predicateValue.Type().In(1))) ||
				(!isMap && !actualValue.Type().ConvertibleTo(predicateValue.Type().In(0))) {
				panic("Given value is not compatible with type of parameter for the predicate.")
			}

			args := []reflect.Value{actualValue}
			if isMap {
				args = append([]reflect.Value{keyValueIfMap}, args...)
			}

			return predicateValue.Call(args)[0].Bool()
		}
	}

	expected := expectedOrPredicate

	return func(keyValueIfMap, actualValue reflect.Value) bool {
		if isMap {
			actualValue = keyValueIfMap
		}

		if expected == nil || actualValue.IsZero() {
			return actualValue.Interface() == expected
		}

		return reflect.DeepEqual(actualValue.Interface(), expected)
	}
}

// IsFunction returns if the argument is a function.
func IsFunction(in interface{}, num ...int) bool {
	funcType := reflect.TypeOf(in)

	result := funcType != nil && funcType.Kind() == reflect.Func

	if len(num) >= 1 {
		result = result && funcType.NumIn() == num[0]
	}

	if len(num) == 2 {
		result = result && funcType.NumOut() == num[1]
	}

	return result
}

// IsPredicate returns if the argument is a predicate function.
func IsPredicate(in interface{}, inTypes ...reflect.Type) bool {
	if len(inTypes) == 0 {
		inTypes = append(inTypes, nil)
	}

	funcType := reflect.TypeOf(in)

	result := funcType != nil && funcType.Kind() == reflect.Func

	result = result && funcType.NumOut() == 1 && funcType.Out(0).Kind() == reflect.Bool
	result = result && funcType.NumIn() == len(inTypes)

	for i := 0; result && i < len(inTypes); i++ {
		inType := inTypes[i]
		result = inType == nil || inType.ConvertibleTo(funcType.In(i))
	}

	return result
}
