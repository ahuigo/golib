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
	// time=2023-11-15T14:56:46.901+08:00 level=INFO msg=hello url=http://example.com/a/b/c count=3
}
