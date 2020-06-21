package main

import (
	"fmt"
	"runtime"
	"testing"
)

const maxGortns = 100

func benchmarkMaxProcs(b *testing.B, maxProcs int) {
	fmt.Printf("%d procs\n", maxProcs)
	runtime.GOMAXPROCS(maxProcs)
	for i := 0; i < b.N; i++ {
		buildImage(maxGortns)
	}
}

func Benchmark1(b *testing.B) { benchmarkMaxProcs(b, 1) }
func Benchmark2(b *testing.B) { benchmarkMaxProcs(b, 2) }
func Benchmark3(b *testing.B) { benchmarkMaxProcs(b, 3) }
func Benchmark4(b *testing.B) { benchmarkMaxProcs(b, 4) }
func Benchmark5(b *testing.B) { benchmarkMaxProcs(b, 5) }
func Benchmark6(b *testing.B) { benchmarkMaxProcs(b, 6) }
func Benchmark7(b *testing.B) { benchmarkMaxProcs(b, 7) }
func Benchmark8(b *testing.B) { benchmarkMaxProcs(b, 8) }
