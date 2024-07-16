package main

import (
	"flag"
	"fmt"
	"log"
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

func longfun1() {
	for i := 0; i < 2e3; i++ {
		for i := 0; i < 1e6; i++ {
		}
	}
}
func longfun2() {
	for i := 1; i < 2e3; i++ {
		for i := 0; i < 1e6; i++ {
		}
	}
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
func hello() {
	fmt.Printf("hello world1!\n")
	start := time.Now()
	sleep()
	longfun1()
	fmt.Println("longfun1 time:", time.Since(start))
	sleep()
	longfun2()
	fmt.Printf("hello world2!\n")
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

	// ... rest of the program ...
	hello()

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
