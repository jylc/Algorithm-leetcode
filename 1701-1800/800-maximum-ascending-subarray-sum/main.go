package main

import "fmt"

// https://leetcode.cn/problems/maximum-ascending-subarray-sum/
// 1800. 最大升序子数组和

func maxAscendingSum(nums []int) int {
	sum := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			ans = max(ans, sum)
			sum = nums[i]
		}
	}
	ans = max(ans, sum)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	sum := maxAscendingSum(nums)
	fmt.Println(sum)
}
