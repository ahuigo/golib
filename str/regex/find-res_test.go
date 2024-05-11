package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
	"regexp"
	"testing"
)

func findStr() {
	r := regexp.MustCompile(`^(?P<Year>\d{4})-(?P<Month>\d{2})`)

	res := r.FindString(`2015-05-27`) // res : "2015-05"

	fmt.Printf("res:%#v\n", res)
}

func TestFindResult(t *testing.T) {
	findStr()
}
