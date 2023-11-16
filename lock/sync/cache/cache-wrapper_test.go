package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type cachedObjType struct {
	val       interface{}
	createdAt time.Time
	err       error
}
type cachedFn[Ctx any, K any, V any] struct {
	mu             sync.RWMutex
	cacheMap       sync.Map
	routineOnceMap sync.Map
	timeout        time.Duration
	keyLen         int
	getFunc        func(Ctx, K) (V, error)
}

// Cache Function with ctx and 1 parameter
func NewCacheFn2[Ctx any, K any, V any](getFunc func(Ctx, K) (V, error)) *cachedFn[Ctx, K, V] {
	return &cachedFn[Ctx, K, V]{getFunc: getFunc, keyLen: 2}
}

// Cache Function with 1 parameter
func NewCacheFn1[K any, V any](getFunc func(K) (V, error)) *cachedFn[context.Context, K, V] {
	getFunc0 := func(ctx context.Context, key K) (V, error) {
		return getFunc(key)
	}
	return &cachedFn[context.Context, K, V]{getFunc: getFunc0, keyLen: 1}
}

// Cache Function with no parameter
func NewCacheFn0[V any](getFunc func() (V, error)) *cachedFn[context.Context, int8, V] {
	getFunc0 := func(ctx context.Context, i int8) (V, error) {
		return getFunc()
	}
	return &cachedFn[context.Context, int8, V]{getFunc: getFunc0, keyLen: 0}
}

// Invoke cached function with 1 parameter
func (c *cachedFn[Ctx, K, V]) Get1(key K) (V, error) {
	var ctx Ctx
	return c.Get2(ctx, key)
}

// Invoke cached function with no parameter
func (c *cachedFn[any, int, V]) Get0() (V, error) {
	var ctx any
	var key int
	// key = 0                                    // error: cannot use 0 (untyped int constant) as uint8 value in assignment
	fmt.Printf("cache key: %#v, %T\n", key, key) // cache key: 0, uint8
	return c.Get2(ctx, key)
}

func (c *cachedFn[Ctx, K, V]) SetTimeout(timeout time.Duration) *cachedFn[Ctx, K, V] {
	c.mu.Lock()
	c.timeout = timeout
	c.mu.Unlock()
	return c
}

// Invoke cached function with no parameter
func (c *cachedFn[Ctx, K, V]) Get2(key1 Ctx, key2 K) (V, error) {
	// pkey
	var pkey any = key2
	if _, hasCtx := any(key1).(context.Context); hasCtx || c.keyLen <= 1 {
		// ignore context key
		kind := reflect.TypeOf(key2).Kind()
		if kind == reflect.Map || kind == reflect.Slice {
			pkey = fmt.Sprintf("%#v", key2)
		}
	} else {
		pkey = fmt.Sprintf("%#v,%#v", key1, key2)
	}

	// check cache
	needRefresh := false
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
			val, err := c.getFunc(key1, key2)
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
	getUserScore := func(c *gin.Context, arg map[int]int) (int, error) {
		executeCount++
		fmt.Println("select score from db where id=", arg[0], time.Now())
		time.Sleep(10 * time.Millisecond)
		return 98, errors.New("db error")
	}

	// Cacheable Function
	getUserScoreFromDbWithCache := NewCacheFn2(getUserScore).SetTimeout(time.Hour).Get2 // getFunc can only accept 1 parameter

	// Parallel invocation of multiple functions.
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	parallelCall(func() {
		score, err := getUserScoreFromDbWithCache(ctx, map[int]int{0: 1})
		fmt.Println(score, err)
		score, err = getUserScoreFromDbWithCache(ctx, map[int]int{0: 2})
		fmt.Println(score, err)
		getUserScoreFromDbWithCache(ctx, map[int]int{0: 3})
	}, 10)

	if executeCount != 3 {
		t.Error("executeCount should be 3")
	}

}

func TestCacheFuncWithNilContext(t *testing.T) {
	getUserScore := func(c context.Context, arg map[int]int) (int, error) {
		return 98, errors.New("db error")
	}
	getUserScoreFromDbWithCache := NewCacheFn2(getUserScore).SetTimeout(time.Hour).Get2 // getFunc can only accept 1 parameter
	var ctx context.Context
	getUserScoreFromDbWithCache(ctx, map[int]int{0: 1})
}
