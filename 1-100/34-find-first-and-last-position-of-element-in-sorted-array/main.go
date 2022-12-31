package main

import "fmt"

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	if n == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			left, right := mid, mid
			for i := mid; i < n; i++ {
				if nums[i] == target {
					right++
				} else {
					break
				}
			}
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					left--
				} else {
					break
				}
			}
			return []int{left + 1, right - 1}
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
func main() {
	nums := []int{2, 2}
	target := 2
	ans := searchRange(nums, target)
	fmt.Println(ans)
}
