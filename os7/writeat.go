package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("1.go", os.O_RDWR, os.ModePerm)
	f.WriteAt([]byte("widuu"), 10) //在偏移10的地方写入
	b := make([]byte, 20)
	d, _ := f.ReadAt(b, 10)    //偏移10的地方开始读取
	fmt.Println(string(b[:d])) //widuudhellowordhello
}
