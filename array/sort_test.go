package main

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	age int
}

// case: sort struct of slice (inplace)
func sortStructSlice() {
	users := []User{
		User{age: 5},
		User{age: 1},
		User{age: 11},
		User{age: 0},
	}
	// 回调函数可以直接访问user.age
	sort.Slice(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})
	fmt.Printf("users: %v\n", users)
}

func sortBase() {
	// 支持IntSlice, StringSlice...等基础类型
	sl := []string{"mumbai", "london", "tokyo", "seattle"}
	sort.Sort(sort.StringSlice(sl))
	fmt.Println(sl)

	intSlice := []int{3, 5, 6, 4, 2, 293, -34}
	sort.Sort(sort.IntSlice(intSlice))
	fmt.Println(intSlice)
}

func TestSort(t *testing.T) {
	sortStructSlice()
}
