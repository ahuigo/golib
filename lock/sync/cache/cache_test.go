package cache

import (
	"sync"
	"testing"
)

// UserInfo 包装你的用户信息对象
type UserInfo struct {
	Name string
	Age  int
}

var (
	routineOnceMap sync.Map
	redisMap       sync.Map
)

func getUserInfoFromDbWithCache(username string) UserInfo {
	getUserInfoFromDb := func(name string) UserInfo {
		println("call db")
		return UserInfo{Name: name}
	}
	value, ok := redisMap.Load(username)
	if ok {
		return value.(UserInfo)
	}

	// 使用 sync.Once 确保只有一个协程执行db操作
	var once sync.Once
	onceInterface, loaded := routineOnceMap.LoadOrStore(username, &once)

	if loaded { // 如果已经存在，则等待
		oncePtr := onceInterface.(*sync.Once)
		oncePtr.Do(func() {})
	} else { // 第一次访问，进行DB查询
		once.Do(func() {
			userInfo := getUserInfoFromDb(username)
			redisMap.Store(username, userInfo)
		})
	}

	val, _ := redisMap.Load(username)
	return val.(UserInfo)
}

// 缓存redis或其它; 并且多个协程同时执行时，只执行一次（其它协程被阻塞）
func TestRedisCacheGroutine(t *testing.T) {
	batchCall(t, func() {
		userinfo := getUserInfoFromDbWithCache("alex")
		t.Log(userinfo)
	})
}

func batchCall(t *testing.T, fn func()) {
	var wg sync.WaitGroup
	for k := 0; k < 10; k++ {
		wg.Add(1)
		go func(i int) {
			fn()
			wg.Done()
		}(k)
	}
	wg.Wait()
}
