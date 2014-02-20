package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Create("1.go")
	defer f.Close()
	fmt.Println(reflect.ValueOf(f).Type()) //*os.File
}
