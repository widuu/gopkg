package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	w, _ := os.OpenFile("test1.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer w.Close()
	n, err := r.WriteString(w, "This is <b>HTML</b>!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	d, _ := ioutil.ReadFile("test1.txt") //32
	fmt.Println(string(d))               //This is &lt;b&gt;HTML&lt;/b&gt;!
}
