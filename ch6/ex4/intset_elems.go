// Add a method Elems that returns a slice containing the elements of the set, suitable for
// iterating over with a range loop.
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, sword := range s.words {
		if i < len(t.words) {
			s.words[i] = sword & t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, sword := range s.words {
		if i >= len(t.words) {
			break
		}
		s.words[i] = sword &^ t.words[i]
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements in s.
func (s *IntSet) Len() int {
	var len uint64
	for _, word := range s.words {
		for word > 0 {
			len += word & 1
			word >>= 1
		}
	}
	return int(len)
}

// Remove removes x from s.
func (s *IntSet) Remove(x int) {
	word, bits := x/64, uint(x%64)
	if len(s.words) < word {
		return
	}
	s.words[word] &= ^(1 << bits)
}

// Clear removes all elements from s.
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of s.
func (s *IntSet) Copy() *IntSet {
	copiedWords := make([]uint64, len(s.words))
	for i, word := range s.words {
		copiedWords[i] = word
	}
	return &IntSet{copiedWords}
}

// AddAll adds all of its arguments to the s.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// Elems returns a slice containing the elements of s.
func (s *IntSet) Elems() []int {
	elems := []int{}
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<j) == 1 {
				elems = append(elems, i*64+j)
			}
		}
	}
	return elems
}
