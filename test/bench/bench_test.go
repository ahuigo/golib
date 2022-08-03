package gotest
import "testing"
import "time"
import "fmt"
// go test -v -bench=. benchmark_test.go
func add1(a,b int) int {
    time.Sleep(1000*1000*1000)
    return a+b
}
func add2(a,b int) int {
    var s = 0
    for i := 0; i < 1e9; i++ {
        s +=1
    }
    return s
}
func calc(a,b int) {
    add1(a,b)
    add2(a,b)
}

func Benchmark_calc(b *testing.B) {
    b.StopTimer() //调用该函数停止压力测试的时间计数
    b.StartTimer() //重新开始时间
    for i := 0; i < b.N; i++ {
        calc(4, 5)
    }
}

func TestHelloWorld(t *testing.T) {
    //t.Errorf("Oh! %q", "some error")
    t.Log("hello world")
}

func Benchmark_Alloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", i)
    }
}
