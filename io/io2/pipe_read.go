package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go w.Write([]byte("hello widuu"))
	d := make([]byte, 11)
	n, _ := r.Read(d) //从管道里读取数据
	fmt.Println(string(d))
	fmt.Println(n)
}
