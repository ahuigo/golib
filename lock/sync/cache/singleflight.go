package cache

import "sync"

var callGroup *Group
var onceGroup sync.Once

type call struct {
	wg  sync.WaitGroup
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func GetGroupSingleton() *Group {
	onceGroup.Do(func() { callGroup = &Group{mu: sync.Mutex{}} })
	return callGroup
}

// 这个Do 和sync.Once.Do 不一样的是，它可能会执行多次。除非多协程同时访问，只会执行一次（其它协程会被wait阻塞且不执行），是协程安全的, 适合redis
func (g *Group) Do(key string, fn func() error) error {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.err
	}

	c := new(call)

	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.err
}
