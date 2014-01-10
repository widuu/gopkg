package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	f1, _ := os.Open("test1.txt")
	f2, _ := os.Open("test.txt")
	defer f1.Close()
	defer f2.Close()
	reader := io.MultiReader(f1, f2) //*io.multiReader
	fmt.Println(reflect.TypeOf(reader))
	p := make([]byte, 10)
	var total int
	var data string
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println("read end", total) //read end 17
			break
		}
		total = total + n
		data = data + string(p[:n])
	}
	fmt.Println("read value", data)  //read value widuu2hello widuu
	fmt.Println("read count", total) //	read count 17
}
