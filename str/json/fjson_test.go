package t

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

// bench: https://www.reddit.com/r/golang/comments/16u956j/the_latest_json_encodedecode_benchmarks_and/

// https://github.com/dtgorski/jsonlex: Fast JSON lexer (tokenizer) with no memory footprint and no garbage collector pressure (zero heap allocations).
func TestSonic(t *testing.T) {
	type Stu struct {
		Name []byte
		age  int
	}
	v := Stu{[]byte("ahui"), 1}
	b, err := sonic.Marshal(&v)
	fmt.Println(string(b), err)
	v = Stu{}
	err = sonic.Unmarshal(b, &v)
	fmt.Println(v, string(v.Name), err)

}

func TestFjson(t *testing.T) {
	type Stu struct {
		Name []byte
		age  int
	}
	v := Stu{[]byte("ahui"), 1}
	b, err := json.Marshal(&v)
	fmt.Println(string(b), err)
	v = Stu{}
	err = json.Unmarshal(b, &v)
	fmt.Println(v, string(v.Name), err)

}
