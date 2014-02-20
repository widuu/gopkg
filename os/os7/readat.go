package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("11.go")
	b := make([]byte, 20)
	n, _ := f.ReadAt(b, 15)
	fmt.Println(n)
	fmt.Println(string(b[:n]))
}
