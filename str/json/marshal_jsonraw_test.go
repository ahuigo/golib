package t

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	var c interface{} = json.RawMessage(`"x"\n`)
	if v, ok := c.(json.RawMessage); ok {
		c = string(v)
	}
	if _, err := json.Marshal(c); err != nil {
		panic(err)
	}
}
