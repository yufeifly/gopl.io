package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) {
	n, err := cw.Writer.Write(content)
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
	cw, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "123")
	fmt.Println(*counter)
	cw.Write([]byte("456"))
	fmt.Println(*counter)
}