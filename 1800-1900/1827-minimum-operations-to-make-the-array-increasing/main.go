package main

import "fmt"

// https://leetcode.cn/problems/minimum-operations-to-make-the-array-increasing/
// 1827. 最少操作使数组递增

func minOperations(nums []int) int {
	ans := 0
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			ans += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return ans
}
func main() {
	nums := []int{1, 5, 2, 4, 1}
	ans := minOperations(nums)
	fmt.Println(ans)
}
