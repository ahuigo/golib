package main

import (
	"testing"

	_ "m/pkg"

	"github.com/samber/lo"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	"os"

	_ "github.com/ahuigo/gofnext"
)

const src = `package foo
import "fmt"
import "github.com/samber/lo"

func test() {
	lo.Contains([]int{1, 2}, 1)
}



func Bar(s string) string { 
    fmt.Println("hello")
    test()
    return s + "-Foo" 
}
`

func TestX(t *testing.T) {
	lo.Contains([]int{1, 2}, 1)
	var goPath string = os.Getenv("GOPATH")
	i := interp.New(interp.Options{GoPath: goPath, Env: os.Environ()})
	i.Use(stdlib.Symbols)
	// i.ImportUsed()
	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}
	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	println(r)
}
