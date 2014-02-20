package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	linkerror := &os.LinkError{
		Op:  "widuu",
		Old: "old",
		New: "new error",
		Err: errors.New("error test"),
	}
	fmt.Println(linkerror.Error())
}
