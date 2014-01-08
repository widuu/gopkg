package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, _ := os.Open("write1.txt")
	b := make([]byte, 2)
	defer r.Close()
	var total int
	for {
		n, err := io.ReadAtLeast(r, b, 8)
		if err == nil {
			fmt.Println("Read enough value:", string(b)) //	Read enough value: hello widuu
		}
		if err == io.ErrUnexpectedEOF { //读取了的数据小于我们限定的最小读取数据8
			fmt.Println("Read fewer value:", string(b[0:n]))
		}
		if err == io.EOF { //读完了 输出
			fmt.Println("Read end total", total) //Read end total 11
			break
		}
		if err == io.ErrShortBuffer {
			fmt.Println("buf too Short")
			os.Exit(1)
		}
		total = total + n
	}
}
