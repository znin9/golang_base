package main

import (
	"fmt"
	"reflect"
)

// 类型断言，一般用于顶层接口向下转型
// 语法格式: value,ok := x.(T)
func main() {
	// interface{} 向下转型
	var v interface{} = 100
	i, b := v.(int)
	if !b {
		fmt.Println("类型断言失败")
	} else {
		fmt.Printf("i=%d\n", i)
	}

	// Hello类型向下转型
	var p Hello = Person{}
	p2, b := p.(Person)
	fmt.Printf("type=%v value=%v b=%v\n", reflect.TypeOf(p2), reflect.ValueOf(p2), b)
	// Output: type=main.Person value={} b=true
	p3, b := p.(interface{})
	fmt.Printf("type=%v value=%v b=%v\n", reflect.TypeOf(p3), reflect.ValueOf(p3), b)
	// Output: type=main.Person value={} b=true
}

type Hello interface {
	SayHello(name string)
}

type Person struct {
}

func (receiver Person) SayHello(name string) {
	fmt.Println("hello:", name)
}
