package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownUserError
	err = "这个我自己设定"
	fmt.Println(err.Error()) //user: unknown user 这个我自己设定
}
