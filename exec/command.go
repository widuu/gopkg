package main

import (
	"fmt"
	"os/exec"
)

func main() {
	argv := []string{"-a"}
	c := exec.Command("ls", argv...)
	d, _ := c.Output()
	fmt.Println(string(d)) //因为装的git bash所以可以用ls -a
	/*
	 *	.
	 *	..
	 *	command.go
	 *	lookpath.go
	 */
}
