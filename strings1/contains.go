package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("widuu", "wi")) //true
	fmt.Println(strings.Contains("wi", "widuu")) //false
}
