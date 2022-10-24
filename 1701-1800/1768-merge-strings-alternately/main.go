package main

// https://leetcode.cn/problems/merge-strings-alternately/
// 1768. 交替合并字符串

func mergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	ans := make([]byte, 0)
	minLen := min(len1, len2)
	for i := 0; i < minLen; i++ {
		ans = append(ans, word1[i], word2[i])
	}
	if len1 > minLen {
		ans = append(ans, word1[minLen:]...)
	} else if len2 > minLen {
		ans = append(ans, word2[minLen:]...)
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
