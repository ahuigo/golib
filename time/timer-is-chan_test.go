package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIsChan(t *testing.T) {
	var timerC <-chan time.Time // timerC

	for {
		fmt.Printf("timerC1:%v\n", timerC)
		if timerC == nil {
			fmt.Println("reset timerC")
			timerC = time.After(5 * time.Second)
		}
		fmt.Printf("timerC2:%v\n", timerC)

		select {
		case <-timerC:
			fmt.Println("5s over!")
			timerC = nil // set to nil!!!!!!!!!!!!
		}

		fmt.Println("loop go")
		time.Sleep(time.Second) // simulate work
	}
}
