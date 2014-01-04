package main

import (
	"fmt"
	"os"
)

func main() {
	p, _ := os.FindProcess(7588) //现在我本机windows下的cmd
	fmt.Println(p)               //	&{7588 224 0}
	err := p.Kill()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("杀死进程")
}
