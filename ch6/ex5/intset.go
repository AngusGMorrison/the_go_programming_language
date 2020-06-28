// The type of each word used by IntSet is uint64, but 64-bit arithmetic may be inefficient on a
// 32-bit platform. Modify the program to use the uint type, which is the most efficient unsigned
// integer type for the platform. Instead of dividing by 64, define a constant holding the effective
// size of uint in bits, 32 or 64. You can use the perhaps too-clever expression
// 32 << (^uint(0) >> 63) for this purpose.
//
package intset

import (
	"bytes"
	"fmt"
)

const wordSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
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

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements in s.
func (s *IntSet) Len() int {
	var len uint
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
	word, bits := x/wordSize, uint(x%wordSize)
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
	copiedWords := make([]uint, len(s.words))
	for i, word := range s.words {
		copiedWords[i] = word
	}
	return &IntSet{copiedWords}
}
