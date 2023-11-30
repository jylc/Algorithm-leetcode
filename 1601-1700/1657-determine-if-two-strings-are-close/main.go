package main

import "slices"

// https://leetcode.cn/problems/determine-if-two-strings-are-close/description/?envType=daily-question&envId=2023-11-30
// 1657. 确定两个字符串是否接近
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	var w1 int
	var w2 int
	var cnt1 [26]int
	var cnt2 [26]int

	for _, c := range word1 {
		w1 |= 1 << (c - 'a')
		cnt1[c-'a']++
	}

	for _, c := range word2 {
		w2 |= 1 << (c - 'a')
		cnt2[c-'a']++
	}

	slices.Sort(cnt1[:])
	slices.Sort(cnt2[:])
	return w1 == w2 && slices.Equal(cnt1[:], cnt2[:])
}

func main() {

}
