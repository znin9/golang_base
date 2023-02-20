package main

import (
	"fmt"
	"strconv"
)

// strconv.Parse系列
// strconv.ParseInt(s string,base int,bitSize int) (int64,error)
func main() {
	i, err := strconv.ParseInt("+100", 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("i=%d\n", i)

	ui, err := strconv.ParseUint("100", 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ui=%d\n", ui)

	b, err := strconv.ParseBool("TRUE")
	if err != nil {
		panic(err)
	}
	fmt.Printf("b=%v\n", b)

	f, err := strconv.ParseFloat("1111.22222222", 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("f=%v\n", f)
}
