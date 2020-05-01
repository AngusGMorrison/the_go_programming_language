// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses string.Join. (Section 1.6
// illustrates part of the time package, and Section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.

package benchmarks

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range os.Args[1:] {
			s += s + arg + sep
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}
