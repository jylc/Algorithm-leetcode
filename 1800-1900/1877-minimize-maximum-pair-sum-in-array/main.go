package main

import (
	"math"
	"sort"
)

//1877. 数组中最大数对和的最小值
//https://leetcode-cn.com/problems/minimize-maximum-pair-sum-in-array/

func minPairSum(nums []int) int {
	sort.Ints(nums)
	cnt := 0
	maxNum := math.MinInt32

	for i, j := 0, len(nums)-1; i < j; {
		maxNum = max(maxNum, nums[i]+nums[j])
		i++
		j--
		cnt++
	}
	return maxNum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
