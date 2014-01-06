package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-ll")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	d, _ := ioutil.ReadAll(stdout) //从管道里读取数据
	cmd.Wait()
	fmt.Println(string(d))
	/*
		total 5
		-rw-r--r--    1 Administ Administ      268 Jan  4 17:07 combinedoutput.go
		-rw-r--r--    1 Administ Administ      277 Jan  4 17:01 command.go
		-rw-r--r--    1 Administ Administ      209 Jan  4 16:46 lookpath.go
		-rw-r--r--    1 Administ Administ      241 Jan  4 17:07 output.go
		-rw-r--r--    1 Administ Administ      271 Jan  4 17:13 run.go
		-rw-r--r--    1 Administ Administ      438 Jan  4 17:24 start.go
		-rw-r--r--    1 Administ Administ      497 Jan  4 17:32 stderrpipe.go
		-rw-r--r--    1 Administ Administ      552 Jan  6 14:10 stdinpipe.go
		-rw-r--r--    1 Administ Administ      235 Jan  6 14:14 stdoutpipe.go
	*/
}
