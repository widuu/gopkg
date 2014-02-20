package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("1.go")
	fmt.Println(f.Name()) //输出1.go
}
