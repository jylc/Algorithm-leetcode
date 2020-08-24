package main

import "strings"

//459. 重复的子字符串
//https://leetcode-cn.com/problems/repeated-substring-pattern/

func repeatedSubstringPattern(s string) bool {
	s1 := s + s
	s2 := s[1 : len(s1)-1]
	if strings.Contains(s2, s) {
		return true
	} else {
		return false
	}
}

func main() {

}
