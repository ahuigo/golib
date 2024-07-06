package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	var validID = regexp.MustCompile(`http(s)?://([\w\-]+\.hdmap\.momenta\.cn|localhost|m)(:\d+)?$`)
	validID = regexp.MustCompile(`^([a-z\d]+-)?osm[a-z\d\-]+\.hdmap\.momenta\.cn$`)
	fmt.Println(validID.MatchString("osm3.hdmap.momenta.cn"))         //true
	fmt.Println(validID.MatchString("dev-osm3.hdmap.momenta.cn"))     //true
	fmt.Println(validID.MatchString("staging-osm3.hdmap.momenta.cn")) //true
	fmt.Println(validID.MatchString("staging.osm3.hdmap.momenta.cn")) //false

	fmt.Println(regexp.MustCompile(`^\d+(\.\d+){3}$`).MatchString("1.23.1.34")) //true
	fmt.Println(regexp.MustCompile(`a|b$`).MatchString("a1"))                   // true
	fmt.Println(regexp.MustCompile(`(a|b)$`).MatchString("a1"))                 // false
	fmt.Println(regexp.MustCompile(`//`).MatchString("a//b"))                 // true
}
