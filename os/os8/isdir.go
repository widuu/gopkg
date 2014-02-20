package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Stat("1.go")
	fi := f.Mode()
	fmt.Println(reflect.TypeOf(fi)) //os.FileMode
	fmt.Println(fi.IsDir())         //false 判断是否是目录
	fmt.Println(fi.IsRegular())     //true  判断是否是个常规文件
	fmt.Println(fi.Perm())          //-rw-rw-rw-	返回的unix浅显
	fmt.Println(fi.String())        //-rw-rw-rw- 	返回文件模式的字符串形式
}
