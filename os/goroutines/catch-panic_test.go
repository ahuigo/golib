package pkg

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func task(t int) (int, error) {
	time.Sleep(time.Duration(randIntn(100)) * time.Millisecond)
	if t <= 0 {
		return 0, nil
	} else if t == 1 {
		return 1, errors.Errorf("some error(task:%d)", t)
	} else {
		s := fmt.Sprintf("(task:%d)", t)
		panic("panic fatal error" + s)
	}
}

type Resp struct {
	n   int
	err error
}

// Server Request: catch routine panic
func handler() (n int, err error) {
	ch1 := make(chan Resp)
	ch2 := make(chan Resp)

	// start task
	go func() {
		fmt.Println("task1 start...")
		n, err := task(1)
		ch1 <- Resp{n, err}
		fmt.Println("task1 end")
	}()
	go func() {
		var n int
		var err error
		defer func() {
			if err := recover(); err != nil {
				err := errors.Errorf("catch task2 panic:%v", err)
				ch2 <- Resp{n, err}
			}
		}()
		fmt.Println("task2 start...")
		n, err = task(2)
		ch2 <- Resp{n, err}
		fmt.Println("task2 end")
	}()

	// get response
	var resp Resp
	select {
	case resp = <-ch1:
		fmt.Printf("ch1:%#v\n", resp.err)
		return resp.n, resp.err
	case resp2 := <-ch2:
		fmt.Printf("ch2:%#v\n", resp2.err)
		return resp2.n, resp2.err
	}
}

func TestCatchPanic(t *testing.T) {
	n, err := handler()
	t.Logf("n=%d,err=%v", n, err)
	time.Sleep(time.Hour)
}

func randIntn(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s).Intn(n)
}
