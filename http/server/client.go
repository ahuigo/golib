package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var httpClient *http.Client

func main() {
	n := 1000
	if len(os.Args) > 1 {
		if x, err := strconv.Atoi(os.Args[1]); err == nil {
			n = x
		}
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 2 * n,
			MaxIdleConns:        2 * n,
		},
		Timeout: 30 * time.Second,
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			do()
			wg.Done()
		}()
	}

	wg.Wait()
}

func do() {
	const url = "http://127.0.0.1:7777"
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}
