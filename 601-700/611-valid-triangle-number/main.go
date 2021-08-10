package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/valid-triangle-number/
//611. 有效三角形的个数
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var res int
	for k := len(nums) - 1; k >= 2; k-- {
		l, r := 0, k-1
		for l < r {
			if nums[l]+nums[r] > nums[k] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}

func main() {

}
