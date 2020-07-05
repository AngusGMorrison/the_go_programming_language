package intset

import (
	"bytes"
	"fmt"
)

// A BitIntSet32 is a set of small non-negative integers. Its zero value represents the empty set.
type BitIntSet32 struct {
	words []uint32
}

// Has reports whether the set contains the non-negative value x.
func (s *BitIntSet32) Has(x int) bool {
	word, bit := x/32, uint32(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *BitIntSet32) Add(x int) {
	word, bit := x/32, uint32(x%32)
	// A single allocation to extend the length of s.words reduces runtime ~25%.
	s.allocMaxLenRequired(word)
	s.words[word] |= 1 << bit
}

// AddAll adds a variable number of non-negative integers to the set.
func (s *BitIntSet32) AddAll(ints ...int) {
	// Find the largest word and allocate sufficient memory to hold it and all other words in a
	// single operation. ~13% runtime reduction.
	max := 0
	for _, x := range ints {
		if x > max {
			max = x
		}
	}

	s.allocMaxLenRequired(max / 32)
	for _, x := range ints {
		s.Add(x)
	}
}

func (s *BitIntSet32) allocMaxLenRequired(word int) {
	if word >= len(s.words) {
		growth := make([]uint32, word-len(s.words)+1)
		s.words = append(s.words, growth...)
	}
}

// UnionWith sets s to the union of s and t.
func (s *BitIntSet32) UnionWith(t IntSet) {
	switch t := t.(type) {
	case *BitIntSet32:
		for i, tword := range t.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	default:
		s.AddAll(t.Ints()...)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *BitIntSet32) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *BitIntSet32) Ints() []int {
	var ints []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				ints = append(ints, i*32+j)
			}
		}
	}
	return ints
}
