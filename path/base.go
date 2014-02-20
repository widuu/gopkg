package main

import (
	"fmt"
	"path"
)

func main() {
	path := path.Base("D:/gopkg/gopkg")
	fmt.Println(path) //gopkg
}
