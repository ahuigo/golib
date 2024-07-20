package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	var x = []byte{1, 2, 3, 4, 5, 6}
	var y = []int8{1, 2, 3, 4, 5, 6}
	xBytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	yBytes, err := json.Marshal(y)
	if err != nil {
		panic(err)
	}
	fmt.Printf("uint8 %s, int8: %s", string(xBytes), string(yBytes))
}
