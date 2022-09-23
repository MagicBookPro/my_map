package mymap

import "testing"

func BenchmarkCommonMapTestFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonMapTestFunc()
	}
}
