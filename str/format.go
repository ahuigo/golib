package main

import (
	"fmt"
	"text/template"
	"strings"
)

type String string
func (s String) Format(data map[string]interface{}) (out string, err error) {
	t := template.Must(template.New("").Parse(string(s)))
    builder := &strings.Builder{}
	if err = t.Execute(builder, data); err != nil {
		return
	}
	out = builder.String()
    return
}


func main() {
    const tmpl = `Hi {{.Name}}!  {{range $i, $r := .Roles}}{{if $i}}, {{end}}{{.}}{{end}}`
    //const tmpl = `Hi {{.Name}}!  {{range $i, $r := .Roles}}{{if $i}}, {{end}}{{$r}}{{end}}`
	data := map[string]interface{}{
		"Name":     "Bob",
		"Roles":    []string{"dbteam", "uiteam", "tester"},
	}

	s ,_:= String(tmpl).Format(data)
	fmt.Println(s)
	fmt.Println("length=", len("1中"))
}
