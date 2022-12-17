package main

// https://leetcode.cn/problems/trapping-rain-water
// 42. 接雨水
func trap(height []int) int {
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		minV := min(maxLeft[i], maxRight[i])
		if minV > height[i] {
			sum += minV - height[i]
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
