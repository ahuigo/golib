package main

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

type User struct {
	Username string
	Age      int
}

func run(name string, logger *zap.Logger) {
	println(name)
	defer logger.Sync()
	defer println("")

	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	loggerS := logger.Named("日志名").Sugar()
	loggerS.Debugw("Sugar", "key1", map[string]interface{}{"k": 1}, "key2:", User{})
	loggerS.Error("some fatal")
}

func TestZapDevProd(t *testing.T) {
	// format NewProduction json序列化输出; NewDevelopment: 普通Info format
	logger, _ := zap.NewDevelopment()
	run("dev:", logger)
	logger, _ = zap.NewProduction()
	run("prod:", logger)
	logger = zap.NewExample()
	run("example:", logger)
}
