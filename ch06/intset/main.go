package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Printf("%s:%d\n", x.String(), x.Len()) // "{1 9 144}:3"

	y.Add(9)
	y.Add(42)
	fmt.Printf("%s:%d\n", y.String(), y.Len()) // "{9 42}:2"

	x.UnionWith(&y)
	fmt.Printf("%s:%d\n", x.String(), x.Len()) // "{1 9 42 144}:4"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len returns the number of elements in the set.
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += bits.OnesCount64(word)
	}
	return len
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
