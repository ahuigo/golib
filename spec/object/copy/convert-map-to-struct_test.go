package demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MyStruct struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserId    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
}

func TestMap2Struct(t *testing.T) {
	m := make(map[string]interface{})
	m["id"] = "2"
	m["name"] = "jack"
	m["user_id"] = "123"
	m["created_at"] = 5
	fmt.Println(m)

	// convert map to json-string
	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))

	// convert json to struct
	s := MyStruct{}
	json.Unmarshal(jsonString, &s)
	fmt.Println(s)

}
