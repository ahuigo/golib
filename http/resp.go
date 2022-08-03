package main

import (
	"fmt"
	"net/http"
	"log"
	"bytes"
)


func main() {

	resp, err := http.Get("https://httpbin.org/get?ahui=mo")
	if err != nil {
		// handle error
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto %s\n", resp.Proto)
	fmt.Printf("resp content length %d\n", resp.ContentLength)
	fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)
	fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)
    fmt.Println("----------\nbody:")
    buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
