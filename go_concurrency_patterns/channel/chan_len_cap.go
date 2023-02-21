package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Println(len(ch), cap(ch))
	// Output: 0 0
	go func() {
		ch <- 10
	}()
	fmt.Println(<-ch)
	// 10

	ch2 := make(chan int, 1)
	fmt.Println(len(ch2), cap(ch2))
	// Output: 0 1
}
