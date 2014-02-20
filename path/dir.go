package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/a/b/c")) // /a/b
	fmt.Println(path.Dir(""))       // .
}
