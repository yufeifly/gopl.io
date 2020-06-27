package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	//conn.Close()

	// 这里可以这么理解：我们先把网络连接的写的部分关闭，然后再从通道中读数据。这就说明即使关闭了连接的写部分，还是可以从通道中读取
	//类型断言，调用*net.TCPConn的方法CloseWrite()只关闭TCP的写连接
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done // wait for background goroutine to finish
}
