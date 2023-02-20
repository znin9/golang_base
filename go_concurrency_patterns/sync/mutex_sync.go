package main

import (
	"fmt"
	"sync"
)

// 使用sync.Mutex实现goroutine的同步
var (
	l    sync.Mutex
	name = ""
)

func f() {
	name = "hello,world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock() // 若当前l,并未unlock,那么会阻塞在这里
	fmt.Println(name)
}
