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
	fmt.Println(p)
	if err := p.Signal(os.Kill); err != nil { //向系统发送一个干掉他的信号 所以我们熟悉的记事本不出来了
		fmt.Println(err)
	}
}
