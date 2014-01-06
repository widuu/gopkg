package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println(path)
	dir, _ := os.Getwd()
	err := os.Mkdir(dir+path+"md", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建目录" + dir + path + "md成功")
}
