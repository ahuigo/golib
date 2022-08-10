package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func ExampleMapWriteLock() {
	m := make(map[string]interface{})
	//并发写：fatal error: concurrent map writes
	go func() {
		for {
			m["abc"] = 134
			time.Sleep(1)
		}
	}()
	go func() {
		for {
			m["abc"] = User{}
			time.Sleep(1)
		}
	}()

	// 读和写引发: fatal error: concurrent map iteration and map write
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(m)
	}

}
