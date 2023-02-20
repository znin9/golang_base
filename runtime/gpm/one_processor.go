package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 测试只有一个 processor
func main() {
	// GOMAXPROCS设置可同时执行的CPU的最大数量，并返回先前的设置
	// runtime.NumCPU的值。如果n＜1，则不会更改当前设置。当调度程序改进时，此调用将消失。
	runtime.GOMAXPROCS(1) // 这里将 1 修改为2 可以直观的在标准输出看到直观的效果

	var wg sync.WaitGroup
	wg.Add(2)
	go f(&wg)
	go f2(&wg)
	wg.Wait()
}

func f(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i * 100)
	}
}
