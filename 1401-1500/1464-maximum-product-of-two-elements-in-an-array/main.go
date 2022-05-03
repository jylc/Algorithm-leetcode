package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/maximum-product-of-two-elements-in-an-array/
//1464. 数组中两元素的最大乘积
func maxProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return (nums[n-1] - 1) * (nums[n-2] - 1)
}
func main() {
	nums := []int{3, 4, 5, 2}
	ans := maxProduct(nums)
	fmt.Println(ans)
}
