package fxdemo

import (
	"log"
	"os"
	"testing"

	"go.uber.org/dig"
)

// dry-run
// go test -v %
func TestDryRun1(t *testing.T) {
	var err error
	// Dry Run
	c := dig.New()
	// c = dig.New(dig.DryRun(true))

	type Config struct {
		Prefix string
	}

	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}
	c.Provide(func() (*Config, error) {
		return &Config{Prefix: "[foo] "}, nil
	})
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked") // not run if dryRun
	})
	if err != nil {
		panic(err)
	}
}
