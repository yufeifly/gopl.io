package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordsCounter int

func(w *WordsCounter) Write(p []byte) (int, error) {
	s := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanLines)
	sum := 0
	for bs.Scan() {
		sum++
	}
	*w = WordsCounter(sum)
	return sum, nil
}

func main () {
	var c WordsCounter
	fmt.Fprintf(&c, "hello world 123\nnihaowoshi\nbi")
	fmt.Println(c) //输出 3
}

