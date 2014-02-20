package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("bin") //打开一个目录
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	ff, _ := f.Readdir(10) //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
	for i, fi := range ff {
		fmt.Printf("filename %d: %+v\n", i, fi.Name()) //我们输出文件的名称
	}
}
