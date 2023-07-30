package cache

import (
	"sync"
	"testing"
)

type CacheFn[K comparable, V any] struct {
	redisMap       sync.Map
	routineOnceMap sync.Map
	getFunc        func(K) V
}

func NewCacheFn[K comparable, V any](getFunc func(K) V) *CacheFn[K, V] {
	return &CacheFn[K, V]{getFunc: getFunc}
}

// 1. 缓存到 redisMap 或其它存储; 2.如果多个协程同时执行时，只执行一次（其它协程被阻塞）
func (c *CacheFn[K, V]) Get(key K) V {
	value, ok := c.redisMap.Load(key)
	if ok {
		return value.(V)
	} else {
		var once sync.Once
		onceInterface, loaded := c.routineOnceMap.LoadOrStore(key, &once)
		if loaded { // 如果有其它协程在执行，则等待它结束
			oncePtr := onceInterface.(*sync.Once)
			oncePtr.Do(func() {})
		} else { // 第一次访问，进行DB查询
			once.Do(func() {
				value = c.getFunc(key)
				c.redisMap.Store(key, value)
			})
		}
		val, _ := c.redisMap.Load(key)
		return val.(V)
	}
}

func TestCacheFuncWrapperGeneric(t *testing.T) {
	type UserInfo struct {
		Name string
		Age  int
	}

	// 原始函数
	getUserInfoFromDb := func(name string) UserInfo {
		println("get info from db:", name)
		return UserInfo{Name: name}
	}

	// 带缓存的函数
	getUserInfoFromDbWithCache := NewCacheFn(getUserInfoFromDb) // getFunc 只接受一个参数，怎么接收多个参数呢？

	// 多个协程同时执行
	batchCall := func(t *testing.T, fn func()) {
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

	// 多次调用函数, 只执行一次
	batchCall(t, func() {
		userinfo := getUserInfoFromDbWithCache.Get("alex")
		t.Log(userinfo)
	})
}
