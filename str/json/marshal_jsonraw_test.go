package t

import (
	"encoding/json"
	"testing"
)

func TestMarshalJsonRawString(t *testing.T) {
	var c interface{} = json.RawMessage(`"x"\n`)
	if v, ok := c.(json.RawMessage); ok {
		c = string(v) // 必须转换为string，否则会输出error
	}
	if out, err := json.Marshal(c); err != nil {
		panic(err)
	} else {
		t.Logf("%s\n", string(out))
	}

}

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
