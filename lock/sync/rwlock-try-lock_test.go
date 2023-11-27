package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/petermattis/goid"
)

func TestRwlockWithTryLock(t *testing.T) {
	var wg sync.WaitGroup
	for k := 0; k < 2; k++ {
		wg.Add(1)
		go func() {
			runWithRWlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

var tmpOnce sync.RWMutex
var (
	tryLockCount = 0
	pkeyLock     = &tmpOnce
)

func runWithRWlock() (retv int, err error) {
checkCache:
	tryLockCount++
	pkeyLock.RLock()
	time.Sleep(10 * time.Millisecond)
	pkeyLock.RUnlock()

	isLocked := pkeyLock.TryLock() // why do all the goroutines lock return false?????

	if !isLocked {
		pkeyLock.Lock()
		if tryLockCount < 3 {
			pkeyLock.Unlock()
			goto checkCache
		}
		fmt.Printf("failedTry: gid=%d,count=%d,time=%v\n", goid.Get(), tryLockCount, time.Now())
	} else {
		fmt.Printf("successTry: gid=%d,count=%d,time=%v\n", goid.Get(), tryLockCount, time.Now())
	}
	defer pkeyLock.Unlock()

	return
}
