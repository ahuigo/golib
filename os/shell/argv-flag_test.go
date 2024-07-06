package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestFlagUsage(t *testing.T) {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	// go run regexp.go -h
	cmd := flag.String("cmd", "compile", "a string")
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	fmt.Println(*cmd, *port)
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		// flag.PrintDefaults()
	}
	flag.Usage()
}
