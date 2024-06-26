package cmd

import (
	"fmt"
	"rds"
	"testing"
	"time"
)

func TestTTL(t *testing.T) {
	client := rds.RDB()
	// test get error
	val, err := client.Get("ahuikey1").Result()
	if err != nil {
		fmt.Printf("get err:%v, val:%v\n", err.Error(), val)
	}

	// test ttl
	err = client.Set("key", "value", time.Duration(1000*1000)).Err()
	if err != nil {
		println("set err:", err.Error())
	}

	ttl, err := client.TTL("key").Result() //-1 if no ttl, -2 if expired
	if err != nil {
		panic(err)
	}
	fmt.Printf("ttl:%#v, (ttl==-2):%v\n", ttl, ttl == -2)
}
