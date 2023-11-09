// https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
package lock

import (
	"fmt"
	"sync"
	"testing"
)

type User struct {
}
type Key struct {
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
