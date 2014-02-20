package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "web_widuu"
	if strings.HasSuffix(s, "_widuu") {
		fmt.Println(strings.TrimSuffix(s, "_widuu")) //web
		return
	}
	fmt.Println(s)

}
