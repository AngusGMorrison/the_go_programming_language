package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(rand.Uint64())
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(rand.Uint64())
	}
}
