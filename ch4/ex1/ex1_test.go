package main

import (
	"crypto/sha256"
	"testing"
)

var h1 [32]byte = sha256.Sum256([]byte("cbthafdgadfgq"))
var h2 [32]byte = sha256.Sum256([]byte("ertwertwertwe"))

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount(h1, h2)
	}
}

func BenchmarkPopCountWithStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountWithStdLib(h1, h2)
	}
}
