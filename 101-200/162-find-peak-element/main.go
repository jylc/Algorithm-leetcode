package main

import (
	"math"
)

//https://leetcode-cn.com/problems/find-peak-element/
//162. 寻找峰值
//
func findPeakElement(nums []int) int {
	maxVal := math.MinInt32
	index := 0
	for i, num := range nums {
		if maxVal < num {
			maxVal = num
			index = i
		}
	}
	return index
}

func findPeakElement1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		m := left + (right-left)/2
		if nums[m] < nums[m+1] {
			left = m + 1
		} else {
			right = m
		}
	}
	return left
}
func main() {

}
