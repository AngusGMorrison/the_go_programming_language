// Write a set of tests for IntSet (§6.5) that checks its behaviour after each operation is
// equivalent to a set based on built-in maps. Save your implementation for benchmarking in
// Exercise 11.7.
package intset

import (
	"testing"
)

func TestAddHas(t *testing.T) {
	tests := [][]int{
		{},
		{1, 64, 75},
	}

	for _, testHas := range tests {
		set := &IntSet{}
		intmap := make(map[int]bool)

		for _, i := range testHas {
			set.Add(i)
			if i < 0 {
				i *= -1 // insert absolute value of i into map
			}
			intmap[i] = true
		}

		testNotHas := []int{2, 9, 306}
		for _, s := range [][]int{testHas, testNotHas} {
			for _, i := range s {
				_, ok := intmap[i]
				if set.Has(i) != ok {
					t.Errorf("set.Has(%d): want %t", i, ok)
				}
			}
		}
	}

}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		set1Has, set2Has []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{13, 29}},
		{[]int{13, 29}, []int{}},
		{[]int{13, 29}, []int{13, 29}},
		{[]int{2, 54, 241}, []int{5, 38, 111}},
		{[]int{2, 54, 241}, []int{5, 38, 111, 316}},
		{[]int{2, 54, 241, 316}, []int{5, 38, 111}},
	}

	for _, test := range tests {
		set1 := &IntSet{}
		set2 := &IntSet{}
		intmap := make(map[int]bool)

		for _, i := range test.set1Has {
			set1.Add(i)
			intmap[i] = true
		}

		for _, i := range test.set2Has {
			set2.Add(i)
			intmap[i] = true
		}

		set1.UnionWith(set2)
		for k := range intmap {
			if !set1.Has(k) {
				t.Errorf("(%v).UnionWith(%v): set1 missing expected value %d", set1, set2, k)
			}
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		in   []int
		want string
	}{
		{[]int{}, "{}"},
		{[]int{1, 214, 320}, "{1 214 320}"},
	}

	for _, test := range tests {
		set := &IntSet{}
		for _, i := range test.in {
			set.Add(i)
		}

		got := set.String()
		if got != test.want {
			t.Errorf("(%v).String(): want %q, got %q", test.in, test.want, got)
		}
	}
}
