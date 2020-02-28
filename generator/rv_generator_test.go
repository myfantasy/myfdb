package generator

import (
	"testing"
)

func BenchmarkRvGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RvGet()
	}
}

func BenchmarkRvGet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RvGet2()
	}
}
