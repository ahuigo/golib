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
reflect.NewAt + unsafe.Pointer 可以访问 unexported interface field(struct/ptr/...)
1. fmt.Println/Printf 就是这样的
2. 不可以写入, 只能读取

深度遍历序列可以用 github.com/ahuigo/gofnext/serial.DeepSerial
*/
func TestReflecPointerUnexported(t *testing.T) {
	m := pri.New()
	b, err := json.Marshal(m)
	fmt.Printf("json:%#v, %v\n", string(b), err)
	fmt.Printf("printf:%#v\n", m)

	v := reflect.ValueOf(m).Elem() // 取pointer 的 Elem

	fmt.Println("\nreflect:")
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		field := v.Field(i)

		// method1: Get the field, returns https://golang.org/pkg/reflect/#StructField
		pv := reflect.NewAt(fieldType.Type, unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()

		// method2(recommended): Get the field, returns https://golang.org/pkg/reflect/#Value
		rv := "<unknown>"
		switch field.Kind() { // 如果直接用field.Interface() 会报错: unexported field or method
		case reflect.String:
			rv = `"` + field.String() + `"`
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv = fmt.Sprintf("%d", field.Int())
		case reflect.Struct, reflect.Ptr:
			if field.CanInterface() {
				fi := field.Interface()
				rv = fmt.Sprintf("%#v", fi)
			} else {
				rv = "<cannot return value from unexported field or method>"
			}

		}

		fmt.Printf("k:%s, pv:%v, rv:%v\n", fieldType.Name, pv, rv)
	}
}

/*
 * 非pointer struct 不能访问 unexported field(只能先转成pointer)
 */
func TestReflectStructUnexported(t *testing.T) {
	val2 := "world"
	m := struct {
		val  int
		val2 *string
		Name string
	}{1, &val2, "hello"}
	b, err := json.Marshal(m)
	fmt.Printf("json:%#v, %v\n", string(b), err)
	fmt.Printf("print:%#v\n", m)

	// v := reflect.ValueOf(m).Elem() // 取pointer 的 Elem
	v := reflect.ValueOf(m)

	fmt.Println("\nreflect:")
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		field := v.Field(i)

		// method2(recommended): Get the field, returns https://golang.org/pkg/reflect/#Value
		rv := "<unknown>"
		switch field.Kind() {
		case reflect.String:
			rv = `"` + field.String() + `"`
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv = fmt.Sprintf("%d", field.Int())
		case reflect.Struct, reflect.Ptr:
			if field.CanInterface() {
				fi := field.Interface()
				rv = fmt.Sprintf("%#v", fi)
			} else {
				/* error: reflect.Value.UnsafeAddr of unaddressable value (xx)
				1. 只有指针类型的值才能获取其字段的地址。非指针类型的值是不可寻址的
				2. 转成pointer　才能访问通过UnsafeAddr访问到 unexported field
				*/
				// fi := reflect.NewAt(fieldType.Type, unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
				// rv = fmt.Sprintf("%#v", fi)
				rv = "<cannot return value from unexported field or method>"
			}

		}

		fmt.Printf("k:%s, rv:%v\n", fieldType.Name, rv)
	}
}

func TestReflectStructUnexportedConvertPtr(t *testing.T) {
	val2 := "world"
	m := struct {
		val  int
		val2 *string
		Name string
	}{1, &val2, "hello"}
	b, err := json.Marshal(m)
	fmt.Printf("json:%#v, %v\n", string(b), err)
	fmt.Printf("print:%#v\n", m)

	v1 := reflect.ValueOf(m)
	// 只有point类型的值才能获取其field的地址。非指针类型的值是不可寻址的
	vPtr := reflect.New(v1.Type()).Elem()
	vPtr.Set(v1)

	fmt.Println("\nreflect:")
	for i := 0; i < vPtr.NumField(); i++ {
		fieldType := vPtr.Type().Field(i)
		field := vPtr.Field(i)

		pv := reflect.NewAt(fieldType.Type, unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()

		// method2(recommended): Get the field, returns https://golang.org/pkg/reflect/#Value
		rv := "<unknown>"
		switch field.Kind() {
		case reflect.String:
			rv = `"` + field.String() + `"`
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv = fmt.Sprintf("%d", field.Int())
		case reflect.Struct, reflect.Ptr:
			if field.CanInterface() {
				fi := field.Interface()
				rv = fmt.Sprintf("%#v", fi)
			} else {
				rv = "<cannot return value from unexported field or method>"
			}

		}

		fmt.Printf("k:%s, rv:%v, pv:%v\n", fieldType.Name, rv, pv)
	}
}
