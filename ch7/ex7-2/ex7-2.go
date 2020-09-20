package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) {
	n, err := cw.Writer.Write(content)
	//fmt.Println(" wirter: ", cw.Writer)
	if err != nil {
		return n, err
	}
	cw.Count += n
	return n, nil
}

func CountingWriter(writer io.Writer) (io.Writer, *int) {
	cw := CountWriter{
		Writer: writer,
	}
	return &cw, &(cw.Count)
}

func main() {
	f, err := os.Open("C:\\Users\\yufei\\go\\src\\github.com\\yufeifly\\gopl.io\\gopl.io\\ch7\\ex7-2\\text")
	if err != nil {
		log.Printf("open file err: %v\n", err)
		os.Exit(1)
	}
	cw, _ := CountingWriter(f)
	//fmt.Println(cw)
	fmt.Fprintf(cw, "%s", "abc")

	//fmt.Println(*counter)
	//fmt.Println(cw)
	cw.Write([]byte("def"))
	//fmt.Println(*counter)
	//fmt.Println(cw)
}
