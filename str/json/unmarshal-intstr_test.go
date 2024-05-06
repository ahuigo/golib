package t

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
)

type IntStr struct {
	value int
}
type Int2 int

type Item struct {
	Price IntStr `json:"price"`
	Age   Int2   `json:"age"`
}

func (v *Int2) UnmarshalJSON(b []byte) (err error) {
	s, n := "", float64(0)
	if err = json.Unmarshal(b, &s); err == nil {
		intVar, _ := strconv.Atoi(s)
		*v = Int2(intVar)
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		*v = Int2(n)
	}
	return
}
func (v *IntStr) UnmarshalJSON(b []byte) (err error) {
	s, n := "", float64(0)
	if err = json.Unmarshal(b, &s); err == nil {
		intVar, _ := strconv.Atoi(s)
		v.value = intVar
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		v.value = int(n)
	}
	return
}

func TestUnmarshalCustom(t *testing.T) {
	item1 := &Item{}
	item2 := &Item{}
	item3 := &Item{}
	bytes := []byte(`{"price":"200","age":"10"}`)
	_ = json.Unmarshal(bytes, item1)
	_ = gojson.Unmarshal(bytes, item2)
	_ = sonic.Unmarshal(bytes, item3)
	t.Logf("%#v\n", item1)
	t.Logf("%#v\n", item2)
	t.Logf("%#v\n", item3)
}
