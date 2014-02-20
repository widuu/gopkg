package main

import (
	"fmt"
	"os"
)

func main() {
	//这个和stat差不多
	f, _ := os.Lstat("1.go")
	fmt.Println(f) //	&{1.go {32 {2883792444 30345044} {2884358853 30345044} {3464863792 30345293} 0 43} {0 0} D:\test\1.go 0 0 0}
}
