package main

import(
"fmt"
)

/*
练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
*/
func main(){
	fmt.Println(RecoverTest(20))//返回 20
}

/*
1.原来只定义返回类型，现在给返回值取一个适当的名字,直接使用内部匿名函数修改这个值
2.使用defer机制，defer后面的函数调用会被延迟执行,遇到pannic后才会调用
3.利用闭包，函数内部使用匿名函数可以访问外部函数的变量
4.利用recover机制 会捕获pannic异常
*/
func RecoverTest(x int)(result int){
	defer func (){
		recover()
		result=x
	}()
	panic(x)
}
