/*
ex6.1 Implement these additional methods: 
Click here to view code image
func (*IntSet) Len() int // return the number of elements
func (*IntSet) Remove(x int) // remove x from the set
func (*IntSet) Clear() // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/

// Package intset provides a set of integers based on a bit vector.
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
	fmt.Println("word: ", word)
	fmt.Printf("%8b\n", word)
	fmt.Println("bit: ", bit)
	fmt.Printf("%8b\n", bit)

	for word >= len(s.words) {
		fmt.Println(s.words)
		s.words = append(s.words, 0)
		fmt.Println(s.words)
	}
	fmt.Println("1 << bit:", 1 << bit)
	s.words[word] |= 1 << bit
	fmt.Println("s.words[word]: ", s.words[word])
	fmt.Printf("%8b\n", s.words[word])
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

// returns the number of elements
func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		count += popcount(word)
	}
	return
}

// returns number of 1 bits
func popcount(x uint64) (count int) {
	for x != 0 {
		count++
		x &= x - 1
	}
	return
}
