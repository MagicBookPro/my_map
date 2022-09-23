package mymap

import "sync"

type CommonMap[Key, Value any] struct {
	mu *sync.Mutex
	m  map[any]Value
}

func NewCommonMap[Key, Value any]() IMap[Key, Value] {
	return &CommonMap[Key, Value]{
		mu: &sync.Mutex{},
		m:  make(map[any]Value),
	}
}

func (c *CommonMap[Key, Value]) Load(key Key) (value Value, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok = c.m[key]
	return value, ok
}

func (c *CommonMap[Key, Value]) Strore(key Key, value Value) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value

}

func (c *CommonMap[Key, Value]) Delete(key Key) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
	c.mu.Unlock()

}

func CommonMapTestFunc() {
	wg := &sync.WaitGroup{}
	wg.Add(100 * 10000)
	m := NewCommonMap[int, int]()
	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				m.Strore(i, i)
				m.Load(i)
				wg.Done()
			}

		}()
	}
	wg.Wait()
}
