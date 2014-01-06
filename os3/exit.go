package main

import (
	"fmt"
	"os"
)

func main() {
	func() {
		for {
			fmt.Println("这个是匿名函数")
			os.Exit(1) //输出exit status 1中断操作
		}
	}()
}
