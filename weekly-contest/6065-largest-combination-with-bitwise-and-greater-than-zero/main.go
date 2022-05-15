package main

import (
	"fmt"
)

//https://leetcode.cn/contest/weekly-contest-293/problems/largest-combination-with-bitwise-and-greater-than-zero/
//6065. 按位与结果大于零的最长组合
func largestCombination(candidates []int) int {
	max := 0
	for i := 0; i < 30; i++ {
		tmp := 0
		for _, candidate := range candidates {
			tmp += candidate >> i & 1
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
func main() {
	candidates := []int{16, 17, 71, 62, 12, 24, 14}
	ans := largestCombination(candidates)
	fmt.Println(ans)
}
