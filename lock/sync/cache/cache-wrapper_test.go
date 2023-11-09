package cache

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

type cachedObjType struct {
	val       interface{}
	createdAt time.Time
	err       error
}
type cachedFn[K any, V any] struct {
	mu             sync.RWMutex
	cacheMap       sync.Map
	routineOnceMap sync.Map
	timeout        time.Duration
	getFunc        func(K) (V, error)
}

func NewCacheFn[K any, V any](getFunc func(K) (V, error)) *cachedFn[K, V] {
	return &cachedFn[K, V]{getFunc: getFunc}
}

func NewCacheFn0[V any](getFunc func() (V, error)) *cachedFn[int8, V] {
	getFunc0 := func(i int8) (V, error) {
		return getFunc()
	}
	return &cachedFn[int8, V]{getFunc: getFunc0}
}

func (c *cachedFn[string, V]) Get0() (V, error) {
	// var s any
	var s string
	// s = "abc" // error: cannot use "abc" (untyped string constant) as string value in assignment
	fmt.Printf("cache key: %#v, %T\n", s, s)
	return c.Get(s)
}

/*
func (c *cachedFn[int, V]) Get0() (V, error) {
	var s int = 100 //error: cannot use 100 (untyped int constant) as int value in variable declaration
	fmt.Printf("cache key: %#v, %T\n", s, s)
	return c.Get(s)
}
*/

func (c *cachedFn[K, V]) SetTimeout(timeout time.Duration) *cachedFn[K, V] {
	c.mu.Lock()
	c.timeout = timeout
	c.mu.Unlock()
	return c
}

func (c *cachedFn[K, V]) Get(key K) (V, error) {
	needRefresh := false
	value, ok := c.cacheMap.Load(key)
	if ok {
		cachedObj := value.(*cachedObjType)
		if c.timeout > 0 && time.Since(cachedObj.createdAt) > c.timeout {
			needRefresh = true
		}
	}
	if !ok || needRefresh {
		var tmpOnce sync.Once
		oncePtr := &tmpOnce
		onceInterface, loaded := c.routineOnceMap.LoadOrStore(key, oncePtr)
		if loaded {
			oncePtr = onceInterface.(*sync.Once)
		}
		oncePtr.Do(func() {
			val, err := c.getFunc(key)
			c.cacheMap.Store(key, &cachedObjType{val: &val, err: err})
		})
		value, _ = c.cacheMap.Load(key)
	}
	cachedObj := value.(*cachedObjType)
	return *(cachedObj.val).(*V), cachedObj.err
}

func TestCacheFuncWithNoParam(t *testing.T) {
	type UserInfo struct {
		Name string
		Age  int
	}

	// Original function
	getUserInfoFromDb := func() (UserInfo, error) {
		fmt.Println("select * from db limit 1", time.Now())
		time.Sleep(10 * time.Millisecond)
		return UserInfo{Name: "Anonymous", Age: 9}, errors.New("db error")
	}

	// Cacheable Function
	getUserInfoFromDbWithCache := NewCacheFn0(getUserInfoFromDb) // getFunc can only accept zero parameter
	_ = getUserInfoFromDbWithCache

	// Parallel invocation of multiple functions.
	parallelCall(func() {
		userinfo, err := getUserInfoFromDbWithCache.Get0()
		fmt.Println(userinfo, err)
	}, 10)
}

// Parallel caller via goroutines
func parallelCall(fn func(), times int) {
	var wg sync.WaitGroup
	for k := 0; k < times; k++ {
		wg.Add(1)
		go func() {
			fn()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCacheFuncWithOneParam(t *testing.T) {
	type UserInfo struct {
		Name string
		Age  int
	}

	// Original function
	getUserInfoFromDb := func(name string) (UserInfo, error) {
		fmt.Println("select * from db where name=", name, time.Now())
		time.Sleep(10 * time.Millisecond)
		return UserInfo{Name: name, Age: 9}, errors.New("db error")
	}

	// Cacheable Function
	getUserInfoFromDbWithCache := NewCacheFn(getUserInfoFromDb) // getFunc can only accept 1 parameter

	// Parallel invocation of multiple functions.
	parallelCall(func() {
		userinfo, err := getUserInfoFromDbWithCache.Get("alex")
		fmt.Println(userinfo, err)
		userinfo, err = getUserInfoFromDbWithCache.Get("John")
		fmt.Println(userinfo, err)
	}, 10)

}
