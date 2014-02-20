package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	b := filepath.HasPrefix("D:/test/golang/", "D:/")
	fmt.Println(b) //true
}
