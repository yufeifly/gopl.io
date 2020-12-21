package main

import (
	"fmt"
	"github.com/yufeifly/gopl.io/ch9/bank2"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Alice:
	wg.Add(1)
	go func() {
		bank2.Deposit(200)                            // A1
		fmt.Println("now we have =", bank2.Balance()) // A2
		wg.Done()
	}()

	// Bob:
	wg.Add(1)
	go func() {
		go bank2.Deposit(100)                         // A1
		fmt.Println("now we have =", bank2.Balance()) // A2
		wg.Done()
	}()
	wg.Wait()
}
