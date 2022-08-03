package gotest
import "testing"
import "time"
import "fmt"
// go test -v -bench=. ./bench_parallel_test.go
func add1(a,b int) int {
    time.Sleep(500*time.Millisecond)
    //time.Sleep(time.Second)
    return a+b
}
func add2(a,b int) int {
    var s = 0
    for i := 0; i < 1e9; i++ {
        s +=1
    }
    return s
}
func calc(i,j int) {
    fmt.Printf("======start:%d,%d\n",i,j)
    add1(i,j)
    //add2(i,j)
    fmt.Printf("======end:%d,%d\n",i,j)
}


func Benchmark_Parralel_ahui(b *testing.B) {
    var gi = 0
      // RunParallel will create GOMAXPROCS goroutines
    // and distribute work among them.
    b.RunParallel(func(pb *testing.PB) {
        gi++
        i:=gi // GOMAXPROCS goroutines(num of cpu)
        j:=0
        fmt.Println("start-----------:",gi,i)
        for pb.Next() {
            j++
            fmt.Printf("middle:gi=%d,i=%d,j=%d\n",gi,i,j) // gi>=i(i是闭包)
            calc(i,j)
        }
        fmt.Printf("end:gi=%d,i=%d,j=%d\n",gi,i,j)
    })
}
