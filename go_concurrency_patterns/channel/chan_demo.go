package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		ch <- 10
	}()
	rev := <-ch
	fmt.Println(rev)
	// Output: 10

}
