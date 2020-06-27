package main

import (
	"fmt"
	"strconv"
)

func EvalRPN(tokens []string) int {
	st := []int{}
	for _, val := range tokens {
		len := len(st)
		switch val {
		case "+":
			st = append(st[:len-2], st[len-2]+st[len-1])
		case "-":
			st = append(st[:len-2], st[len-2]-st[len-1])
		case "*":
			st = append(st[:len-2], st[len-2]*st[len-1])
		case "/":
			st = append(st[:len-2], st[len-2]/st[len-1])
		default:
			num, _ := strconv.Atoi(val)
			st = append(st, num)
		}
	}
	return st[0]
}

func main() {
	str := []string{"2", "1", "+", "3", "*"}
	fmt.Println(EvalRPN(str))
}
