package t

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalDataType(t *testing.T) {
	contents := []byte(`{
	"key": "value",
	"key2": {}
}
`)

	customStructure := map[string]string{}
	err := json.Unmarshal(contents, &customStructure)
	t.Logf("%#v\n", customStructure)
	t.Logf("err:%s\n", err)

}
