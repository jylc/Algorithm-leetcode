package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
//462. 最少移动次数使数组元素相等 II
func minMoves2(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	res := 0
	for i < j {
		res += nums[j] - nums[i]
		j--
		i++
	}
	return res
}
func main() {
	nums := []int{1, 0, 0, 8, 6}
	ans := minMoves2(nums)
	fmt.Println(ans)
}
