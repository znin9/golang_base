package main

import "fmt"

// 类型断言和switch语句一起使用
func main() {
	var v interface{} = true

	// v.(type)
	switch v.(type) {
	case string:
		fmt.Println("string type")
	case int:
		fmt.Println("int type")
	case int64:
		fmt.Println("int64 type")
	case bool:
		fmt.Println("bool type")
	case interface{}:
		fmt.Println("interface{} type")
	default:
		fmt.Println("unknown type")
	}
}
