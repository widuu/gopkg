package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.Readlink("src") //我的linux返回的是/home/widuu/go/src
	fmt.Println(data)
}
