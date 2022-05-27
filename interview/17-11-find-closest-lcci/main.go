package main

import (
	"fmt"
)

//https://leetcode.cn/problems/find-closest-lcci/
//面试题 17.11. 单词距离
func findClosest(words []string, word1 string, word2 string) int {
	index1, index2 := -1, -1
	ans := len(words)
	for i, word := range words {
		if word == word1 {
			index1 = i
		} else if word == word2 {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 {
			ans = min(ans, abs(index1, index2))
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"
	ans := findClosest(words, word1, word2)
	fmt.Println(ans)
}
