package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir, _ := filepath.Rel("D:/test/", "D:/test/path/filepath/rel.go")
	fmt.Println(dir) //path\filepath\rel.go
}
