package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Clean("a/c"))               // a/c
	fmt.Println(path.Clean("a//c"))              //以一个/代替// , a/c
	fmt.Println(path.Clean("a/c/."))             //清除. , a/c
	fmt.Println(path.Clean("a/c/b/.."))          // 清除内部..以及前面的元素b; a/c
	fmt.Println(path.Clean("111/../a/c"))        // 清除内部..以及前面的元素111; a/c
	fmt.Println(path.Clean("/../a/b/../././/c")) // 清除/..开头,..以及前面的元素b; /a/c
}
