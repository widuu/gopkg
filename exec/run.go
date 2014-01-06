package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c, err := exec.Command("date")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Run()
	if err != nil {
		fmt.Println(err)
	}
	d, _ := c.Output()
	fmt.Println(string(d)) //	Sat Jan  4 17:07:36
}
