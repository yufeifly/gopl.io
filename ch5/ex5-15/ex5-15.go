package main

import (
	"errors"
	"fmt"
	"os"
)

func max(digits ...int) (int,error) {
	if len(digits) < 1 {
		return -1, errors.New("need at least one arg")
	}
	maxV := digits[0]
	for _,val := range digits {
		if val > maxV {
			maxV = val
		}
	}
	return maxV,nil
}

func main() {
	val,err := max()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("max value is : %d\n",val)

}