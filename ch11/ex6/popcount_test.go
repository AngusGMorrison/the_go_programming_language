// Write benchmarks to compare the popcount implementation in Section 2.6.2 with your solutions
// to Exercise 2.4 and Exercise 2.5. At what point does the table-based approach break even?
package popcount

import (
	"testing"
)

func benchmarkPopCountTable(b *testing.B, invocs int) {
	// Simulate the additional work of initializing the table (in the test context it is
	// initialized once for both PopCountTable and PopCountClear, then once more for PopCountTable).
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < invocs; j++ {
			PopCountTable(uint64(j))
		}
	}
}

func BenchmarkPopCountTable1(b *testing.B)       { benchmarkPopCountTable(b, 1) }
func BenchmarkPopCountTable10(b *testing.B)      { benchmarkPopCountTable(b, 10) }
func BenchmarkPopCountTable100(b *testing.B)     { benchmarkPopCountTable(b, 100) }
func BenchmarkPopCountTable1000(b *testing.B)    { benchmarkPopCountTable(b, 1000) }
func BenchmarkPopCountTable10000(b *testing.B)   { benchmarkPopCountTable(b, 10000) }
func BenchmarkPopCountTable100000(b *testing.B)  { benchmarkPopCountTable(b, 100000) }
func BenchmarkPopCountTable1000000(b *testing.B) { benchmarkPopCountTable(b, 1000000) }

func benchmarkPopCountClear(b *testing.B, invocs int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < invocs; j++ {
			PopCountClear(uint64(j))
		}
	}
}

func BenchmarkPopCountClear1(b *testing.B)       { benchmarkPopCountClear(b, 1) }
func BenchmarkPopCountClear10(b *testing.B)      { benchmarkPopCountClear(b, 10) }
func BenchmarkPopCountClear100(b *testing.B)     { benchmarkPopCountClear(b, 100) }
func BenchmarkPopCountClear1000(b *testing.B)    { benchmarkPopCountClear(b, 1000) }
func BenchmarkPopCountClear10000(b *testing.B)   { benchmarkPopCountClear(b, 10000) }
func BenchmarkPopCountClear100000(b *testing.B)  { benchmarkPopCountClear(b, 100000) }
func BenchmarkPopCountClear1000000(b *testing.B) { benchmarkPopCountClear(b, 1000000) }
