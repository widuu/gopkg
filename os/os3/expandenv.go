package main

import (
	"fmt"
	"os"
)

func main() {
	data := "GOBIN PATH $GOBIN"
	fmt.Println(os.ExpandEnv(data)) //输出我本地的环境变量的GOBIN的地址GOBIN PATH C:\Go\bin
}
