package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu web")
	b := make([]byte, 10)
	n, err := reader.ReadAt(b, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b[:n])) //llo widuu
}
