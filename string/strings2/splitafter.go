package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.SplitAfter("a,b,c,d", ",")) //[a, b, c, d]
}
