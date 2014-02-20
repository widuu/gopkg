package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f1, _ := os.Create("1.txt")
	f2, _ := os.Create("2.txt")
	writer := io.MultiWriter(f1, f2)
	writer.Write([]byte("widuu"))
	//千万别这么逻辑来 ，我这是测试用的哈
	r1, _ := ioutil.ReadFile("1.txt")
	r2, _ := ioutil.ReadFile("2.txt")
	fmt.Println(string(r1)) //widuu
	fmt.Println(string(r2)) //widuu
}
