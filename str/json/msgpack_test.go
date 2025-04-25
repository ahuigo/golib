package t

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

type Stu struct {
	I int
	i int
	M map[string]any
}

func MarshalMsgpack(v any) ([]byte, error) {
	return msgpack.Marshal(v)
}

// UnmarshalMsgpack 解析 MessagePack 格式的字节切片 data 并将结果存储在 v 指向的值中。
// v 必须是一个指向目标数据结构（如结构体、map、slice、基本类型等）的指针，类似于 encoding/json.Unmarshal。
// 如果 v 是 nil 或者不是指针，UnmarshalMsgpack 会返回错误。
//
// 对于数字类型：
//   - 当反序列化到 interface{} 或 map[string]any 时，msgpack 通常会尝试将整数
//     解码为 int64 或 uint64，浮点数解码为 float64，这比 json 默认将所有数字
//     解码为 float64 更能保留类型信息。
//   - 当反序列化到具体的类型（如 *int, *float32, *MyStruct）时，会进行相应的类型转换。
//
// 对于结构体：
//   - 默认情况下，它会匹配导出字段名（大小写敏感）或 'msgpack' 标签。
//   - 嵌套的结构体如果存在于 map[string]any 的值中，反序列化到 map[string]any 时
//     通常会被解码为 map[string]interface{}。
func UnmarshalMsgpack(data []byte, v any) error {
	return msgpack.Unmarshal(data, v)
}

func TestMsgpack(t *testing.T) {
	fn := func(v any) {
		buf, err := MarshalMsgpack(v)
		if err != nil {
			t.Fatalf("Msgpack encoding failed: %v", err)
		}
		fmt.Printf("Encoded data: %x\n", buf)
		var obj Stu
		err = UnmarshalMsgpack(buf, &obj)
		if err != nil {
			t.Fatalf("Msgpack decoding failed: %v", err)
		}
		fmt.Printf("Decoded map: %#v\n", obj)
		if obj.M != nil {
			fmt.Printf("Decoded map.M.u8: %T\n", obj.M["u8"])
		}
	}
	fn(Stu{I: 1, i: 1, M: map[string]any{
		"u8": uint8(8),
	}}) // 直接传递结构体
	var stu *Stu
	fn(stu) // 传递指向结构体的指针

}
func TestMsgpackDynamic(t *testing.T) {
	// !!! 重要：需要告诉编码器如何处理自定义结构体，否则它可能不知道如何编码 Stu
	// 选项1：如果 Stu 实现了 msgpack.CustomEncoder/Decoder 接口
	// 选项2：注册扩展类型 (较复杂)
	// 选项3：让它默认使用类似 map 的方式编码 (可能需要配置)
	// enc.UseJSONTag(true) // 或者其他配置，让它像 JSON 一样处理导出字段
	// 或者确保 Stu 字段导出，它可能默认按 map 处理
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)

	data := map[string]any{
		"int":      42,
		"u8":       uint8(8),
		"int8":     int8(8),
		"f32":      float32(3.14),
		"f64":      1.6,
		"stu":      Stu{I: 2, i: 2},
		"largeInt": 1234567890123456789,
	}
	type Data struct {
		M   map[string]any
		Stu Stu
	}
	err := enc.Encode(Data{
		M:   data,
		Stu: Stu{I: 1, i: 1},
	})
	if err != nil {
		t.Fatalf("Msgpack encoding failed: %v", err)
	}

	dec := msgpack.NewDecoder(&buf)
	var data1 Data
	err = dec.Decode(&data1) // 解码到 map[string]any
	if err != nil {
		t.Fatalf("Msgpack decoding failed: %v", err)
	}
	obj := data1.M
	fmt.Printf("Decoded map: %#v\n", data1)
	fmt.Printf("Type of obj[\"int\"]: %T\n", obj["int"]) // 可能是 int64 或 uint64
	fmt.Printf("Type of obj[\"u8\"]: %T\n", obj["u8"])   // 可能是 int64 或 uint64
	fmt.Printf("Type of obj[\"f32\"]: %T\n", obj["f32"]) // 可能是 float64 或 float32 (取决于库)
	fmt.Printf("Type of obj[\"f64\"]: %T\n", obj["f64"]) // 可能是 float64 或 float32 (取决于库)
	fmt.Printf("Type of obj[\"stu\"]: %T\n", data1.Stu)  // 可能是 map[string]interface{}
}
