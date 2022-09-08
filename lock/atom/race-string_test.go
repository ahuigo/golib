package main

// 	go run race-string.go | grep -E '^WHAT$$'
import (
	"fmt"
	"testing"
	"time"
)

// 交替打印字符
// go-lib/goroutine/race-string.go

func TestRaceString(t *testing.T) {
	const (
		FIRST  = "Hello world!"
		SECOND = "F*CK"
	)
	var s string = FIRST
	go func() {
		i := 1
		for {
			i = 1 - i
			if i == 0 {
				s = FIRST
			} else {
				s = SECOND
			}
			time.Sleep(10 * time.Nanosecond)
			//time.Sleep(10*time.Second)
		}
	}()

	for {
		s1 := s
		if s1 == FIRST || s1 == SECOND {
			fmt.Println(s)
		} else {
			fmt.Printf("---------------\nlast string:%v\n", s1)
			panic("err string")
		}
		time.Sleep(10 * time.Nanosecond)
	}
}
