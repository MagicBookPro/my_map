package mymap

import "sync"

type Table[Key, Value any] struct {
	mu    *sync.RWMutex
	lines map[any]Value
}

func (t *Table[Key, Value]) set(key Key, value Value) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.lines[key] = value
}

func (t *Table[Key, Value]) get(key Key) (value Value, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	got, gotten := t.lines[key]
	return got, gotten
}

func (t *Table[Key, Value]) del(key Key) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.lines, key)
}

func (t *Table[Key, Value]) tRange(f func(Key, Value) bool) bool {
	for k, v := range t.lines {
		if !f(k.(Key), v) {
			return false
		}
	}

	return true
}

func NewRWMap[Key, Value any](length uint) IMap[Key, Value] {
	rwMap := RWMap[Key, Value]{
		tables: map[uint]*Table[Key, Value]{},
		length: length,
	}

	for i := uint(0); i < length; i++ {
		rwMap.tables[i] = &Table[Key, Value]{
			mu:    &sync.RWMutex{},
			lines: map[any]Value{},
		}
	}

	return &rwMap
}
