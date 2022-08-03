package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
	"regexp"
    "flag"
)

func compile(){
	var validID = regexp.MustCompile(`http(s)?://([\w\-]+\.hdmap\.momenta\.works|localhost|m)(:\d+)?$`)
	validID = regexp.MustCompile(`^([a-z\d]+-)?osm[a-z\d\-]+\.hdmap\.momenta\.works$`)
	fmt.Println(validID.MatchString("osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("dev-osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("staging-osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("staging.osm3.hdmap.momenta.works"))  //true
	fmt.Println(regexp.MustCompile(`^\d+(\.\d+){3}$`).MatchString("1.23.1.34"))

}

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
    // go run regexp.go -h
    cmd := flag.String("cmd", "compile", "a string")
    flag.Parse()
    if *cmd == "compile" {
        compile()
    }
}
