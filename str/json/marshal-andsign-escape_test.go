package t

import (
	"bytes"
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
func TestMarshalNoescape(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	c := "a&b"                      //noescape: and sign: & => \u0026, no!!!!
	encoder := json.NewEncoder(buf) //为了美观，默认多一个换行符
	encoder.SetEscapeHTML(false)
	// encoder.SetIndent("", " ")

	if err := encoder.Encode(c); err != nil {
		panic(err)
	} else {
		fmt.Printf("out:%#v\n", buf.String())
	}
}
