package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("md") //删除我们刚创建的md目录
	if err != nil {
		fmt.Println(err)
	}
}
