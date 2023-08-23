package pkg

import (
	"fmt"
	"testing"
	"time"
)

/**
两种deadlock
1. read dst, 但是dst 没有输入就会是deadlock (除非close(dst))
1. write dst,但是dst 没有read就会是deadlock

***/

func TestDeadLock(t *testing.T) {
	dst := make(chan int)
	n := 1
	go func() {
		dst <- 1
		time.Sleep(1 * time.Second)
		//close(dst) //防止 fatal error: all goroutines are asleep - deadlock!
	}()

	for n = range dst { //read dst, 但是dst 没有输入就会是deadlock
		fmt.Println(n)
		if n == 4 {
			break
		}
	}
}
