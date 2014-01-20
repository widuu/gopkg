package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu web")
	n, err := reader.Read(make([]byte, 10))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) //10
}
