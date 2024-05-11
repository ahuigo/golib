package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
	"regexp"
	"testing"
)

func nameMatch() {
	r := regexp.MustCompile(`^(?P<Year>\d{4})-(?P<Month>\d{2})`)

	res := r.FindStringSubmatch(`2015-05-27`) // res may be: []string(nil)
	names := r.SubexpNames()

	result := make(map[string]string)
	for i, _ := range res {
		if i != 0 {
			result[names[i]] = res[i]
		}
	}
	fmt.Printf("%#v\n", result)
}

func nameMatch2() {
	regex := *regexp.MustCompile(`(?s)(\d{4})-(\d{2})-(\d{2})`)
	txt := `2009-03-22
    2018-02-25`
	res := regex.FindAllStringSubmatch(txt, -1)
	for i := range res {
		//like Java: match.group(1), match.gropu(2), etc
		fmt.Printf("year: %s, month: %s, day: %s\n", res[i][1], res[i][2], res[i][3])
	}
}
func TestGroup(t *testing.T) {
	nameMatch()
	fmt.Println("-------")
	nameMatch2()
}
