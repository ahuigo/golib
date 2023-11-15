package cache

import (
	"errors"
	"fmt"
	"reflect"
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

func (c *cachedFn[int, V]) Get0() (V, error) {
	var s int
	// s = 0                                    // error: cannot use 0 (untyped int constant) as uint8 value in assignment
	fmt.Printf("cache key: %#v, %T\n", s, s) // cache key: 0, uint8
	return c.Get(s)
}

func (c *cachedFn[K, V]) SetTimeout(timeout time.Duration) *cachedFn[K, V] {
	c.mu.Lock()
	c.timeout = timeout
	c.mu.Unlock()
	return c
}

func (c *cachedFn[K, V]) Get(key K) (V, error) {
	var pkey any = key
	needRefresh := false
	kind := reflect.TypeOf(key).Kind()
	if kind == reflect.Map || kind == reflect.Slice {
		pkey = fmt.Sprintf("%#v", key)
	}
	value, hasCache := c.cacheMap.Load(pkey)
	if hasCache {
		cachedObj := value.(*cachedObjType)
		if c.timeout > 0 && time.Since(cachedObj.createdAt) > c.timeout {
			needRefresh = true
		}
	}
	if !hasCache || needRefresh {
		var tmpOnce sync.Once
		oncePtr := &tmpOnce
		//1. clean up routineOnceMap key
		if needRefresh {
			c.routineOnceMap.Delete(pkey)
		}
		// 2. load or store routineOnceMap key
		onceInterface, loaded := c.routineOnceMap.LoadOrStore(pkey, oncePtr)
		if loaded {
			oncePtr = onceInterface.(*sync.Once)
		}
		// 3. Execute getFunc(only once)
		oncePtr.Do(func() {
			val, err := c.getFunc(key)
			createdAt := time.Now()
			c.cacheMap.Store(pkey, &cachedObjType{val: &val, err: err, createdAt: createdAt})
		})
		value, _ = c.cacheMap.Load(pkey)
	}
	cachedObj := value.(*cachedObjType)
	return *(cachedObj.val).(*V), cachedObj.err
}

func TestCacheFuncWithNoParam(t *testing.T) {
	type UserInfo struct {
		Name string
		Age  int
	}

	executeCount := 0
	// Original function
	getUserInfoFromDb := func() (UserInfo, error) {
		executeCount++
		fmt.Println("select * from db limit 1", time.Now())
		time.Sleep(10 * time.Millisecond)
		return UserInfo{Name: "Anonymous", Age: 9}, errors.New("db error")
	}

	// Cacheable Function
	getUserInfoFromDbWithCache := NewCacheFn0(getUserInfoFromDb).SetTimeout(500 * time.Millisecond).Get0 // getFunc can only accept zero parameter
	_ = getUserInfoFromDbWithCache

	// Parallel invocation of multiple functions.
	parallelCall(func() {
		userinfo, err := getUserInfoFromDbWithCache()
		fmt.Println(userinfo, err)
	}, 10)

	// Test timeout
	_, _ = getUserInfoFromDbWithCache()
	time.Sleep(600 * time.Millisecond)
	_, _ = getUserInfoFromDbWithCache()

	if executeCount != 2 {
		t.Error("executeCount should be 2", ", but get ", executeCount)
	}
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

/*
func (c *cachedFn[string, V]) Get0() (V, error) {
	// var s any
	var s string
	// s = "abc" // error: cannot use "abc" (untyped string constant) as string value in assignment
	fmt.Printf("cache key: %#v, %T\n", s, s)
	return c.Get(s)
}
*/

/*
func (c *cachedFn[int, V]) Get0() (V, error) {
	var s int = 100 //error: cannot use 100 (untyped int constant) as int value in variable declaration
	fmt.Printf("cache key: %#v, %T\n", s, s)
	return c.Get(s)
}
*/

func TestCacheFuncWithOneParam(t *testing.T) {
	// Original function
	executeCount := 0
	getUserScore := func(arg map[int]int) (int, error) {
		executeCount++
		fmt.Println("select score from db where id=", arg[0], time.Now())
		time.Sleep(10 * time.Millisecond)
		return 98, errors.New("db error")
	}

	// Cacheable Function
	getUserScoreFromDbWithCache := NewCacheFn(getUserScore).SetTimeout(time.Hour).Get // getFunc can only accept 1 parameter

	// Parallel invocation of multiple functions.
	parallelCall(func() {
		score, err := getUserScoreFromDbWithCache(map[int]int{0: 1})
		fmt.Println(score, err)
		score, err = getUserScoreFromDbWithCache(map[int]int{0: 2})
		fmt.Println(score, err)
		getUserScoreFromDbWithCache(map[int]int{0: 3})
	}, 10)

	if executeCount != 3 {
		t.Error("executeCount should be 3")
	}

}
