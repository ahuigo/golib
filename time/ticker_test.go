package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick2 at")
		}
	}
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
