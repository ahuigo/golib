package cmd

import (
	"fmt"
	"rds"
	"testing"
	"time"
)

func TestTTL(t *testing.T) {
	client := rds.RDB()
	// test ttl
	err := client.Set("key", "value", 10*time.Second).Err()
	if err != nil {
		println("set err:", err.Error())
	}

	ttl, err := client.TTL("key").Result() //-1 if no ttl, -2 if expired(or not existed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ttl:%#v, (ttl==-2):%v\n", ttl, ttl == -2)
}
