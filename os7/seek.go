package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 10)
	f, _ := os.Open("1.go")
	defer f.Close()
	f.Seek(1, 0) //相当于开始位置偏移1
	n, _ := f.Read(b)
	fmt.Println(string(b[:n])) //原字符package 输出ackage
}
