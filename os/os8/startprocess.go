package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	p, err := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr) //windows打开文件夹下1.txt
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p) //&{5488 240 0}
}
