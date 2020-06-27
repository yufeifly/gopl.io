package main

import (
	//      "filepath"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

/*
练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
*/

//接收命令行参数-v
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	//接收命令行参数,多个路径
	flag.Parse()
	roots := flag.Args()
	//如果没传递任何路径,给默认值
	if len(roots) == 0 {
		roots = []string{"/"}
	}

	for {
		sumFileSize(roots)
		time.Sleep(20 * time.Second)
	}
}

func sumFileSize(roots []string) {
	//发送和接收文件字节大小的channel
	fileSizes := make(chan int64)
	//goroutine的计数器
	var n sync.WaitGroup
	//循环命令行传递的路径
	for _, root := range roots {
		n.Add(1)
		//启动goroutine计算
		go walkDir(root, &n, fileSizes)
	}
	//启动goroutine,等待所有计算目录的goroutine结束
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//定时显示目录进度发送的channel
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	//select和loop循环,多路复用
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			//计算目录数,计算字节大小
			nfiles++
			nbytes += size
		case <-tick:
			//接收到定时channel打印进度
			printDiskUsage(nfiles, nbytes)
		}
	}
	//最后打印总计
	printDiskUsage(nfiles, nbytes) // final totals
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := path.Join(dir, entry.Name())
			//开启多个goroutine进行递归
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	//使用计数信号量逻辑限制太多并发
	sema <- struct{}{}
	entries, err := ioutil.ReadDir(dir)
	<-sema
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
