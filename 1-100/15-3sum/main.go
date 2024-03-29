package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/3sum/
//15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{1, 1}
	ans := threeSum(nums)
	fmt.Println(ans)
}
