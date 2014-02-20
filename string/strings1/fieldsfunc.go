package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.FieldsFunc("widuunhellonword", split)) //	[widuu hello word]根据n字符分割
}

func split(s rune) bool {
	if s == 'n' {
		return true
	}
	return false
}
