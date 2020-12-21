package main

import (
	"flag"
	"fmt"
	"github.com/yufeifly/gopl.io/ch7/tempconv"
)

// 这里会设置默认值
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}