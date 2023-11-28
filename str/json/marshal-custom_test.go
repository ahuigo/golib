package t

import (
	"encoding/json"
	"fmt"
	"testing"
    "strings"
)

type UserMarshalCustom struct {
    Name  string
    Array []uint8
}

func (t *UserMarshalCustom) MarshalJSON() ([]byte, error) {
    var array string
    if t.Array == nil {
        array = "null"
    } else {
        array = strings.Join(strings.Fields(fmt.Sprintf("%d", t.Array)), ",")
    }
    jsonResult := fmt.Sprintf(`{"Name":%q,"Array":%s}`, t.Name, array)
    return []byte(jsonResult), nil
}

func TestMarshalCustom(t *testing.T) {
     v := &UserMarshalCustom{"Go", []uint8{'h', 'e', 'l', 'l', 'o'}}
    out, _:= json.Marshal(v)
    fmt.Println(string(out))
}
