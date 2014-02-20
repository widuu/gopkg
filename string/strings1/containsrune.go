package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsRune("widuu", rune('w'))) //true
	fmt.Println(strings.ContainsRune("widuu", 20))        //fasle
}
