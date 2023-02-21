package main

import "fmt"

func main() {
	// 多重for循环可以选择break指定的循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				fmt.Println(i, j)
				goto INDEX
			}
		}
	}
INDEX:
	println("end")
}
