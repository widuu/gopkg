package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	r, w := io.Pipe()
	fmt.Println(reflect.TypeOf(r)) //*io.PipeReader
	fmt.Println(reflect.TypeOf(w)) //*io.PipeWriter
}
