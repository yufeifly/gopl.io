package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("launch the rocket!")
}

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(500 * time.Millisecond)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
