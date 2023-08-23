package pkg

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Store struct {
	Age int
}

var store *Store
var count = 0

func GetStore() *Store {
	// 注意，加锁
	if store != nil {
		return store
	}
	// 小心写成了 store := &Store{}
	store = &Store{}
	count++
	fmt.Printf("GetStore:%v,create count:%d\n", *store, count)
	return store
}

func TestMultiSafe(t *testing.T) {
	randIntn := func(n int) int {
		s := rand.NewSource(time.Now().UnixNano())
		return rand.New(s).Intn(n)
	}
	count := 5
	fn := func() {
		n := randIntn(5)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Println(n, "s elasped!")
		GetStore()

	}
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
	fmt.Println(count, "tasks done")
}
