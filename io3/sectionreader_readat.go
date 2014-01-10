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
	n, err := sr.ReadAt(p, 1)
	if err == io.EOF {
		fmt.Println(string(p[:n])) //	lo w
	}

}
