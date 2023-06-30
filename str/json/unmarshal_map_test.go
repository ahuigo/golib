package t

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalMap(t *testing.T) {
	data := []byte(`
{
	"key": "value",
	"key2": [123456]
}
`)

	var obj interface{}
	err := json.Unmarshal(data, &obj)
	t.Logf("%#v\n", obj)
	t.Logf("err:%s\n", err)

}
