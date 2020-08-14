package _01_300

import (
	"math"
	"sort"
)

func MinSubArrayLen(s int, nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	if n == 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans==math.MaxInt32{
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
