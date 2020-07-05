package intset

import (
	"math/rand"
	"testing"
)

const (
	randmax = 1 << 24
	count   = 100
)

var (
	randInts1 = make([]int, 0)
	randInts2 = make([]int, 0)
)

func init() {
	for i := 0; i < count; i++ {
		randInts1 = append(randInts1, rand.Intn(randmax))
		randInts2 = append(randInts2, rand.Intn(randmax))
	}
}

func benchmarkAdd(b *testing.B, set IntSet) {
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(randmax))
	}
}

func benchmarkAddAll(b *testing.B, set IntSet) {
	for i := 0; i < b.N; i++ {
		set.AddAll(randInts1...)
	}
}

func benchmarkHas(b *testing.B, set IntSet) {
	set.AddAll(randInts1...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, x := range randInts1 {
			set.Has(x)
		}
	}
}

func benchmarkUnionWith(b *testing.B, newSet func() IntSet) {
	set1, set2 := newSet(), newSet()
	set1.AddAll(randInts1...)
	set2.AddAll(randInts2...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

func benchmarkString(b *testing.B, set IntSet) {
	set.AddAll(randInts1...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

// BitIntSet64 benchmarks
func BenchmarkBitIntSet64Add(b *testing.B)    { benchmarkAdd(b, &BitIntSet64{}) }
func BenchmarkBitIntSet64AddAll(b *testing.B) { benchmarkAddAll(b, &BitIntSet64{}) }
func BenchmarkBitIntSet64Has(b *testing.B)    { benchmarkHas(b, &BitIntSet64{}) }
func BenchmarkBitIntSet64UnionWith(b *testing.B) {
	benchmarkUnionWith(b, func() IntSet { return &BitIntSet64{} })
}
func BenchmarkBitIntSet64String(b *testing.B) { benchmarkHas(b, &BitIntSet64{}) }

// BitIntSet32 benchmarks
func BenchmarkBitIntSet32Add(b *testing.B)    { benchmarkAdd(b, &BitIntSet32{}) }
func BenchmarkBitIntSet32AddAll(b *testing.B) { benchmarkAddAll(b, &BitIntSet32{}) }
func BenchmarkBitIntSet32Has(b *testing.B)    { benchmarkHas(b, &BitIntSet32{}) }
func BenchmarkBitIntSet32UnionWith(b *testing.B) {
	benchmarkUnionWith(b, func() IntSet { return &BitIntSet32{} })
}
func BenchmarkBitIntSet32String(b *testing.B) { benchmarkHas(b, &BitIntSet32{}) }

// MapIntSet benchmarks
func BenchmarkMapIntSetAdd(b *testing.B)    { benchmarkAdd(b, &MapIntSet{map[int]bool{}}) }
func BenchmarkMapIntSetAddAll(b *testing.B) { benchmarkAddAll(b, &MapIntSet{map[int]bool{}}) }
func BenchmarkMapIntSetHas(b *testing.B)    { benchmarkHas(b, &MapIntSet{map[int]bool{}}) }
func BenchmarkMapIntSetUnionWith(b *testing.B) {
	benchmarkUnionWith(b, func() IntSet { return &MapIntSet{map[int]bool{}} })
}
func BenchmarkMapIntSetString(b *testing.B) { benchmarkHas(b, &MapIntSet{map[int]bool{}}) }
