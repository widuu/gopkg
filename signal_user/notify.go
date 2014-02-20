package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//设置一个channel来发送信号
	c := make(chan os.Signal, 1)
	// 一直运行一直到收到一个信号
	signal.Notify(c, os.Interrupt, os.Kill)

	//终端信号的接收
	//signal.Stop(c)   //程序运行结过exit status 2
	s := <-c

	fmt.Println("Got signal:", s) //当我停止运行时 Got signal: interrupt

}
