package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexFunc("nihaoma", split)) //6
}

func split(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}
