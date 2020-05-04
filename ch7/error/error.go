package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))
	var err error = syscall.Errno(3)
	fmt.Println(err.Error())
	fmt.Println(err)
}
