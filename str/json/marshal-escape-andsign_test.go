package t

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalEscapeHTML(t *testing.T) {
	var c string = "a&b" //and sign: & => \u0026
	if out, err := json.Marshal(c); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
}
func TestMarshalNoescapeHTML(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	c := "a&b" //noescape: and sign: & => \u0026, no!!!!
	// c := json.RawMessage("a&b") // []byte 不base64, json 编码校验会出问题
	encoder := json.NewEncoder(buf) //为了美观，默认多一个换行符
	encoder.SetEscapeHTML(false)    // 这个不是base64, 为是html escape
	// encoder.SetIndent("", " ")

	if err := encoder.Encode(c); err != nil {
		panic(err)
	} else {
		fmt.Printf("out:%#v\n", buf.String())
	}
}
