package mymap

type IMap[Key, Value any] interface {
	// Load 从kv中读取一个Key,获取它的Value以及是否存在
	Load(key Key) (value Value, ok bool)

	// Store 向kv中存储一个Key-Value
	Store(key Key, value Value)

	// Delete 在kv中删除一个Key-Value
	Delete(key Key)

	// Range 遍历kv-map,对每一个Key-Value,均调用f,当f返回false的时候,停止遍历
	Range(f func(key Key, value Value) bool)
}
