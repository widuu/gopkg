package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	err := &exec.Error{
		Name: "widuu",
		Err:  errors.New("is not exists widuu"),
	}

	fmt.Println(err.Error()) //	exec: "widuu": is not exists widuu
}
