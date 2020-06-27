package main

import "fmt"

func IsHappy(n int) bool {
	tmp := n
	arr := []int{n}
	for {
		s := 0
		vec := func(n int) (vec []int) {
			for n > 0 {
				vec = append(vec, n%10)
				n = n / 10
			}
			return
		}(tmp)
		for _, v := range vec {
			s += v * v
		}
		if s == 1 {
			return true
		}
		for _, val := range arr {
			if val == s {
				return false
			}
		}
		arr = append(arr, s)
		tmp = s
	}
	return false
}

func main() {
	fmt.Println(IsHappy(2))
}
