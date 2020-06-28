package main

import "testing"

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildImage(1) // 311 ms
	}
}

func BenchmarkFewGortns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildImage(2) // 161 ms
	}
}

func BenchmarkOptimalGortns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildImage(5) // 88 ms
	}
}

func BenchmarkHighGortns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildImage(10) // 84 ms
	}
}
