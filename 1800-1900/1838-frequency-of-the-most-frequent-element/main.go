package main

import (
	"fmt"
	"sort"
)

//1838. 最高频元素的频数
//https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element/
//滑动窗口+前缀和
func maxFrequency(nums []int, k int) (ans int) {
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	for r, v := range nums {
		l := sort.Search(r, func(l int) bool {
			return (r-l)*v-sum[r]+sum[l] <= k
		})
		ans = max(ans, r-l+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 3, 6}
	k := 2
	res := maxFrequency(nums, k)
	fmt.Println(res)
}
