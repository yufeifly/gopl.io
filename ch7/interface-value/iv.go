package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte{'n','i','\n'})
	fmt.Printf("%T\n",w)

	w = new(bytes.Buffer)
	w.Write([]byte("hello"))
	fmt.Fprintln(os.Stdout,w)
	fmt.Printf("%T\n",w)

	var x interface{} = time.Now()
	fmt.Fprintln(os.Stdout,x)
	fmt.Printf("%T\n",x)
}
