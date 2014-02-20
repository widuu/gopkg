package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("hello widuu")
	w, _ := os.Create("test.txt")
	defer w.Close()
	n, err := reader.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) //11
}
