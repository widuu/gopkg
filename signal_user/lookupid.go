package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.LookupId("S-1-5-21-955939588-3462822645-4196941772-500")
	if err != nil {
		fmt.Println(err) //Unknown directory 一样的windows bug
	}
	fmt.Println(user.HomeDir)
}
