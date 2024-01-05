// https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
package lock

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type User struct {
}
type Key struct {
}

func TestMapNotOk(t *testing.T) {
	// 1. test map
	var sm sync.Map
	key := &Key{}

	// 2.1 add 100 by reference
	v, _ := sm.Load(key)
	fmt.Printf("%#v\n", v == nil)
}

func TestMapUpdate(t *testing.T) {
	// 1. test map
	m := map[string]int{
		"age": 1,
	}
	mt := m
	mt["age"] = 2
	fmt.Printf("ori m: %#v\n", m)

	// 2. store in sync.Map
	var sm sync.Map
	key := &Key{}
	sm.Store(key, m)

	// 2.1 add 100 by reference
	v, _ := sm.Load(key)
	m2 := v.(map[string]int)
	m2["age"] += 100

	// 2.2 read m
	v, _ = sm.Load(key)
	m3 := v.(map[string]int)
	fmt.Printf("%#v\n", m3)
}

func TestSyncMapRange(t *testing.T) {
	var sm sync.Map
	sm.Store("key1", 1)
	sm.Store("key2", "v2")
	sm.Store("key3", "v3")
	sm.Range(func(k, v interface{}) bool {
		fmt.Println("k:", k.(string))
		return k.(string) != "key2"
	})
}

func TestMapKey(t *testing.T) {
	// 1. key
	key1 := User{}
	key2 := &map[string]int{
		"age": 1,
	}

	// 2. store in sync.Map
	var sm sync.Map
	var m = 100

	// 2.2 read
	sm.Store(key1, m)
	_, ok := sm.Load(key1) // key1 参数不是指针，所以不会相等
	if !ok {
		t.Fatalf("key1(struct) should exist")
	}

	sm.Store(key2, m)
	if _, ok := sm.Load(key2); !ok {
		t.Fatalf("key2(*map) should exist")
	}

	// 1. array+struct+pointer 可以作为key(会被自动序列化)
	bk := []byte("bk")
	sm.Store(&bk, m) // ok(所有的指针都可以作为key)

	// 2. map + slice 不可以作为key
	// sm.Store(key2, m) // panic(map[string]int)
	// sm.Store([]byte(""), m) // panic([]byte)
	// sm.Store([]int{}, m) // panic([]int)

}

func TestMapKeyArray(t *testing.T) {
	var sm sync.Map
	key1 := [2]int{1}
	key2 := [2]int{1}
	val1 := 10
	sm.Store(key1, val1)
	val2, ok := sm.Load(key2) // array 会被自动序列化(copy)
	if !ok {
		t.Fatalf("fail load")
	} else if val2 != val1 {
		t.Fatalf("fail load%d,%d", val1, val2)
	}
}

func TestMapKeyNil(t *testing.T) {
	var sm sync.Map

	// key1 key2 都是interface{}类型，相同的key
	var key1 any // interface{} is hashable (nil 指针)
	var key2 context.Context
	// key3 key4 都是指针类型，不同的key
	var key3 *int  // key3作为指针，is hashable
	var key4 []int // slice/map, is not hashable
	_, _ = key3, key4

	fmt.Printf("%#v\n", key1)
	fmt.Printf("%v,%#v\n", key2 == nil, key2)

	val1 := 10
	sm.Store(key1, val1)
	val2, ok := sm.Load(key2)
	if !ok {
		t.Fatalf("fail load")
	} else if val2 != val1 {
		t.Fatalf("fail load%d,%d", val1, val2)
	}
}

func TestPlainMapKeyNil(t *testing.T) {
	// key1 key2 都是interface{}类型，相同的key
	var sm map[any]int
	var key1 any // interface{} is hashable (nil 指针)
	var key2 context.Context
	// key3 key4 都是指针类型，不同的key
	var key3 *int  // key3作为指针，is hashable
	var key4 []int // slice/map, is not hashable
	_, _ = key3, key4

	fmt.Printf("v:%#v,%v,%v,%v\n", key1, key2, key3, key4)
	fmt.Printf("T:%T,%T,%T,%T\n", key1, key2, key3, key4)

	val1 := 10
	val2, ok := sm[key2]
	if !ok {
		t.Fatalf("fail load")
	} else if val2 != val1 {
		t.Fatalf("fail load%d,%d", val1, val2)
	}
}
func TestKeyNilReflect(t *testing.T) {
	var key1 *int
	var key2 any
	rv1 := reflect.ValueOf(key1)
	rt1 := reflect.TypeOf(key1)
	rv2 := reflect.ValueOf(key2)
	rt2 := reflect.TypeOf(key2)
	// output: true, true, true, panic
	fmt.Printf("rv1.kind:%#v\n", rv1.Kind() == reflect.Ptr)
	fmt.Printf("rt1.kind:%#v\n", rt1.Kind() == reflect.Ptr)
	fmt.Printf("rv2.kind:%#v\n", rv2.Kind() == reflect.Invalid)
	fmt.Printf("rt2.kind:%#v\n", rt2.Kind()) // panic, type(any) 是无效的type, 没有kind()
}
