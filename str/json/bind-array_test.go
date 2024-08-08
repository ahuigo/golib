package t

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBindArray(t *testing.T) {
	type Obj struct {
		Name    string   `json:"name"`
		PyFiles []string `json:"py_files" `
	}
	var obj Obj
	data := []byte(`{"name":"ahuigo","apy_files": ["i"]}`)
	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&obj); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", obj)
}
