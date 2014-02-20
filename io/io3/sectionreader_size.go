package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()
	sr := io.NewSectionReader(f, 2, 5)
	fmt.Println(sr.Size()) //5
	p := make([]byte, 10)
	sr.Seek(1, 0) //相当于起始的地址偏移1
	n, err := sr.Read(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p[:n])) //lo w
	fmt.Println(sr.Size())     //5
}
