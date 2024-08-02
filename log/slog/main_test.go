package d

import (
	"log/slog"
	"os"
	"testing"
)

var Slogger *slog.Logger

func init() {
	Slogger = getSlogger(false)
}
func getSlogger(isJson bool) (logger *slog.Logger) {
	handlerOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	if isJson {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))
	}
	// logger2 := logger.With("url", "http://example.com/a/b/c")
	return logger
}

func TestSlogger(t *testing.T) {
	logger := getSlogger(false)
	logger.Info("hello", "count", 3)
	// time=2024-08-02T21:30:55.812+08:00 level=INFO source=/main_test.go:30 msg=hello count=3

	logger = getSlogger(true)
	logger.Info("hello", "count", 3)
	// {"time":"2024-08-02T21:30:55.812395+08:00","level":"INFO","source":{"function":"go-lib/log/slog.TestSlogger","file":"/main_test.go","line":34},"msg":"hello","count":3}

}
