package t

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
)

func TestGoJson(t *testing.T) {
	v := []byte("ahui中国")
	b, err := json.Marshal(v)
	fmt.Println(string(b))
	fmt.Println(err)
}
