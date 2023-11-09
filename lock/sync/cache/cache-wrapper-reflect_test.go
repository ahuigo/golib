package cache

import (
	"reflect"
	"sync"
	"testing"
)

type Function interface{}

func CacheFunc(f Function) func(...interface{}) interface{} {
	var (
		onceMap    sync.Map // 用于多协程阻塞
		redisCache sync.Map
	)

	fn := reflect.ValueOf(f)
	return func(args ...interface{}) interface{} {
		var key string
		for i := 0; i < len(args); i++ {
			key += reflect.ValueOf(args[i]).String()
		}

		value, ok := redisCache.Load(key)
		if ok {
			return value
		}

		in := make([]reflect.Value, len(args))
		for i := range args {
			in[i] = reflect.ValueOf(args[i])
		}

		var once sync.Once
		onceInterface, loaded := onceMap.LoadOrStore(key, &once)
		if loaded { // 如果已经存在，则等待
			oncePtr := onceInterface.(*sync.Once)
			oncePtr.Do(func() {})
		} else { // 第一次访问，进行DB查询
			once.Do(func() {
				res := fn.Call(in)
				redisCache.Store(key, res[0].Interface())
			})
		}
		val, _ := redisCache.Load(key)
		return val
	}
}

// 缓存redis或其它; 并且多个协程同时执行时，只执行一次（其它协程被阻塞）
func TestRedisCacheGroutineWrapperReflect(t *testing.T) {
	type UserInfo struct {
		Name string
		Age  int
	}
	getUserInfoFromDb := func(name string) UserInfo {
		println("call db")
		return UserInfo{Name: name}
	}

	// getUserInfoFromDbWithCache := Cached(getUserInfoFromDb).(func(string) UserInfo)
	getUserInfoFromDbWithCache := CacheFunc(getUserInfoFromDb)

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

	batchCall(t, func() {
		userinfo := getUserInfoFromDbWithCache("alex")
		t.Log(userinfo)
	})
}
