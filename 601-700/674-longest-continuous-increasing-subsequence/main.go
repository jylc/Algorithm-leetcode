package main

import "fmt"

//https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
//674. 最长连续递增序列
func findLengthOfLCIS(nums []int) int {
	l, r := 0, 1
	max := 0
	if len(nums) == 1 {
		return 1
	}
	for r < len(nums) {
		for r < len(nums) && nums[r-1] < nums[r] {
			r++
		}
		if max < r-l {
			max = r - l
		}
		l = r
		r++
	}
	return max
}
func main() {
	//nums := []int{1, 3, 5, 4, 7}
	nums := []int{2, 2, 2, 2, 2}
	ans := findLengthOfLCIS(nums)
	fmt.Println(ans)
}
