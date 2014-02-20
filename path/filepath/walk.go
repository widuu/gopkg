package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkfunc(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}

func main() {
	filepath.Walk("D:/test/path/filepath", walkfunc)
}
