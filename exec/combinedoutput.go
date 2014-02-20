package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arv := []string{"-a"}
	c := exec.Command("ls", arv...)
	d, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d)) //输出和command里边的output一样哈
}
