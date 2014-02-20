package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/a/b/c", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建文件夹" + dir + "/a/b/c成功")
}
