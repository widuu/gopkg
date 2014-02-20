package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	r.Close()
	_, err := w.Write([]byte("hello widuu")) //具体可以看下边的func (w *PipeWriter) Write(data []byte) (n int, err error)

	if err == io.ErrClosedPipe {
		fmt.Println("管道已经关闭无法写入") //管道已经关闭无法写入
	}
}
