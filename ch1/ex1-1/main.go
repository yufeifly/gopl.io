package main

import (
	"fmt"
	"os"
)

func PrintArgs() {
	for ind, val := range os.Args {
		fmt.Println("ind: ", ind, " val: ", val)
	}
}

func main() {
	PrintArgs()
}
