package m

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	wg := sync.WaitGroup{}

	fn := func() {
		defer func() {
			wg.Done()
		}()
		fmt.Println(time.Now())
		time.Sleep(time.Second * 3)
	}
	count := 0
	for {
		begin := time.Now()
		fmt.Println("begin:\n", begin)
		wg.Add(1)
		go fn()
		wg.Wait()
		fmt.Println("end:\n", time.Now())
		elasped := time.Since(begin)
		if elasped < 3*time.Second {
			time.Sleep(3 * time.Second)
		}
		if count > 20 {
			break
		}
	}

}
