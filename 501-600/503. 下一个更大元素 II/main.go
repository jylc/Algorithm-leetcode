package main

import "fmt"

//503. 下一个更大元素 II
//https://leetcode-cn.com/problems/next-greater-element-ii/
//遍历:时间O(n^2)，空间O(n)
func nextGreaterElements(nums []int) []int {
	var res []int
	n := len(nums)

	for i := 0; i < n; i++ {
		t := nums[i]
		flag := false
		for j := 1; j < n; j++ {
			k := (i + j) % n
			if nums[k] > t {
				res = append(res, nums[k])
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, -1)
		}
	}
	return res
}

//单调栈:时间O(n)，空间O(n)
func nextGreaterElements1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}
func main() {
	nums := []int{1, 2, 1}
	res := nextGreaterElements1(nums)
	fmt.Print("res = ", res)
}
