package t

import (
	"encoding/json"
	"strconv"
	"testing"
)

type IntStr struct {
	value int
}

type Item struct {
	Price IntStr `json:"price"`
	Age   StrInt `json:"age"`
}

type StrInt int

func (v *StrInt) UnmarshalJSON(b []byte) (err error) {
	s, n := "", float64(0)
	if err = json.Unmarshal(b, &s); err == nil {
		intVar, _ := strconv.Atoi(s)
		*v = StrInt(intVar)
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		*v = StrInt(n)
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
	item := &Item{}
	bytes := []byte(`{"price":"200","age":"10"}`)
	_ = json.Unmarshal(bytes, item)
	t.Logf("%#v\n", item)
}
