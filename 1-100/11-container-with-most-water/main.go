package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/container-with-most-water/
//11. 盛最多水的容器
func min(a, b int) (ans int) {
	if a < b {
		ans = a
	} else {
		ans = b
	}
	return
}
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left != right {
		h := min(height[left], height[right])
		w := right - left
		if ans < h*w {
			ans = h * w
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//height := []int{1, 1}
	ans := maxArea(height)
	fmt.Println(ans)
}
