package main

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
	word, bit := x/64, uint(x%64) // 找到是第几个字，该字的第几个偏移
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

func (s *IntSet) AddAll(xs...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t. 并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith (t *IntSet) {
	var length int
	if len(s.words) < len(t.words) {
		length = len(s.words)
	} else {
		length = len(t.words)
	}
	s.words = s.words[:length]
	for i:=0; i< length;i++ {
		s.words[i] &= t.words[i]
	}
}

func(s *IntSet) Len() int {
	cnt := 0
	for _, tword := range s.words {
		if tword == 0 {
			continue
		}
		for i:=0;i<64;i++ {
			if tword & (1 << i) != 0 {
				cnt++
			}
		}
	}
	return cnt
}

func (s *IntSet) Remove(x int) {
	word, bit := x / 64, uint64(x%64)
	if word < len(s.words) {
		s.words[word] = s.words[word] &^ (1<<bit)
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var intset IntSet
	for _, word := range s.words {
		intset.words = append(intset.words, word)
	}
	return &intset
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

func main () {
	var x,y IntSet
	x.AddAll(1,144,9,19,12,14,16,11)
	y.AddAll(1,25,9,10,11)
	x.IntersectWith(&y)
	fmt.Println(x.String())
	fmt.Println(&x)

}