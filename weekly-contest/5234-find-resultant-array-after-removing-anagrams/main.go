package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/contest/weekly-contest-293/problems/find-resultant-array-after-removing-anagrams/
//5234. 移除字母异位词后的结果数组

func removeAnagrams(words []string) []string {
	tmp_words := make([]string, len(words))
	for i, word := range words {
		tmp := []rune(word)
		sort.Slice(tmp, func(i, j int) bool {
			if tmp[i] <= tmp[j] {
				return true
			} else {
				return false
			}
		})
		tmp_words[i] = string(tmp)
	}
	ans := make([]string, 0)
	left, right := 0, 0
	for right < len(tmp_words) {
		ans = append(ans, words[left])
		for right < len(tmp_words) && tmp_words[left] == tmp_words[right] {
			right++
		}
		left = right
	}
	return ans
}
func main() {
	words := []string{"a", "b", "a"}
	ans := removeAnagrams(words)
	fmt.Println(ans)
}
