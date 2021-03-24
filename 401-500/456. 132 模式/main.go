package main

import "math"

//456. 132 模式
//https://leetcode-cn.com/problems/132-pattern/
func find132pattern(nums []int) bool {
	temp := make([]int, 0)
	right := math.MinInt64
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < right {
			return true
		}

		//单调栈；从右到左遍历，寻找i+1到len(nums)-1中最大值，此值比nums[i]小
		for len(temp) > 0 && nums[i] > temp[len(temp)-1] {
			//若栈不为空且nums[i]大于栈顶元素则不断弹出元素；栈底元素始终是i+1到n-1的最大元素，而right是次大元素
			right = temp[len(temp)-1]
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, nums[i])
	}
	return false
}
func main() {

}
