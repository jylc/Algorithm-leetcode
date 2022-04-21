package main

import "fmt"

//https://leetcode-cn.com/problems/longest-consecutive-sequence/
//128. 最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int, 0)
	max := 0
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			left, right := m[num-1], m[num+1]
			m[num] = left + right + 1
			m[num-left] = m[num]
			m[num+right] = m[num]
			if m[num] > max {
				max = m[num]
			}
		}
	}
	return max
}
func main() {
	nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	ans := longestConsecutive(nums)
	fmt.Println(ans)
}
