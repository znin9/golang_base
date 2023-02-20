package main

import "fmt"

func main() {
	var sli []int
	for i := 0; i < 10; i++ {
		sli = append(sli, i)
	}
	fmt.Println(sli, len(sli), cap(sli))

	// 准备拷贝最后五个元素到新的slice
	dest := make([]int, 5)
	fmt.Println(dest, len(dest), cap(dest))
	copyNum := copy(dest, sli[5:])
	fmt.Printf("copyNum=%d dest=%v len(dest)=%d cap(dest)=%d\n", copyNum, dest, len(dest), cap(dest))
	// Output:
	// [0 1 2 3 4 5 6 7 8 9] 10 16
	// [0 0 0 0 0] 5 5
	// copyNum=5 dest=[5 6 7 8 9] len(dest)=5 cap(dest)=5
}
