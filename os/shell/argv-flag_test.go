package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

type headers struct{ http.Header }

func (h headers) String() string {
	buf := &bytes.Buffer{}
	if err := h.Write(buf); err != nil {
		return ""
	}
	return buf.String()
}

// Set implements the flag.Value interface for a map of HTTP Headers.
func (h headers) Set(value string) error {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("header '%s' has a wrong format", value)
	}
	key, val := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	if key == "" || val == "" {
		return fmt.Errorf("header '%s' has a wrong format", value)
	}
	h.Header[key] = append(h.Header[key], val)
	return nil
}

func TestFlagUsage(t *testing.T) {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	// go run regexp.go -h
	port := 0
	header := headers{http.Header{}}
	os.Args = []string{"./main", "-cmd", "del", "-p", "4500", "-header", `key: token1`}
	cmd := flag.String("cmd", "compile", "a string")
	flag.IntVar(&port, "p", 8080, "port to listen on")
	flag.Var(&header, "header", "http headers")
	flag.Parse()

	fmt.Println(*cmd, port, header)
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		// flag.PrintDefaults()
	}
	flag.Usage()
}
