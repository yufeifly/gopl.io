package main

import "fmt"

func HammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		num = num & (num - 1)
		sum++
	}
	return sum
}

func main() {
	fmt.Println(HammingWeight(7))
}
