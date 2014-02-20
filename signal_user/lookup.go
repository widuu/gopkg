package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Lookup("widuu\\Administrator")
	if err != nil {
		/*
		  我是windows平台 返回Unknown directory为什么呢？在标准库里我们可以看到有个BUG是这样写的
		  Lookup and LookupId functions do not set Gid and HomeDir fields in the User struct returned on windows.
		  这里我们就明白了，widnows唉唉唉~~~~
		*/
		fmt.Println(err)
	}
	fmt.Println(user.HomeDir)
}
