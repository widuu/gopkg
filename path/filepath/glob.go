package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	m, _ := filepath.Glob("*.go")
	for i, v := range m {
		fmt.Println(i, v) //0 abs.go 1 base.go 2 fromslash.go 3 glob.go
	}
}
