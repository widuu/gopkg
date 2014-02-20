package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 4)) //[a, b, c, d,r]
	fmt.Println(strings.SplitAfterN("a,b,c,d,r", ",", 5)) //[a, b, c, d, r]
}
