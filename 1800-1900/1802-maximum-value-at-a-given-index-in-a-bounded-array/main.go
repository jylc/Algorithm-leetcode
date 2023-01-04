package main

import "sort"

// https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/
// 1802. 有界数组中指定下标处的最大值

func maxValue(n int, index int, maxSum int) int {
	return sort.Search(maxSum, func(mid int) bool {
		left := index
		right := n - index - 1
		return mid+f(mid, left)+f(mid, right) >= maxSum
	})

}
func f(big, length int) int {
	if length == 0 {
		return 0
	}
	if length <= big {
		return (2*big + 1 - length) * length / 2
	}
	return (big+1)*big/2 + length - big
}

func main() {

}
