package mymap

type IMap[Key, Value any] interface {
	Load(key Key) (value Value, ok bool)
	Strore(key Key, value Value)
	Delete(key Key)
}
