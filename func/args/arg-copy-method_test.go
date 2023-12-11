package main

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
}

// 即使u=*User, 也会copy
func (u User) NotChanged() {
	u.Name = "changed"
}

// 即使u=User{}, 也会传引用
func (u *User) Changed() {
	u.Name = "changed"
}

func TestArgMethod(t *testing.T) {
	// not changed
	u1 := &User{"original"}
	u1.NotChanged()
	fmt.Println(u1)

	// not changed
	u2 := User{"ahui"}
	u2.Changed()
	fmt.Println(u2)

}
