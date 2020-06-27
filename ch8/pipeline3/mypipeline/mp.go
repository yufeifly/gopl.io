package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	//close(out)
}

func square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	//close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}
