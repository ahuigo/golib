package unexported

import (
	"encoding/json"
	"fmt"
	"go-lib/spec/reflect/unexported/pri"
	"reflect"
	"testing"
	"unsafe"
)

/*
*
reflect.NewAt + unsafe.Pointer 可以访问 unexported field
1. fmt.Println/Printf 就是这样的
2. 不可以写入, 只能读取

深度遍历序列可以用 github.com/ahuigo/gofnext/serial.DeepSerial
*/
func TestReflecPointerUnexported(t *testing.T) {
	m := pri.New()
	b, err := json.Marshal(m)
	fmt.Printf("json:%#v, %v\n", string(b), err)
	fmt.Printf("print:%#v\n", m)

	v := reflect.ValueOf(m).Elem() // 取pointer 的 Elem

	fmt.Println("\nreflect:")
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		field := v.Field(i)

		// method1: Get the field, returns https://golang.org/pkg/reflect/#StructField
		pv := reflect.NewAt(fieldType.Type, unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()

		// method2(recommended): Get the field, returns https://golang.org/pkg/reflect/#Value
		rv := "<unknown>"
		switch field.Kind() {
		case reflect.String:
			rv = `"` + field.String() + `"`
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv = fmt.Sprintf("%d", field.Int())
		}

		fmt.Printf("k:%s, pv:%v, rv:%v\n", fieldType.Name, pv, rv)
	}
}
func TestReflectStructUnexported(t *testing.T) {
	m := struct {
		val  int
		Name string
	}{1, "hello"}
	b, err := json.Marshal(m)
	fmt.Printf("json:%#v, %v\n", string(b), err)
	fmt.Printf("print:%#v\n", m)

	// v := reflect.ValueOf(m).Elem() // 取pointer 的 Elem
	v := reflect.ValueOf(m)

	fmt.Println("\nreflect:")
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		field := v.Field(i)

		// error!!!!
		// pv := reflect.NewAt(fieldType.Type, unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()

		// method2(recommended): Get the field, returns https://golang.org/pkg/reflect/#Value
		rv := "<unknown>"
		switch field.Kind() {
		case reflect.String:
			rv = `"` + field.String() + `"`
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv = fmt.Sprintf("%d", field.Int())
		}

		fmt.Printf("k:%s, rv:%v\n", fieldType.Name, rv)
	}
}
