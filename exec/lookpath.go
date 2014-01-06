package main

import (
	"fmt"
	"os/exec"
)

func main() {
	f, err := exec.LookPath("php")
	if err != nil {
		fmt.Println("not install php")
	}
	fmt.Println(f) // 输出我的phpD:\PHP\\php.exe
}
