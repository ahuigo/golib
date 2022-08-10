package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	foo := func(ctx context.Context) {
		time.Sleep(0 * time.Second)
		cancel()
	}
	go foo(ctx)
	// <-ctx.Done()
	<-ctx.Done()
	fmt.Println("done")

	// Output:
	// done
}
