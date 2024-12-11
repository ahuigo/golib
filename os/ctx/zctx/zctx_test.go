package zctx
// 演示context: 支持cancelFns + error
import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

type cancelFn func()
type Zcontext struct {
	context.Context
	cancelFns []cancelFn
	mu        sync.Mutex   // protects following fields
	done      atomic.Value // of chan struct{}, created lazily, closed by first cancel call
	err       error        // set to non-nil by the first cancel call
}

var cancelCtxKey int

// ErrCanceled is the error returned by Zcontext.Err when the context is canceled.
var ErrCanceled = errors.New("context canceled")

func WithCancelFns(parent context.Context) (ctx *Zcontext, cancel func()) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := Zcontext{Context: parent}
	return &c, func() { c.cancel(true, ErrCanceled) }
}

func (c *Zcontext) AddCancelFns(cancelFns ...cancelFn) {
	c.cancelFns = append(c.cancelFns, cancelFns...)
}

func (c *Zcontext) Cancel() {
	c.cancel(true, ErrCanceled)
}

func (c *Zcontext) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	d, _ := c.done.Load().(chan struct{})
	if d == nil {
		c.done.Store(closedchan)
	} else {
		close(d)
	}

	for _, cancelFn := range c.cancelFns {
		cancelFn()
	}
	c.cancelFns = nil
	c.mu.Unlock()
}

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

func (c *Zcontext) Value(key interface{}) interface{} {
	if key == &cancelCtxKey {
		return c
	}
	return c.Context.Value(key)
}

func (c *Zcontext) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}

func (c *Zcontext) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func TestCancelFns(t *testing.T) {
	ctx, cancel := WithCancelFns(context.Background())
	// foo.go
	ctx.AddCancelFns(
		func() {
			println("close pg")
		},
		func() {
			println("close server")
		},
	)

	// defer.go
	cancel()
}
