package main

import (
	"fmt"
	"os"
)

func main() {
	p, _ := os.FindProcess(6816) //我的windows下的cmd进程
	fmt.Println(p)               //&{6816 236 0}
}
