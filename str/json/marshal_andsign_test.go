package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalAddSign(t *testing.T) {
	c := "a&b"
	if out, err := json.Marshal(c); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
}
