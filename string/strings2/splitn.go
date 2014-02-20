package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.SplitN("a,b,c", ",", 2)) //[a b,c]
}
