package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "widuu_xiaowei"
	if strings.HasPrefix(s, "widuu_") {
		fmt.Println(strings.TrimPrefix(s, "widuu_")) //xiaowei
		return
	}

	fmt.Println(s)
}
