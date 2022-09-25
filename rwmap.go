package mymap

type RWMap[Key, Value any] struct {
	tables map[uint]*Table[Key, Value]
	length uint
}

func (R *RWMap[Key, Value]) Load(key Key) (value Value, ok bool) {
	index := hash(key) % R.length
	return R.tables[index].get(key)
}

func (R *RWMap[Key, Value]) Store(key Key, value Value) {
	index := hash(key) % R.length
	R.tables[index].set(key, value)
}

func (R *RWMap[Key, Value]) Delete(key Key) {
	index := hash(key) % R.length
	R.tables[index].del(key)
}

func (R *RWMap[Key, Value]) Range(f func(key Key, value Value) bool) {
	for _, t := range R.tables {
		if !t.tRange(f) {
			break
		}
	}
}
