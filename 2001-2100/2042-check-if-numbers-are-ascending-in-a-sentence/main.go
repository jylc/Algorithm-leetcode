package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/check-if-numbers-are-ascending-in-a-sentence/
// 2042. 检查句子中的数字是否递增

func areNumbersAscending(s string) bool {
	splits := strings.Split(s, " ")
	nums := make([]int, 0)
	for _, split := range splits {
		num, err := strconv.Atoi(split)
		if err == nil {
			nums = append(nums, num)
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] <= 0 {
			return false
		}
	}
	return true
}

func main() {

}
