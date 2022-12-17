package main

import (
	"fmt"
)

// https://leetcode.cn/problems/jump-game-ii/
// 45. 跳跃游戏 II
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	i := 0
	for j := 1; j < len(nums); j++ {
		for i+nums[i] < j {
			i++
		}
		dp[j] = dp[i] + 1
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	ans := jump(nums)
	fmt.Println(ans)
}
