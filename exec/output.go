package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c)) //	Sat Jan  4 17:07:36     2014 这个是标准库里的例子
}
