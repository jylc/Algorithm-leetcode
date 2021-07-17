package main

import (
	"fmt"
)

//剑指 Offer 42. 连续子数组的最大和
//https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
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
	nums := []int{-2, -1}
	res := maxSubArray(nums)
	fmt.Println(res)
}
