package bank

import "fmt"

var balance int

func Deposit(amount int) {
	balance = balance + amount
	fmt.Println(balance)
}
func Balance() int {
	fmt.Println("balance:", balance)
	return balance
}
