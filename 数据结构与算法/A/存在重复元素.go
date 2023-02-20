package main

import "fmt"

/*
*
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
*/
func main() {
	nums := []int{1, 2, 3, 1}
	duplicate := containsDuplicate(nums)
	fmt.Println(duplicate)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] != false {
			return true
		}
		m[v] = true
	}
	return false
}
