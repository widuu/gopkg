package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("widuu", "uu")) //1
	fmt.Println(strings.Count("widuu", "u"))  //2
}
