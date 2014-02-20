package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "widuuweb"
	fmt.Println(strings.Replace(s, "u", "a", 1))  //widauweb
	fmt.Println(strings.Replace(s, "u", "a", -1)) //widaaweb
	fmt.Println(s)                                //widuuweb
}
