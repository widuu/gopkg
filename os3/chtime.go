package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	err := os.Chtimes("2.go", time.Now(), time.Now()) //改变时间
	if err != nil {
		fmt.Println(err)
	}
	fi, _ := os.Stat("2.go")
	fmt.Println(fi.ModTime()) //输出时间 2013-12-29 20:46:23.0005257 +0800 +0800
}
