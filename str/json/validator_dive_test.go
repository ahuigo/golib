package t

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/goccy/go-json"
)

func TestValidator(t *testing.T) {
	type Request struct {
		Args []string `json:"args" validate:"required,dive,required"`
	}

	// 假设我们有一个JSON请求
	jsonStr := `{"args": ["foo", "bar", "baz",""]}`

	// 将JSON解析为Request结构体
	var req Request
	err := json.Unmarshal([]byte(jsonStr), &req)
	t.Logf("%#v\n", req)
	if err != nil {
		t.Fatal(err)
	}

	// 使用validator验证Request结构体
	validate := validator.New()
	err = validate.Struct(req)
	if err == nil {
		t.Fatal("not validate")
	}
	t.Log(err)
}
