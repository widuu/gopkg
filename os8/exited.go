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
	fmt.Println(p) //&{7764 252 0}
	pw, _ := p.Wait()
	fmt.Println(pw.Exited()) //当我关闭记事本的时候进程结束，所以就算true
}
