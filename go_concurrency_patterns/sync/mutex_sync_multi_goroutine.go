package main

import (
	"fmt"
	"sync"
	"time"
)

// 验证在不同的协程之间同时lock,是否unlock之后的执行顺序是相同的

var (
	l    sync.Mutex
	name string
	w    sync.WaitGroup
)

func f(id int) {
	defer w.Done()
	l.Lock() // 会发生锁的竞争,多个goroutine阻塞在这时，不确定哪个goroutine先执行
	time.Sleep(time.Second * 1)
	fmt.Println(id, name)
	l.Unlock()
}

func main() {
	for i := 1; i <= 10; i++ {
		w.Add(1)
		go f(i)
	}
	w.Wait()
}
