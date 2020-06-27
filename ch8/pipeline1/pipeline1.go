package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	square := make(chan int)
	// Counter
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
			//time.Sleep(1*time.Second)
		}
		//fmt.Println(1, "end")
		close(naturals)
	}()
	// Square
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			square <- x * x
			//time.Sleep(1 * time.Second)
		}
		close(square)
	}()
	//Printer (in main goroutine)
	for {
		ans, ok := <-square
		if !ok {
			break
		}
		fmt.Println(ans)
	}

}
