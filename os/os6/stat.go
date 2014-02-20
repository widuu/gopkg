package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Stat("1.go")
	fmt.Println(f.Size())
}
