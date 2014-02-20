package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("2.go", os.O_RDWR, os.ModePerm)
	n, _ := f.WriteString("hello word widuu")
	fmt.Println(n)
	b := make([]byte, n)
	f.Seek(0, 0)
	c, _ := f.Read(b)
	fmt.Println(string(b[:c]))
}
