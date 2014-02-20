package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	var output bytes.Buffer
	cmd := exec.Command("cat")
	cmd.Stdout = &output
	stdin, _ := cmd.StdinPipe()
	cmd.Start() //执行开始
	stdin.Write([]byte("widuu test"))
	stdin.Close()
	cmd.Wait()                                        //等待执行完成
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: widuu test!
}
