package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("test2.txt")
	reader := io.TeeReader(r, w)
	fmt.Println(reflect.TypeOf(reader)) //*io.teeReader
	p := make([]byte, 10)
	n, _ := reader.Read(p)
	fmt.Println(string(p[:n])) //hello widu
}
