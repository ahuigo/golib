package main

import (
    "github.com/go-redis/redis/v7"
    "fmt"
    "time"
        "encoding/json"

)

type User struct{
    Name string 
    Height int 
    age int
}


var client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

func main() {
    user := User{"ahuigo", 20, 20}

    ustr,_ := json.Marshal(user)


    // set string
	client.Set("key", ustr, 0)
    val, err := client.Get("key").Result()
    fmt.Printf("Get struct val:%#v, err:%v\n", val, err!=nil)

    // set expire
	err = client.Set("key", ustr, time.Duration(1000*1000)).Err()
    if err!=nil{
        fmt.Printf("Set err:%#v, err:%v\n", err, err)
    }
	val,err = client.Get("key").Result()
    fmt.Printf("Get nil val:%#v, err:%v\n", val, err)


    // set expire
	err = client.Set("key", 1, 0).Err()
    if err!=nil{
        fmt.Printf("Set err:%#v, err:%v\n", err, err)
    }
	val,err = client.Get("key").Result()
    fmt.Printf("Get nil val:%#v, err:%v\n", val, err)

}


