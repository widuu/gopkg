package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a := filepath.ToSlash("D:\\test\\path\\pathfile")
	fmt.Println(a) //D:/test/path/pathfile
}
