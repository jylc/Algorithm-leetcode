package main

import "sort"

//31. 下一个排列
//https://leetcode-cn.com/problems/next-permutation/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	sort.Ints(nums[i+1:])
}
func main() {

}
