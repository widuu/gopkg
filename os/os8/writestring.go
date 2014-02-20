package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("2.go", os.O_RDWR, os.ModePerm)
	n, _ := f.WriteString("hello word widuu") //写入字符串
	fmt.Println(n)
	b := make([]byte, n)
	f.Seek(0, 0) //一定要把偏移地址归0否则就一直在写入的结尾处
	c, _ := f.Read(b)
	fmt.Println(string(b[:c])) //返回hello word widuu
}
