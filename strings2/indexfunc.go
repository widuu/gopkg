package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexFunc("nihaoma", split)) //3
}

func split(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}
