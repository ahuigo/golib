package t

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalCoder(t *testing.T) {
	fc := []byte(`{
		"$schema": "../../../.schema/relation_tuple.schema.json",
		"namespace": "videos"
	  } `)

	type RelationTuple struct {
		Schema    string `json:"$schema"`
		Namespace string `json:"namespace"`
	}

	var r RelationTuple
	decoder := json.NewDecoder(bytes.NewReader(fc))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&r); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", r)

}
