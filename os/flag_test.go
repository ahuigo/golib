package o

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
    "flag"
    "testing"
)


func TestFlag(t *testing.T) {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
    // go run regexp.go -h
    cmd := flag.String("cmd", "compile", "a string")
    port := flag.String("port", "8080", "port to listen on")
    flag.Parse()
    fmt.Println(*cmd, *port)
}
