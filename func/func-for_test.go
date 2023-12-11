package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type User struct {
	Name string
}

func TestFuncFor(t *testing.T) {
	user := User{"ahui"}
	fmt.Println(user) //ahui

	for {
		err := getUsers()
		fmt.Println("noerr:", err)
		time.Sleep(time.Second * 3)
	}

}

func getUsers() (err error) {
	n := rand.Intn(2)
	fmt.Println("getUser1: ", n)
	fmt.Println("getUser: ", n)
	if n > 0 {
		err = fmt.Errorf("some")
	}
	return
}
