package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalCase(t *testing.T) {
	contents := []byte(`
{
	"$schema": "../../../.schema/relation_tuple.schema.json",
	"key": "value",
	"key2": [123456],
	"age_a": "18",
	"age_b": "19",
	"AgE_c": "19",
    "has": null,
	"AgE_d": "19"
}
`)

	type Custom struct {
		Key   string
		Key2  json.RawMessage
		Age_a int `json:",string"`
		Age_B int `json:",string"`
		Age_C int `json:",string"`
		AgeD  int `json:",string"` // only　this　one　can not　be　unmarshaled
		Has  *bool `json:"has"`
	}

	customStructure := &Custom{}
	err := json.Unmarshal(contents, customStructure)
	fmt.Printf("%#v\n", customStructure)
    fmt.Printf("has:%#v, isNil:%v\n", customStructure.Has, customStructure.Has==nil)
	if err != nil {
		t.Logf("err:%s\n", err.Error())
	} else {
		t.Logf("no err:%v\n", err == nil)
	}

}
