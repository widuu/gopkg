package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//官方的标准库的例子
	cmd := exec.Command("sleep", "5") //执行等待5秒
	err := cmd.Start()                //开始执行
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Waiting for command to finish...\n")
	err = cmd.Wait() //wait下边的函数等待执行完成
	fmt.Printf("Command finished with error: %v\n", err)
}
