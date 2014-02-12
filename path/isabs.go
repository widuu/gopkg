package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("D:/gopkg/gopkg")) //false
	fmt.Println(path.IsAbs("/dev/null"))      //true
}
