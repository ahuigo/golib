package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	// var ch1 <-chan int
	ch1 := make(chan int, 1)
	ch1 <- 1
	fmt.Println("Tick start:", time.Now())
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick at:", time.Now())
		case <-ch1:
			fmt.Println("sleep 5")
			time.Sleep(5 * time.Second)
		}
	}
}
