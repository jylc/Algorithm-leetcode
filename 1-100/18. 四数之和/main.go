package main

import "sort"

//18. 四数之和
//https://leetcode-cn.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var equals [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				distance := sum - target
				if distance == 0 {
					equals = append(equals, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if distance > 0 {
					for cur := l; l > k && nums[l] == nums[cur]; l-- {
					}
				} else {
					for cur := k; k < l && nums[k] == nums[cur]; k++ {
					}
				}
			}
		}
	}
	return equals
}
func main() {

}
