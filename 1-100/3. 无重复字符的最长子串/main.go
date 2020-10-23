package main

import "strings"

//3. 无重复字符的最长子串
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring_1(s string) int { //滑动窗口
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if s == " " {
		return 1
	}
	if len(s) == 1 {
		return 1
	}
	m := make(map[uint8]bool)
	tmp, max := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; !ok {
				m[s[j]] = true
				tmp++
			} else {
				if tmp > max {
					max = tmp
				}
				m = make(map[uint8]bool)
				tmp = 0
				break
			}
		}
	}
	return max
}
func main() {

}
