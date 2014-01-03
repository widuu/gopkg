package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	f, _ := os.Open("views")
	err := f.Chdir()
	if err != nil {
		fmt.Println(err)
	}
	dir1, _ := os.Getwd()
	fmt.Println(dir1)
}
