package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, _ := os.Open("write.txt")
	b := make([]byte, 20)
	num, err := io.ReadFull(r, b)
	defer r.Close()
	if err == io.EOF {
		fmt.Println("Read end total", num)
	}
	if err == io.ErrUnexpectedEOF {
		fmt.Println("Read fewer value:", string(b[:num])) //Read fewer value: hello widuu，依然是buf长度大于读取的长度
		return
	}

	fmt.Println("Read  value:", string(b)) //如果b是5 就出现这里
}
