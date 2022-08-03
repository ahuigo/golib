package t

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalRaw(t *testing.T) {
	contents := []byte(`
{
	"key": "value",
	"key2": [123456]
}
`)

	type Custom struct {
		Key  string
		Key2 json.RawMessage
	}

	customStructure := &Custom{}
	err := json.Unmarshal(contents, customStructure)
	t.Logf("%#v\n", customStructure)
	t.Logf("err:%s\n", err)

}
