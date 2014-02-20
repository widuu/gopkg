package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Open("1.go")
	fmt.Println(reflect.ValueOf(fileinfo).Type()) //*os.fileStat
	fmt.Println(fileinfo)
}
