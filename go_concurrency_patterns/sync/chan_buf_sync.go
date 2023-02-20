package main

// 带缓冲通道实现通信

var (
	name string
	c    = make(chan int, 10)
)

func f() {
	name = "hello,world"
	c <- 1
}

func main() {
	go f()
	<-c
	println(name)
}
