package rds

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisClient redis.UniversalClient

func RDB() redis.UniversalClient {
	// var client = redis.NewClient(&redis.Options{
	// 		Addr:     "redis:6379",
	// 		Password: "", // no password set
	// 		DB:       0,  // use default DB
	// 	})
	if redisClient != nil {
		return redisClient
	}
	redisAddr := "redis:6379"
	if redisAddr != "" {
		redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{redisAddr},
			DB:    0,
		})
	} else {
		redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
			MasterName: viper.GetString("redisConfig.masterName"),
			Addrs:      viper.GetStringSlice("redisConfig.addrs"),
			Password:   viper.GetString("redisConfig.password"),
			DB:         0,
		})
	}
	return redisClient
}
