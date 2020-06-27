package main

import "fmt"

func trailingZeroes(n int) int {
	fives := 0
	for i := 5; i <= n; i += 5 {
		fives += get5(i)
		//fmt.Println(fives)
	}
	return fives
}

func get5(number int) int {
	five := 0
	for number%5 == 0 {
		number = number / 5
		five++
	}
	return five
}

func main() {
	fmt.Println(trailingZeroes(30))
}
