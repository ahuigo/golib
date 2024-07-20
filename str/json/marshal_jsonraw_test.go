package t

import (
	"encoding/json"
	"testing"
)

/*
默认: escapeHTML为true　会检查json语法, 使用json.RawMessage 很可能会出问题
源码：https://golang.org/src/encoding/json/encode.go go/1.22.1/libexec/src/encoding/json/encode.go

	func addrMarshalerEncoder(e *encodeState, v reflect.Value, opts encOpts) {
		b, err := m.MarshalJSON() //func (json.RawMessage) MarshalJSON()  本身不会转义
		out, err = appendCompact(out, b, opts.escapeHTML) 但是这里会转义
	}
*/
func TestMarshalJsonRawBadCase(t *testing.T) {
	c := json.RawMessage(`abc"`) // 不合法的json字符串
	if out, err := json.Marshal(c); err != nil {
		panic(err)
	} else {
		t.Logf("%s\n", string(out))
	}
}

// 默认json encoder 会转义特殊字符：参考 marshal-escape-andsign_test.go
func TestMarshalRawCase(t *testing.T) {
	data := struct {
		C string
	}{
		C: "a&b", // {"C":"a\u0026b"}
	}
	if out, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		t.Logf("%s\n", string(out))
	}

}
