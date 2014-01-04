package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	p, _ := os.StartProcess(`c:\windows\system32\notepad.EXE`, []string{`c:\windows\system32\notepad.EXE`, `d:/1.txt`}, attr)
	pw, _ := p.Wait()
	fmt.Println(pw) //exit status 0
}
