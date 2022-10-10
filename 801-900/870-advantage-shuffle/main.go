package main

import "sort"

// https://leetcode.cn/problems/advantage-shuffle/
// 870. 优势洗牌

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 0; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}
	sort.Slice(idx1, func(i, j int) bool {
		return nums1[idx1[i]] < nums1[idx1[j]]
	})
	sort.Slice(idx2, func(i, j int) bool {
		return nums2[idx2[i]] < nums2[idx2[j]]
	})

	ans := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[idx1[i]] > nums2[idx2[left]] {
			ans[idx2[left]] = nums1[idx1[i]]
			left++
		} else {
			// 在遍历结束后,nums2实际上是原先nums2的一个后缀，因此用左右指针
			ans[idx2[right]] = nums1[idx1[i]]
			right--
		}
	}
	return ans
}
func main() {

}
