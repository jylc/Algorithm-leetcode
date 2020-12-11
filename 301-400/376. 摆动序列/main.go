package main

//376. 摆动序列
//https://leetcode-cn.com/problems/wiggle-subsequence/
//贪心算法
func wiggleMaxLength(nums []int) int {
	preDiff, curDiff := 0, 0
	cnt := 0
	for i, num := range nums {
		if i == len(nums)-1 {
			cnt++
			break
		}
		curDiff = nums[i+1] - num
		if preDiff <= 0 && curDiff > 0 || preDiff >= 0 && curDiff < 0 {
			cnt++
			preDiff = curDiff
		}
	}
	return cnt
}

//动态规划
func wiggleMaxLength1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}(down, up)
}

func main() {

}
