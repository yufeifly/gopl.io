package main

import "testing"

func TestIsHappy(t *testing.T) {
	var cases = []struct {
		in     int
		expect bool
	}{
		{2, false},
		{19, true},
	}
	for _, val := range cases {
		actual := IsHappy(val.in)
		if actual != val.expect {
			t.Errorf("expect=%v,actual=%v", val.expect, actual)
		}
	}

}
