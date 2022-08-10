package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context, cancel context.CancelFunc) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("context done!n=", n, ctx.Err())
                    close(dst)
					return
				case dst <- n:
					fmt.Println("wait....,n=", n)
					time.Sleep(1 * time.Second)
					fmt.Println("wait after")
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    t := 0
    if t==1{
        d := time.Now().Add(2 * time.Second)
        ctx, cancel = context.WithDeadline(context.Background(), d)
    }
	defer func() {
		println("call cancel")
		cancel()
	}()

	iter := gen(ctx, cancel)
	for n := range iter {
		fmt.Println(n)
		if n == 4 {
			break
		}
	}
}
