package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownUserIdError
	err = 2
	fmt.Println(err.Error()) //user: unknown userid 2
}
