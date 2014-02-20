package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go w.Write([]byte("hello widuu"))
	newerr := errors.New("your daye 突然关闭了")
	w.CloseWithError(newerr)
	data := make([]byte, 10)
	_, err := r.Read(data)
	if err != nil {
		fmt.Println(err) //your daye 突然关闭了
	}
}
