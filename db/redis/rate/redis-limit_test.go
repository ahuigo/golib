package demo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func TestExampleNewLimiter(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// _ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)
	for i := 0; i < 10; i++ {
		res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(4))
		if err != nil {
			panic(err)
		}
		fmt.Println("allowed", res.Allowed, "remaining", res.Remaining, time.Now())
	}

}
