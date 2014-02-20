package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu web")
	r, n, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(r), n) //h,1
}
