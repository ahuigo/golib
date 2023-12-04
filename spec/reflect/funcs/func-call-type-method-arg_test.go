package funcs

import (
	"fmt"
	"reflect"
	"testing"
)

func unmarshalJSONWraper(data []byte, objIface interface{}) error {
	ifaceType := reflect.TypeOf(objIface)
	unmarshalJSONMethod, ok := ifaceType.MethodByName("UnmarshalJSON")
	if !ok {
		panic("UnmarshalJSON method not found")
		// return json.Unmarshal(data, iface)
	}

	args := []reflect.Value{reflect.ValueOf(data)}

	// 或者转换成正常的UnmarshalJSON方法
	//fn := unmarshalJSONMethod.Func.Interface().(func([]byte, obj any) error)

	// 调用UnmarshalJSON方法
	ifaceValue := reflect.ValueOf(objIface)
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
