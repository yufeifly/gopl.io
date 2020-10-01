// 包级变量首先初始化，然后init函数，然后运行main函数
package main

import (
	"fmt"
)

var T int64 = a()

func init() {
	fmt.Println("init in main.go ")
}

func a() int64 {
	fmt.Println("calling a()")
	return 2
}
func main() {
	fmt.Println("calling main")
}
