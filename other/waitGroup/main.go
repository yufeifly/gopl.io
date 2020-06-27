package main

import (
	"fmt"
	"sync"
	"time"
)

func calc(w *sync.WaitGroup, i int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic error.")
		}
	}()

	fmt.Println("calc: ", i)
	time.Sleep(time.Second)
	w.Done()
}

func main() {
	// WaitGroup能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}
	// 阻塞主线程等到所有的goroutine执行完成
	wg.Wait()
	fmt.Println("all goroutine finish")
}
