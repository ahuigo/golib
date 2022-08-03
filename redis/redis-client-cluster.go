	redisAddr := viper.GetString("redisConfig.addr")
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
	service.RedisClient = redisClient
