package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("write.txt")
	num, err := io.Copy(w, w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num) //返回int64的11 打开我的write.txt正是test.txt里边的hello widuu
}
