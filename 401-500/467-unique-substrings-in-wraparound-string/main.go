package main

import "fmt"

//https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/
//467. 环绕字符串中唯一的子字符串
func findSubstringInWraproundString(p string) int {
	chCount := 1
	count := 0
	m := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		if i > 0 && ((p[i] - 'a') == (p[i-1]+1-'a')%26) {
			chCount++
		} else {
			chCount = 1
		}
		if m[p[i]-'a'] < chCount {
			m[p[i]-'a'] = chCount
		}
	}
	for _, v := range m {
		count += v
	}
	return count
}
func main() {
	p := "mnoabzabccb"
	ans := findSubstringInWraproundString(p)
	fmt.Println(ans)
}
