package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.Hostname()
	fmt.Println(data) //我是windows环境下返回我的win的主机名 WIDUU
}
