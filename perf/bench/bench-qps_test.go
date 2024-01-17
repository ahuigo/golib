package gotest

// go test -timeout 30s -run ^TestCustomQps$ . -v

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"

	"sync"
	"time"
)

type Args struct {
	TaskName string
	Num      int
	Cn       int
}

func init() {
	rand.Seed(time.Now().Unix())

}

func getCliArgs() (args Args) {
	tasknamePtr := flag.String("t", "yxh_task0", "taskname")
	numPtr := flag.Int("n", 200, "total number of test tasks")
	cnPtr := flag.Int("c", 100, "concurrent executed tasks ")
	flag.Parse()

	args.TaskName = *tasknamePtr
	args.Num = *numPtr
	args.Cn = *cnPtr

	// dir
	_ = flag.Args()
	return args
}

// 模拟http
func httpRequest(taskname string, i int) error {
	if i%400 == 0 {
		fmt.Printf("i=%+v\n", i)
	}

	ri := 1 + rand.Intn(1000)
	// fmt.Printf("i=%+v,ri=%v\n", i, ri)
	time.Sleep(time.Duration(ri) * time.Millisecond)
	return nil
}

var succNum int = 0

func TestCustomQps(t *testing.T) {
	args := getCliArgs()
	fmt.Printf("args:%+v", args)

	var wg sync.WaitGroup
	ch := make(chan struct{}, args.Cn)

	var st = time.Now()
	for i := 0; i < args.Num; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-ch
                succNum++
			}()
			// mock request
			httpRequest(args.TaskName, i)
		}(i)

	}
	wg.Wait()
	elapsed := time.Since(st)
	fmt.Printf("elapsed:%v, succ=%d,total=%d", elapsed, succNum, args.Num)

}
