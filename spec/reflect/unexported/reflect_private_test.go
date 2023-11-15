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
*/
func TestReflectUnexported(t *testing.T) {
	m := pri.New()
	b, err := json.Marshal(m)
	fmt.Printf("%#v, %v\n", string(b), err)
	fmt.Printf("%#v\n", m)

	v := reflect.ValueOf(m).Elem() // 使用 m 而不是 *m

	for i := 0; i < v.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := v.Type().Field(i)
		value := reflect.NewAt(field.Type, unsafe.Pointer(v.Field(i).UnsafeAddr())).Elem().Interface()

		fmt.Printf("%s: %v\n", field.Name, value)
	}
}
