package main

import "fmt"

//https://leetcode.cn/problems/di-string-match/
//942. 增减字符串匹配
func diStringMatch(s string) []int {
	left, right := 0, len(s)
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			ans = append(ans, left)
			left++
		} else if s[i] == 'D' {
			ans = append(ans, right)
			right--
		}
	}
	ans = append(ans, right)
	return ans
}
func main() {
	str := "III"
	match := diStringMatch(str)
	fmt.Println(match)
}
