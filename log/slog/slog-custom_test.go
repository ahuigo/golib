package d

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"
)

type mlogger struct {
	*slog.Logger
}

func (l *mlogger) Infof(format string, args ...any) {
	// slog.Default().Info(fmt.Sprintf(format, args...))
	l.Logger.Info(fmt.Sprintf(format, args...))
}
func getLogger(app string, level slog.Level) *mlogger {
	opts := &slog.HandlerOptions{Level: level}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	logger2 := logger.With("app", "gonic")
	return &mlogger{logger2}
}

var logger = getLogger("gonic", slog.LevelInfo)

func TestFormat(t *testing.T) {
	logger.Infof("hello %s", "world")
	// time=2024-01-15T15:39:09.812+08:00 level=INFO msg="hello world" app=gonic

}

func TestGroup(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	// slog.Default().With("id", systemID)
	logger2 := logger.With("app", "gonic")
	slog.Group("request",
		"method", "POST",
		"url", "/login")
	slog.SetDefault(logger2)

	logger.Info("finished",
		slog.Group("req",
			slog.String("method", "POST"),
			slog.String("url", "/login"),
		),
		slog.Int("status", 200),
		slog.Duration("duration", time.Second),
		"count", 3,
	)
	// time=2024-01-15T15:15:18.494+08:00 level=INFO msg=finished req.method=POST req.url=/login status=200 duration=1s count=3
	slog.Info("hello", "count", map[string]int{"a": 1})
	//time=2024-01-15T15:15:18.494+08:00 level=INFO msg=hello app=gonic count=map[a:1]

}
