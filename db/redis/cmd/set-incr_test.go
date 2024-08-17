package cmd

import (
	"rds"
	"testing"
	"time"
)

// INCRBY mykey 5 不支持XX
/**
SET key value [EX seconds] [PX milliseconds] [NX|XX]
	EX seconds -- Set the specified expire time, in seconds (a positive integer).
	PX milliseconds -- Set the specified expire time, in milliseconds (a positive integer).
	EXAT timestamp-seconds -- Set the specified Unix time at which the key will expire, in seconds (a positive integer).
	PXAT timestamp-milliseconds -- Set the specified Unix time at which the key will expire, in milliseconds (a positive integer).
	NX -- Only set the key if it does not already exist.
	XX -- Only set the key if it already exists.
*/

func TestSetNX(t *testing.T) {
	client := rds.RDB()
	key := "key4"

	// set string
	isSetted, err := client.SetNX(key, 13, 100*time.Second).Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isSetted)

	val, err := client.IncrBy(key, 10).Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}
