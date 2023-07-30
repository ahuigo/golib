package cache

import (
	"sync"
	"testing"
	"time"
)

// 多个协程同时执行时，只执行一次（其它协程阻塞）; 否则多次
func TestGoroutineOnce(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	getUser := func(name string) User {
		time.Sleep(time.Second)
		return User{Name: name, Age: 18}
	}

	redisCache := map[string]User{}
	flushCacheUser := func(name string) error {
		userinfo := getUser(name)
		redisCache[name] = userinfo
		return nil
	}

	username := "alex"
	var wg sync.WaitGroup
	for k := 0; k < 5; k++ {
		wg.Add(1)
		go func(i int) {
			var userinfo User
			// 这个Do 和sync.Once.Do 不一样的是，它可能会执行多次。除非多协程同时访问，只会执行一次（其它协程会被wait阻塞且不执行），是协程安全的, 适合redis
			err := GetGroupSingleton().Do(username, func() error {
				println(i)
				return flushCacheUser(username)
			})
			_ = userinfo
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(k)
	}
	wg.Wait()

}
