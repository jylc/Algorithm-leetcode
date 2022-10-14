package main

import (
	"fmt"
)

// https://leetcode.cn/problems/minimum-sum-of-squared-difference/description/
// 2333. 最小差值平方和

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	tmp := make([]int, 1e5+1)
	for i := 0; i < len(nums1); i++ {
		tmp[abs(nums1[i]-nums2[i])]++
	}
	k := k1 + k2
	for i := len(tmp) - 1; i > 0 && k > 0; i-- {
		change := min(k, tmp[i])
		tmp[i-1] += change
		k -= change
		tmp[i] -= change
	}
	var ans int64
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != 0 {
			ans += int64(i * i * tmp[i])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func main() {
	nums1 := []int{10, 10, 10, 11, 5}
	nums2 := []int{1, 0, 6, 6, 1}
	k1, k2 := 11, 27
	ans := minSumSquareDiff(nums1, nums2, k1, k2)
	fmt.Println(ans)
}
