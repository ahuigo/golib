package t

import (
	"encoding/json"
	"fmt"
	"testing"
)



func TestMarshalName(t *testing.T) {
    type A struct {
        Name  string
        NameB  string
        Has *bool `json:"has"`
    }
     v := &A{"", "", nil}
    out, _:= json.Marshal(v)
    fmt.Println(string(out)) //{"Name":,"NameB":}
}
