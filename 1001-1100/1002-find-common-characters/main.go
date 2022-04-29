package main

import "fmt"

//https://leetcode-cn.com/problems/find-common-characters/
//1002. 查找共用字符
func min(a, b int) (ans int) {
	if a < b {
		ans = a
	} else {
		ans = b
	}
	return
}

//时间复杂度：O(n(m+|Σ|)),n为数组的长度，m为字符串的平均长度，|Σ|为字符集的长度，空间复杂度：O(|Σ|)
func commonChars(words []string) []string {
	var first map[uint8]int
	for idx, word := range words {
		cur := make(map[uint8]int)
		for i := 0; i < len(word); i++ {
			cur[word[i]]++
		}
		if idx != 0 {
			for k, v := range first {
				first[k] = min(v, cur[k])
			}
		} else {
			first = cur
		}
	}
	ans := make([]string, 0)
	for k, v := range first {
		for i := 0; i < v; i++ {
			ans = append(ans, string(k))
		}
	}
	return ans
}
func main() {
	words := []string{
		"bella", "label", "roller",
	}
	ans := commonChars(words)
	fmt.Println(ans)
}
