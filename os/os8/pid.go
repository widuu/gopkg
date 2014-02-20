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
	fmt.Println(p) //&{44 252 0}
	pw, _ := p.Wait()
	fmt.Println(pw.Pid())        //44
	fmt.Println(pw.String())     //exit status 0
	fmt.Println(pw.Success())    //true
	fmt.Println(pw.Sys())        //{0}
	fmt.Println(pw.SysUsage())   //&{{1728533532 30345488} {1745163891 30345488} {1562500 0} {1562500 0}}
	fmt.Println(pw.SystemTime()) //156.25ms
	fmt.Println(pw.UserTime())   //156.25ms
}
