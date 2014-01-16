package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"hello", "word", "xiaowei"}
	fmt.Println(strings.Join(s, "-")) //	hello-word-xiaowei
}
