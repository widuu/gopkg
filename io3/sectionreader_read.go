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
	p := make([]byte, 10)
	n, err := sr.Read(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p[:n])) //llo w
}
