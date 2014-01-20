package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu")
	b := make([]byte, 8)
	n, _ := reader.Read(b)
	fmt.Println(string(b[:n])) //hello wi 这个时候没有读到要读d
	reader.Seek(2, 1)          //这个时候我们在读取位置 偏移2
	reader.UnreadByte()        //然后向前偏移1字节
	n1, _ := reader.Read(b)
	fmt.Println(string(b[:n1])) //uu
}
