package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalAndSign(t *testing.T) {
	c := "a&b" //and sign: & => \u0026
	if out, err := json.Marshal(c); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
}
