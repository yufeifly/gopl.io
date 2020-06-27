package bank3

import "sync"

var (
	mu      sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	b := balance
	return b
}
