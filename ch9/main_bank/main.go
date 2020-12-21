package main

import (
	"fmt"
	"github.com/yufeifly/gopl.io/ch9/bank"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Alice:
	wg.Add(1)
	go func() {
		bank.Deposit(200)                            // A1
		fmt.Println("now we have =", bank.Balance()) // A2
		wg.Done()
	}()

	// Bob:
	wg.Add(1)
	go func() {
		go bank.Deposit(100)                         // A1
		fmt.Println("now we have =", bank.Balance()) // A2
		wg.Done()
	}()
	wg.Wait()
}
