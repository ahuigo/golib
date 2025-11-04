package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "cpu.pprof", "write cpu profile `file`")

var memprofile = flag.String("memprofile", "mem.pprof", "write memory profile to `file`")

// time.Sleep 函数让当前的 goroutine 进入休眠状态，这种情况下 CPU 将不会执行任何操作，因此在 CPU profile 中是看不到 time.Sleep 的
// 可以使用 runtime/trace 工具
var traceprofile = flag.String("traceprofile", "trace.pprof", "write memory profile to `file`")

func longfun2() {
	for i := 1; i < 2e3; i++ {
		for i := 0; i < 1e6; i++ {
		}
	}
}
func longfun1(max int) {
	for i := 0; i < max; i++ {
		for i := 0; i < 1e6; i++ {
		}
	}
}

func fetchHttpSlow() {
	resp, err := http.Get("http://m:4500/sleep/1")
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Read(make([]byte, 1024))
	defer resp.Body.Close()
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
func hello() {
	fmt.Printf("hello world1!\n")
	start := time.Now()
	fetchHttpSlow()
	sleep()
	longfun1(2e3)
	fmt.Println("longfun1 time:", time.Since(start))
	sleep()
	longfun2()
	fmt.Println("longfun2 time:", time.Since(start))
}

func main() {
	flag.Parse()
	// trace.pporf
	if f, _ := os.Create(*traceprofile); f != nil {
		// go tool trace -http=:9999 trace.pprof
		defer f.Close()
		trace.Start(f)
		defer trace.Stop()
	}
	// cpu.pprof
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		// StartCPUProfile(os.Stdout)
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	ctx := context.Background()
	region := trace.StartRegion(ctx, "main region")
	// ... rest of the program ...
	hello()
	region.End()

	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create memory profile: ", err)
	// 	}
	// 	runtime.GC() // get up-to-date statistics
	// 	if err := pprof.WriteHeapProfile(f); err != nil {
	// 		log.Fatal("could not write memory profile: ", err)
	// 	}
	// 	f.Close()
	// }
}
