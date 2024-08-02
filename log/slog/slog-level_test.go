package d

import (
	"log/slog"
	"os"
	"testing"
)

func TestLevel(t *testing.T) {
	handlerOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))
	logger2 := logger.With("url", "http://example.com/a/b/c")
	slog.SetDefault(logger2)
	slog.Info("hello", "count", 3)
	// time=2024-08-02T21:29:33.889+08:00 level=INFO source=/slog-level_test.go:17 msg=hello url=http://example.com/a/b/c count=3

}
