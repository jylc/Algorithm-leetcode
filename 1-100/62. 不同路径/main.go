package main

import (
	"fmt"
	"math/big"
)

//62. 不同路径
//https://leetcode-cn.com/problems/unique-paths/
//动态规划
func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 0
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//组合数学
func uniquePaths2(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func main() {
	fmt.Print("result = ", uniquePaths(3, 2))
}
