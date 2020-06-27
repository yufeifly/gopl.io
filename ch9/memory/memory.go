package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var x, y int
	wg.Add(1)
	go func() {
		x = 1                     // A1
		fmt.Println("y:", y, " ") // A2
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		y = 1                     // B1
		fmt.Println("x:", x, " ") // B2
		wg.Done()
	}()
	wg.Wait()
}
