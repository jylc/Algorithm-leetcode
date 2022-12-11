package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/sum-of-beauty-of-all-substrings/
// 1781. 所有子字符串美丽值之和
func beautySum(s string) int {
	ans := 0
	for left := 0; left < len(s)-1; left++ {
		cMap := make(map[int]int)
		for right := left; right < len(s); right++ {
			cMap[int(s[right])]++
			ans += beauty(cMap)
		}
	}
	return ans
}

func beauty(m map[int]int) int {
	max := math.MinInt
	min := math.MaxInt
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func main() {
	ans := beautySum("aabcbaa")
	fmt.Println(ans)
}
