package main

import (
	"fmt"
	"sync"
)

// sync.Once

var (
	a    string
	once sync.Once
	wg   sync.WaitGroup
)

func setup() {
	fmt.Println("call setup")
	a = "hello,world"
}

func printA() {
	defer wg.Done()
	once.Do(setup) // 无论调用多次少(即使是多个goroutine调用)都只会执行一次,本质通过Mutex双重检查锁实现的
	println(a)
}

func main() {
	wg.Add(3)
	go printA()
	go printA()
	go printA()
	wg.Wait()
}
