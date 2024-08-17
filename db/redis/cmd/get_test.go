package cmd

import (
	"encoding/json"
	"fmt"
	"rds"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
)

type User struct {
	Name   string
	Height int
	age    int
}

func TestGet(t *testing.T) {
	client := rds.RDB()
	user := User{"ahuigo", 20, 20}

	ustr, _ := json.Marshal(user)

	// set string
	client.Set("key", ustr, 0)
	val, err := client.Get("key").Result()
	fmt.Printf("Get struct val:%#v, err:%v\n", val, err != nil)

	// set expire 1s
	err = client.Set("key", ustr, time.Second).Err()
	if err != nil {
		fmt.Printf("Set err:%#v, err:%v\n", err, err)
	}
	val, err = client.Get("key").Result()
	fmt.Printf("Get nil val:%#v, err:%v\n", val, err)

	// set expire 0
	err = client.Set("key", 1, 0).Err()
	if err != nil {
		fmt.Printf("Set err:%#v, err:%v\n", err, err)
	}

	// 4. get value
	age, err := client.Get("key").Int64()
	if err == redis.Nil {
		t.Fatal("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		t.Logf("Get val:%#v, err:%v\n", age, err)
	}

}
