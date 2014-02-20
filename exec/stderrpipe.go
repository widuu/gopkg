package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	c := exec.Command("mv", "hello")
	i, err := c.StderrPipe()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err = c.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	b, _ := ioutil.ReadAll(i)
	if err := c.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err) //Error: exit status 1 mv: missing file argument Try `mv --help' for more information.
	}
	fmt.Println(string(b))
}
