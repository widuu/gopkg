package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'i'
		}
		return r
	}, "hello waduu")) //hello widuu
}
