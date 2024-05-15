package d

import (
	"log/slog"
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	slog.Info("hello", "count", 3)
	//2022/11/08 15:28:26 INFO hello count=3

}

func TestConsole(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Info("hello", "count", 3)
	// time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3

}
func TestJson(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)
	// {"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}

}
func TestHandler(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger2 := logger.With("url", "http://example.com/a/b/c")
	slog.SetDefault(logger2)
	slog.Info("hello", "count", 3)
	// time=2023-11-15T14:56:46.901+08:00 level=INFO msg=hello url=http://example.com/a/b/c count=3

}
