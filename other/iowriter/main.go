package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyWriter struct {
	Writer io.Writer
}
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

//type fmt.Stringer interface {
//String() string
//}
func main() {
	mw := &MyWriter{Writer: os.Stdout}
	mw.Writer.Write([]byte("hello,yufei"))

	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
