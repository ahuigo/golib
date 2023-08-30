//go:build needTestxxx

// go test -tags=needTestxxx .
package demo

import "testing"

func TestSth(t *testing.T) {
	//  或者可以：t.Skip("skipping testing in short mode")
	if true:
		println("this will be ignored")
}
