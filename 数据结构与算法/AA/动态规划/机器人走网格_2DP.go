package main

import "fmt"

/*
*
二位数组的DP

问题描述: 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	问总共有多少条不同的路径？

解题思路:

	1.定义数组的含义 例如在5x5的格子中，从A到B,那么数组的下标dp[i][j]就表示到达i,j这个位置一共的路径种数。
					A x x x x
					x x x x x
					x x x x x
					x x x x x
					x x x x B
	2.分析数组元素之间的关系。因为机器人只能向下或者向右移动，因此dp[i][j] = dp[i-1][j] + dp[i][j-1]
	3.确定
*/
func main() {
	sli := make([][]int, 7)
	fmt.Println(sli)
}
