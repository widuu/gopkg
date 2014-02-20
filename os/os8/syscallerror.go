package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := &os.SyscallError{
		Syscall: "widuu",
		Err:     errors.New("syscall error"),
	}
	fmt.Println(err.Error()) //widuu: syscall error
}
