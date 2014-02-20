package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a := filepath.FromSlash("D:/test/path/filepath/abs.go")
	fmt.Println(a) //windows D:\test\path\filepath\abs.go
}
