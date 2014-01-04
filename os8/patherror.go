package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	patherr := &os.PathError{
		Op:   "dir",
		Path: "widuu",
		Err:  errors.New("no path"),
	}
	fmt.Println(patherr.Error()) //dir widuu: no path
}
