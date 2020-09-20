package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordsCounter struct {
	words, lines int
}

func (w *WordsCounter) Write(p []byte) (int, error) {
	s := strings.NewReader(string(p))
	s2 := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	ls := bufio.NewScanner(s2)
	bs.Split(bufio.ScanWords)
	ls.Split(bufio.ScanLines)
	wordCnt := 0
	for bs.Scan() {
		wordCnt++
		fmt.Println("bs ", bs.Text())
	}
	lineCnt := 0
	for ls.Scan() {
		lineCnt++
		fmt.Println("ls ", ls.Text())
	}
	*w = WordsCounter{
		words: wordCnt,
		lines: lineCnt,
	}
	return lineCnt, nil
}

func main() {
	var c WordsCounter
	fmt.Fprintf(&c, "hello world 123\nnihaowoshi\nbi")
	fmt.Println("haha ", c.words, c.lines) //输出 3
}
