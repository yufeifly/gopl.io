package main

import (
	"fmt"
	"strings"
)

func Print1(ss []string) string {
	var ans string
	for _, val := range ss[:len(ss)-1] {
		ans += val + " "
	}
	ans += ss[len(ss)-1]
	return ans
}

func Print2(ss []string) string {
	ans := strings.Join(ss, " ")
	return ans
}

func main() {
	strvec := []string{"i", "am", "yufeixiong"}
	s1 := Print1(strvec)
	s2 := Print2(strvec)
	fmt.Println(s1)
	fmt.Println(s2)
}
