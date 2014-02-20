package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs("../clean.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir) //D:\test\path\clean.go
}
