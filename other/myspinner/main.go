package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	ans := fib(45)
	fmt.Printf("\r%v", ans) // \r的意思是回到行首
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func spinner(t time.Duration) {
	flags := `-/|\`
	for {
		for _, v := range flags {
			fmt.Printf("\r%c", v)
			time.Sleep(t)
		}
	}
}
