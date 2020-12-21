package main

import "fmt"

func p(thing interface{}) interface{} {
	t := thing.([]int)
	t = append(t, 1)
	thing = t
	print(thing)
	return thing
}

func print(v interface{}) {
	for _, val := range v.([]int) {
		fmt.Println(val)
	}
}

func adder(a []int) []int {
	a[0] = 100
	a = append(a, []int{11, 22, 33}...)
	fmt.Println("a in main ", a, " ", len(a), " ", cap(a))
	return a
}

type T struct {
	x int
}

// func (t *T) String() string            { return "boo" }
// func (t T) Format(s fmt.State, _ rune) { fmt.Fprint(s, "baa") }

func (t T) String() string             { return "boo" }
func (t T) Format(s fmt.State, _ rune) { fmt.Fprint(s, "baa") }

func main() {
	v := &T{123}
	fmt.Println(v)
	error.Error()
}
