package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu web")
	data, err := reader.ReadByte()
	if err == nil {
		fmt.Println(string(data))
		return
	}
	fmt.Println(err)
}
