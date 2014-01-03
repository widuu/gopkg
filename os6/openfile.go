package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("10.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	fmt.Println(f.Stat())
}
