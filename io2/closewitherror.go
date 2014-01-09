package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	r.Close()
	err := errors.New("管道符关闭了") //errors这个包我们前边已经说过了，就一个方法New不会的可以看看前边的
	r.CloseWithError(err)
	_, err = w.Write([]byte("test"))
	if err != nil {
		fmt.Println(err) //管道符关闭了
	}
}
