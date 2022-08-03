package main

import (
    "github.com/go-redis/redis/v7"
    "fmt"
    "strings"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Test hget

    fields:= strings.Split("foo,resource,bar,baz",",")
	_, err = client.HDel("k3", fields...).Result()
	_, err = client.HSet("k3", "resource", 3).Result()
	val2, err = client.HGet("k3", "resource").Result()
	if err == redis.Nil {
		fmt.Println("k3 does not exist")
		fmt.Printf("%#v", err)
	} else if err != nil {
		panic(err)
	} else {
        fmt.Printf("key3=%#v", val2)
	}

}
