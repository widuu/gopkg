package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 100) //设置读取的字节数
	f, _ := os.Open("1.go")
	n, _ := f.Read(b)
	fmt.Println(n)
	fmt.Println(string(b[:n])) //输出内容
}
