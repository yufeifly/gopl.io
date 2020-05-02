package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LimitedReader struct {
	Reader  io.Reader
	Limit   int
	current int
}

func (r *LimitedReader) Read(b []byte) (int, error) {
	if r.current >= r.Limit {
		return 0, io.EOF
	}

	// 如果b申请得太大了，就将其缩小一点
	if r.current+len(b) > r.Limit {
		b = b[:r.Limit-r.current]
	}

	n, err := r.Reader.Read(b)
	if err != nil {
		return n, err
	}
	r.current += n
	return n, nil
}

func LimitReader(r io.Reader, limit int) io.Reader {
	lr := LimitedReader{
		Reader: r,
		Limit:  limit,
	}
	return &lr
}

func main() {
	file, err := os.Open("/home/alan/go/src/gopl.io/ch7/ex7-5/limit") // 1234567890
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lr := LimitReader(file, 5)
	buf := make([]byte, 13)
	n, err := lr.Read(buf)
	for err == nil {
		fmt.Println(n, string(buf)) // 5 [49 50 51 52 53 0 0 0 0 0]
		n,err = lr.Read(buf)
	}
}