package main

import "fmt"

func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main () {
	fmt.Printf("sum : %d\n",sum([]int{1,2,3,4,5,6,7}...))
}

