package main
import (
    "runtime/pprof"
    "runtime"
    "flag"
    "os"
    "log"
    "fmt"
)
 
var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

func longfun(){
    for i:=0;i<1e6;i++{
    }
}
func hello(){
    for i:=0;i<1e4;i++{
        longfun()
    }
    fmt.Printf("hello world!\n")
}
 
func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
 
    // ... rest of the program ...
    hello()
 
    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
}
