package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/maximum-subarray/
//53. 最大子数组和

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	arr := []int{-1, 1}
	ans := maxSubArray(arr)
	fmt.Println(ans)
}
