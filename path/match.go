package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("*", "alll"))  //true nil
	fmt.Println(path.Match("*", "a/lll")) //false nil
	fmt.Println(path.Match("?", "alll"))  //false nil
	fmt.Println(path.Match("?", "a"))     //true nil
	fmt.Println(path.Match("a", "a"))     //true nil
	fmt.Println(path.Match("\\a", "a"))   //true nil
}
