package main

import (
	"sort"
)

//1818. 绝对差值和
//https://leetcode-cn.com/problems/minimum-absolute-sum-difference/
//寻找最大的差值
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(v - nums1[i])
		sum += diff
		j := rec.Search(v) //如果没有找到v，返回插入v时其所在的索引位置
		if j < n {
			maxn = max(maxn, diff-(rec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-rec[j]-1))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
