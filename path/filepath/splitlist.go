package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir := filepath.SplitList("D:/test/path;D:/test/path/filepath")
	for index, v := range dir {
		fmt.Println(index, v) //0 D:/test/path 1 D:/test/path/filepath

	}
}
