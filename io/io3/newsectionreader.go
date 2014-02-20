package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Open("test.txt")
	sr := io.NewSectionReader(f, 2, 5)
	fmt.Println(reflect.TypeOf(sr)) //*io.SectionReader
}
