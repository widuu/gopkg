package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimFunc("&&&&nihao&&&&", func(r rune) bool {
		if r == '&' {
			return true
		}
		return false
	})) //nihao
}
