package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimLeftFunc("&widuu&", func(r rune) bool {
		if r == '&' {
			return true
		}
		return false
	})) //	widuu&是不是跟上边的那个函数效果一样了
}
