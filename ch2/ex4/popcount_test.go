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

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(rand.Uint64())
	}
}
