# 使用方法
1. 编译server代码和client代码。
```
go build server.go
go build client.go    
```
2. 运行
```
./server
./client fei@localhost:5900 // fei是用户名，5900是server端监听的端口号
// 随后还需要输入密码：123
// 进入界面：/Users/bytedance/go/src/gopl.io/ch8/ex8-2/server/ftp/ftpfile :
```

## put
用来将client端的文件传给server端。具体的操作是：在client界面，输入
```shell script
put somefile // somefile在client二进制文件的同一目录
// put操作会把client端的文件传到server端的/Users/bytedance/go/src/gopl.io/ch8/ex8-2/server/ftp/ftpfile目录
```

## get
用来将server端的文件传到client端。
```shell script
get somefile // somefile在server端的ftpfile目录
// get操作会把server的ftpfile目录中的某一文件下载到client二进制文件所在的目录
```

