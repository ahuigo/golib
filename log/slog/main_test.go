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
	if isJson {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	// logger2 := logger.With("url", "http://example.com/a/b/c")
	return logger
	// logger.Info("hello", "count", 3)
	// time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3
}
func TestSlogger(t *testing.T) {
	logger := getSlogger(false)
	logger.Info("hello", "count", 3)
	// {"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}
}
