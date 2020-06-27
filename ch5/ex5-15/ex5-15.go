package main

import (
	"errors"
	"fmt"
	"os"
)

func max(digits ...interface{}) (interface{}, error) {
	if len(digits) < 1 {
		return -1, errors.New("need at least one arg")
	}
	maxV := digits[0]
	for _, val := range digits {
		if val.(int) > maxV.(int) {
			maxV = val
		}
	}
	return maxV, nil
}

func main() {
	val, err := max([]interface{}{7, 5, 7, 8, 100, 34}...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("max value is : %d\n", val)

}
