package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ztimer := time.NewTimer(0)
	for range ztimer.C {
		//ztimer.Reset(0 * time.Millisecond)
		fmt.Println("haha") // immediately
		break               // a timer only fires once
	}
}
func TestTimerTiker(t *testing.T) {
	ztimer := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-ztimer.C: // 不会累积
			//ztimer.Reset(0 * time.Millisecond)
			fmt.Println("haha")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
