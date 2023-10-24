package cmd

import (
	"fmt"
	"rds"
	"strings"
	"testing"

	"github.com/go-redis/redis/v7"
)

func TestHashSetGet(t *testing.T) {
	client := rds.RDB()
	fields := strings.Split("foo,resource,bar,baz", ",")
	// hkey := "hkey"

	// 1. set
	if _, err := client.HSet("hkey", "foo", 3).Result(); err != nil {
		t.Fatal(err)
	}
	// 1.2 set multi
	if _, err := client.HMSet("hkey", map[string]any{
		"resource": "rds",
		"bar":      "barv",
	}).Result(); err != nil {
		t.Fatal(err)
	}

	// 2. get key
	val, err := client.HGet("hkey", "foo").Result()
	if err == redis.Nil {
		fmt.Println("filed does not exist")
		t.Logf("%#v", err)
	} else if err != nil {
		panic(err)
	} else {
		t.Logf("field=%#v", val)
	}

	// 2.2 get all
	if valMap, err := client.HGetAll("hkey").Result(); err == nil {
		for k, v := range valMap {
			t.Logf("%v=%v\n", k, v)
		}
	}

	// 3. del key list
	if _, err := client.HDel("hkey", fields...).Result(); err != nil {
		t.Fatal(err)
	}

}
