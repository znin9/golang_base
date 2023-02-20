package main

import (
	"fmt"
	"sync/atomic"
)

// 原子的同步操作
func main() {
	var a int64
	// Store
	atomic.StoreInt64(&a, 100)
	// Load
	val := atomic.LoadInt64(&a)
	fmt.Println("val:", val)
	// Add
	addVal := atomic.AddInt64(&a, 100)
	fmt.Println("addVal:", addVal)
	// 交换
	oldVal := atomic.SwapInt64(&a, 300)
	fmt.Println("oldVal:", oldVal)
	// 比较并交换CompareAndSwap,若&a地址的值和old相等则交换
	swapped := atomic.CompareAndSwapInt64(&a, 200, 400)
	fmt.Println("swapped:", swapped)
	swapped2 := atomic.CompareAndSwapInt64(&a, 300, 400)
	fmt.Println("swapped2:", swapped2)
}

/**
sync/atomic包定义了许多原子操作的数据类型
*/
