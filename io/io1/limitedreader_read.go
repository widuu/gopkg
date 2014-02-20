package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, _ := os.Open("test.txt")
	limitedreader := io.LimitedReader{
		R: reader,
		N: 20,
	}
	p := make([]byte, 10)
	var total int
	for {
		n, err := limitedreader.Read(p)
		if err == io.EOF {
			fmt.Println("read total", total)     //read total 11
			fmt.Println("read value", string(p)) //read value hello widuu
			break
		}
		total = total + n

	}

}
