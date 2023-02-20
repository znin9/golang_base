package main

/*
*
求斐波那契数列Fibonacci sequence

问题：f(0)=0 f(2)=1 f(n)=f(n-1)+f(n-2) 求f(n)

0 1 1 2 3 5 8 13 21 34
*/
func main() {
	recursion(9)
}

// recursion 完全递归求解斐波那契数列
func recursion(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return recursion(n-1) + recursion(n-2)
}
