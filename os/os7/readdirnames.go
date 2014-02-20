package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("bin")
	names, err := f.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}
	for i, name := range names {
		fmt.Printf("filename %d: %s\n", i, name)
	}
}
