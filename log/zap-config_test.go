package main

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000000Z0700"))
}

func TestZapConfig(t *testing.T) {
	// 设置编码
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "root",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		// EncodeTime:    zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeTime: logTimeEncoder, // ISO8601 UTC 时间格式
		// EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := zap.Config{
		Level:            atom,                               // 日志级别
		Development:      false,                              // 开发模式，堆栈跟踪
		Encoding:         "console",                          // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                      // 编码器配置
		InitialFields:    map[string]interface{}{"uid": 9},   // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", "./stdout.log"}, // []string{"stdout", "./a.log"}
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	// logger = func(options ...zap.Option) *zap.Logger {
	// 	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, zap.DebugLevel)
	// 	return zap.New(core).WithOptions(options...)
	// }()

	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	// logger.log(zap.DebugLevel, "msg", nil, map[string]interface{}{"err": "error\tline"})
	logger = logger.Named("project")
	logger.Sugar().Error("[MONITOR_LOG_TYPE=DATAERROR] 数据错误", "data")
	logger.Sugar().Debugw("msg", "k1", map[string]interface{}{"err": "error\tline"}, "k2", map[string]interface{}{"name": "ahuigo"})
	logger.Sugar().Debug(map[string]interface{}{"debug": "error\tline"})
	logger.Sugar().Debugf("debugf: k1=%v, k2=%v", "k1", map[string]interface{}{"err": "error\tline"})
	// logger.Info("无法获取网址",
	// 	zap.String("url", "http://www.baidu.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
}
