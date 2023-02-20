package main

import (
	"fmt"
)

// 使用chan控制goroutine执行顺序
var (
	name = ""
	c    = make(chan int)
)

func main() {
	go f()
	// main goroutine 在这里阻塞直到goroutine向c发送消息
	// 实现main goroutine在goroutine执行完后再执行
	<-c
	fmt.Println(name)
}

func f() {
	name = "hello,world"
	// close(c)  无缓冲通道 在这里关闭通道也可实现目的
	c <- 1
}
