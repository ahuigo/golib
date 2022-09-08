package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ztimer := time.NewTimer(0)
	for {
		select {
		case <-ztimer.C:
			//ztimer.Reset(0 * time.Millisecond)
			fmt.Println("haha")
		}
	}
}
