package t

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type StuGob struct {
	I int
	//i int // not supported
}

func init() {
	// 1. 注册StuGob, 2. StuGob 不能有private property
	gob.Register(StuGob{})
}

func TestGobStruct(t *testing.T) {
	// 各种不同类型的数值
	// 序列化
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	data := map[string]any{
		"int":    42,
		"u8":     uint8(8),
		"int8":   int8(8),
		"f32":    float32(8),
		"StuGob": StuGob{1},
	}
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}

	// 反序列化1
	dec := gob.NewDecoder(&buf)
	obj := map[string]any{}
	err = dec.Decode(&obj)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v, %T, %T, %#v\n", obj, obj["int"], obj["u8"], obj["StuGob"])
}
