package main

import "fmt"
import "time"
import "math/rand"

type User struct {
	Name string
}

func main() {
	user := User{"ahui"}
	fmt.Println(user) //ahui
	rand.Seed(time.Now().Unix())

	for {
		err := getUsers()
		fmt.Println("noerr:", err)
		time.Sleep(time.Second * 3)
	}
	fmt.Println("som err:")

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
