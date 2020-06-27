package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	// 调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数。
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	// 如果原始参数已经是切片类型，我们该如何传递给sum？只需在最后一个参数后加上省略符。
	fmt.Printf("sum : %d\n", sum([]int{1, 2, 3, 4, 5, 6, 7}...))
}
