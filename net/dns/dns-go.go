package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	n := int(1e6) //100万次
	fmt.Println("压测开始:", time.Now())
	wg.Add(n)
	for i := 0; i < n; i++ {
		go lookup(&wg)
	}
	wg.Wait()
	fmt.Println("压测结束:", time.Now())

}

func lookup(wg *sync.WaitGroup) {
	ips, err := net.LookupIP("mydomain.com")
	defer wg.Done()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	return
	for _, ip := range ips {
		fmt.Printf("%s\n", ip.String())
	}
}
