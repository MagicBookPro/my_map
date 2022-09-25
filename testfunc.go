package mymap

import (
	"sync"
	"testing"
)

// MapBenchmarkTestFunc 各种IMap的基准测试函数体
func MapBenchmarkTestFunc(m IMap[int, int], goroutines, operations int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(goroutines * operations)
		for j := 0; j < goroutines; j++ {
			go func(index int) {
				for k := 0; k < operations; k++ {
					m.Store(index*operations+k, k)
					m.Load(index*operations + k)
					wg.Done()
				}
			}(j)
		}
		wg.Wait()
	}
}
