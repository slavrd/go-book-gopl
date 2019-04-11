// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

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

// AddAll adds all of the provided values to the set
func (s *IntSet) AddAll(l ...int) {
	for _, item := range l {
		s.Add(item)
	}
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
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

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove removes the provided x from the set
func (s *IntSet) Remove(x int) {

	word, bit := x/64, x%64
	if len(s.words) > word {
		s.words[word] = s.words[word] &^ (1 << uint(bit))
	}
}

// Clear removes all items form the set
func (s *IntSet) Clear() {
	if len(s.words) > 0 {
		s.words = make([]uint64, 0, 0)
	}
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var result IntSet
	result.words = make([]uint64, len(s.words))
	copy(result.words, s.words)
	return &result
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {

	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = s.words[:i]
			break
		}
	}

}

// DifferenceWith sets s to the difference with t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}

}

// SymmetricDifference sets s to the symmetric difference with t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns s's elements as s slice
func (s *IntSet) Elems() *[]int {
	r := make([]int, 0, 0)
	for i, word := range s.words {
		for j := 0; word != 0; j++ {
			if word&1 == 1 {
				r = append(r, 64*i+j)
			}
			word >>= 1
		}
	}
	return &r
}

//!-intset
