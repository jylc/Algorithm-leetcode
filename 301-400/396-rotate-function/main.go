package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/rotate-function/
//396. 旋转函数
func maxRotateFunction(nums []int) int {
	numSum := 0
	f := 0
	for i, num := range nums {
		numSum += num
		f += i * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		if ans < f {
			ans = f
		}
	}
	return ans
}
func main() {
	nums := []int{4, 3, 2, 6, 7, 8}
	ans := maxRotateFunction(nums)
	fmt.Println(ans)
}
