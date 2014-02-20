package main

import (
	"fmt"
	"os/user"
)

func main() {
	/*
		type User struct {
		    Uid      string // 用户id
		    Gid      string // 用户组id
		    Username string //用户名
		    Name     string //用户全名
		    HomeDir  string //用户家目录
		}
	*/
	user, _ := user.Current()
	//我是windows平台
	fmt.Println(user.Gid)      //S-1-5-21-955939588-3462822645-4196941772-513
	fmt.Println(user.HomeDir)  //C:\Users\Administrator
	fmt.Println(user.Uid)      //S-1-5-21-955939588-3462822645-4196941772-500
	fmt.Println(user.Name)     //nil
	fmt.Println(user.Username) //widuu\Administrator
}
