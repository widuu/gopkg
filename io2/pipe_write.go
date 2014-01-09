package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go w.Write([]byte("hello widuu")) //写入的是[]byte,注意官方文档写的是，写入管道阻塞，一直到所有数据的读取结束
	data := make([]byte, 11)
	n, _ := r.Read(data)
	fmt.Println(string(data))     //hello widuu
	fmt.Println("read number", n) //read number 10
}
