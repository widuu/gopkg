package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir := filepath.Base("D:/test/path/filepath/clean.go")
	fmt.Println(dir)
}
