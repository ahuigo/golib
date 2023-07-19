package lock

import (
	"fmt"
	"sync"
	"testing"
    "time"
)

/*
sync.Once 作用:
1.　缓存
2.　并发阻塞: refer to /lock/sync/once
*/
var out = 0;
func TestSyncOnce(t *testing.T) {
	var once sync.Once
	onceBody := func() {
        time.Sleep(5*time.Second)
		fmt.Println("Only once")
        out = 100
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(j int) {
            fmt.Println("end0:", j)
			once.Do(onceBody)
            fmt.Println("end:", j, ", out:", out)
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
