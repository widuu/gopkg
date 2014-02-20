package main

import (
	"fmt"
	"os"
)

func main() {
	r, w, _ := os.Pipe()
	fmt.Println(r, w) //&{0xc08402e120} &{0xc08402e180}
}
