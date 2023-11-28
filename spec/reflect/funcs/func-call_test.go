package funcs

import (
	"fmt"
	"reflect"
	"testing"
)

func unmarshalJSONWraper(data []byte, iface interface{}) error {
	ifaceType := reflect.TypeOf(iface)
	unmarshalJSONMethod, ok := ifaceType.MethodByName("UnmarshalJSON")
	if !ok {
		panic("UnmarshalJSON method not found")
		// return json.Unmarshal(data, iface)
	}

	ifaceValue := reflect.ValueOf(iface)
	args := []reflect.Value{reflect.ValueOf(data)}

	// 调用UnmarshalJSON方法
	result := unmarshalJSONMethod.Func.Call([]reflect.Value{ifaceValue, args[0]})

	// 检查调用结果是否有错误
	if len(result) > 0 && !result[0].IsNil() {
		return result[0].Interface().(error)
	}

	return nil
}

type MyStruct struct {
	Name string `json:"name"`
}

func (m *MyStruct) UnmarshalJSON(data []byte) error {
	m.Name = string(data)
	return nil
}

func TestFuncCallUnmarshal(t *testing.T) {
	data := []byte(`"John"`)

	var my = &MyStruct{}

	_ = unmarshalJSONWraper(data, my)
	fmt.Println(my)
}
