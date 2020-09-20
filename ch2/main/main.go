package main

import (
	"fmt"
	"github.com/yufeifly/gopl.io/gopl.io/ch2/popcount"
)

func main() {
	fmt.Printf("%d\n", popcount.PopCount(127))
	fmt.Printf("%d\n", popcount.PopCount2(127))
}
