package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("1.go")
	fmt.Println(f.Fd()) //我的平台句柄是228
}
