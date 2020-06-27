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

func main() {
	var a []int
	v := p(a)
	fmt.Println(v)
}
