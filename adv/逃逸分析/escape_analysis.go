package main

import "fmt"

func f() *int {
	i := 1
	return &i
}

func main() {
	i := f()
	fmt.Println(*i)
}

// go build -gcfalgs '-m -l' main.go
// go tool compile -S main.go
