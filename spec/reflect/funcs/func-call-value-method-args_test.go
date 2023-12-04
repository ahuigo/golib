package funcs

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo struct {
}

func (c *Foo) Bar(age int, args ...any) string {
	return fmt.Sprintf("age:%d, args:%v", age, args)
}

func TestFuncCallValueArgs(t *testing.T) {
	hashKeyFunc := func(keys ...any) string {
		foo := &Foo{}
		objRefV := reflect.ValueOf(foo)
		methodValue := objRefV.MethodByName("Bar")

		if methodValue.IsValid() {
			argsNum := methodValue.Type().NumIn() // len(keys)
			fmt.Printf("args num:%d\n", argsNum)  // 变参数是1
			// convert HashKeyFunc 方法
			fn := methodValue.Interface().(func(int, ...any) string)
			_ = fn

			// 调用 HashKeyFunc 方法
			reflectKeys := make([]reflect.Value, argsNum)
			reflectKeys[0] = reflect.ValueOf(29)
			reflectKeys[1] = reflect.ValueOf(keys)
			result := methodValue.Call(reflectKeys)

			// 将结果转换为 string
			if len(result) > 0 {
				if bytes, ok := result[0].Interface().(string); ok {
					return bytes
				}
			}
		}

		// 其他类型的处理逻辑
		return "null"
	}
	println(hashKeyFunc(1, 2, 3))

}
