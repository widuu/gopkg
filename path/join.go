package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("a", "b", "c"))          // a/b/c
	fmt.Println(path.Join("a", "", "c"))           // a/c
	fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c
}
