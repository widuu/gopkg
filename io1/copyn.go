package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("write1.txt")
	num, err := io.CopyN(w, r, 5)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()
	b, _ := ioutil.ReadFile("write1.txt")
	fmt.Println(string(b)) //输出 hello
	fmt.Println(num)       //5
}
