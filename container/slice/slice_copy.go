package main

import "fmt"

func main() {
	sli := []int{1, 2, 3}
	sli2 := sli
	sli2[1] = 100
	fmt.Printf("sli=%v addr=%v\n", sli, &sli)
	fmt.Printf("sli2=%v addr=%v\n", sli2, &sli2)
	// sli=[1 100 3] addr=&[1 100 3]
	// sli2=[1 100 3] addr=&[1 100 3]
	sli3 := make([]int, 3, 3)
	copy(sli3, sli)
	sli3[1] = 200
	fmt.Printf("sli3=%v addr=%v\n", sli3, &sli3)
}
