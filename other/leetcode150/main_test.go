package main

import "testing"

func TestEvalRPN(t *testing.T) {
	var cases = []struct {
		in     []string
		expect int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 7},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, val := range cases {
		actual := EvalRPN(val.in)
		if actual != val.expect {
			t.Errorf("expect=%v,actual=%v", val.expect, actual)
		} else {
			t.Log("good")
		}
	}
}
