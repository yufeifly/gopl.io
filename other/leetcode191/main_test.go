package main

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	num := 3
	expect := 2
	actual := HammingWeight(uint32(num))
	if actual != expect {
		t.Errorf("expect=%d,actual=%d", expect, actual)
	}

}
