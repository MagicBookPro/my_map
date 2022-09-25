package mymap

import (
	"testing"
)

// BenchmarkCommonMapTest 测试CommonMap进行100万次读和写的基准性能
func BenchmarkCommonMapTest(b *testing.B) {
	m := NewCommonMap[int, int]()
	MapBenchmarkTestFunc(m, 100, 10000, b)
}

// BenchmarkSimpleMapTest 测试SimpleMap进行100万次读和写的基准性能
func BenchmarkSimpleMapTest(b *testing.B) {
	m := NewSimpleMap[int, int]()
	MapBenchmarkTestFunc(m, 100, 10000, b)
}

// BenchmarkRWMapTest 测试RWMap进行100万次读和写的基准性能
func BenchmarkRWMapTest(b *testing.B) {
	m := NewRWMap[int, int](10007)
	MapBenchmarkTestFunc(m, 100, 10000, b)
}
