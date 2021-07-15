package main

import (
	"sort"
)

//1846. 减小和重新排列数组后的最大元素
//https://leetcode-cn.com/problems/maximum-element-after-decreasing-and-rearranging/
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	if arr[0] != 1 {
		arr[0] = 1
	}
	for i := 1; i < len(arr); i++ {
		if abs(arr[i]-arr[i-1]) > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {

}
