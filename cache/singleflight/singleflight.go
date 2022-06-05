package singleflight

import "sync"

type call struct {
	val interface{}
	err error
	wg  sync.WaitGroup
}

type Group struct {
	mu sync.Mutex // protects m
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock() // 保护 Group 的成员变量 m 不被并发读写
	if g.m == nil {
		g.m = make(map[string]*call) // 延迟初始化
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
