package main

import "fmt"

/*
*
笔记：
*/

func f() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func f1() {
	s := 1
	defer fmt.Println(s)
	s = 10
}

func f2() {
	obj := struct {
		Name string
	}{Name: "zs"}

	defer func() {
		fmt.Println(obj)
		fmt.Println(obj.Name)
	}()

	obj.Name = "lisi"
}

func main() {
	f() // 延迟语句逆序执行
	// Output:
	// 3
	// 2
	// 1
	f1() // 延迟语句如果使用的参数是一个值，则和最开始定义的一样。
	// Output:
	// 1
	f2() // 延迟语句如果使用的是“引用”，那么defer函数中的可能和最初的值不一致。
	// Output:
	// {lisi}
	// lisi

}
