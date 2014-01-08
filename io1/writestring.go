package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	w, _ := os.OpenFile("write1.txt", os.O_RDWR, os.ModePerm)
	n, err := io.WriteString(w, "ni hao ma")
	if err != nil {
		fmt.Println(err) //当我用os.open()的时候木有权限  悲催的~~输出write write1.txt: Access is denied.
	}
	defer w.Close()
	b, _ := ioutil.ReadFile("write1.txt")
	fmt.Println("write total", n) //write total 9
	fmt.Println(string(b))        // ni hao ma
}
