package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()
	reader := io.LimitReader(f, 5)
	p := make([]byte, 5)
	fmt.Println(reflect.TypeOf(reader)) //*io.LimitedReader
	var total int
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println("read value", string(p[:total])) //read value hello
			fmt.Println(total)                           //5
			break
		}
		total = total + n
	}

}
