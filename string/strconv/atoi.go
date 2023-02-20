package main

import (
	"fmt"
	"strconv"
)

// strconv.Atoi(s string)(int,error)
// strconv.Itoa(a int) string
func main() {
	// string -> int
	i, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("i=%d\n", i)

	// int -> string
	s := strconv.Itoa(100)
	fmt.Printf("s=%s\n", s)
}
