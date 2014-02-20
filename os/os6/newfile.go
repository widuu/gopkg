package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("ini.go")
	defer f.Close()
	f1 := os.NewFile(f.Fd(), "ceshi.go") //输如ini.go的句柄
	defer f1.Close()
	fd, _ := f1.Stat()
	fmt.Println(fd.ModTime()) //返回的是ini.go的创建时间2013-11-27 09:11:50.2793737 +0800 CST

}
