package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimerStop(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	count := 0

	go func() {
		for {
			<-timer.C //quit goroutine if stop timmer
			count += 1
			fmt.Println("Timer count: ", count)
			if count == 2 {
				//timer.Reset(150*time.Millisecond)//will stop
				timer.Reset(50 * time.Millisecond) //not stop
				fmt.Println("Timer sleep 2(0),", count)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Timer sleep 2(1),", count)
				continue
			}
			if count >= 3 {
				fmt.Println("Timer break 3,", count)
				break
			}
			timer.Reset(time.Second)
		}
		fmt.Println("Timer quit: ", count)
	}()
	time.Sleep(2100 * time.Millisecond)
	fmt.Printf("sleep stop before: 1\n")
	timer.Stop() //调用 <-timer.C (期间)会stop 协程(timer内ttl需要大于0)
	fmt.Printf("sleep stop after: 2\n")

	time.Sleep(5 * time.Second)
	println("quit")

}
