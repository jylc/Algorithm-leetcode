package main

import (
	"fmt"
	"strings"
)

//459. 重复的子字符串
//https://leetcode-cn.com/problems/repeated-substring-pattern/
//字符串匹配：
func repeatedSubstringPattern(s string) bool {
	s1 := s + s
	s2 := s1[1 : len(s1)-1]
	if strings.Contains(s2, s) {
		return true
	} else {
		return false
	}
}

//枚举：时间复杂度O(n^2)，空间复杂度O(1)
func repeatedSubstringPattern1(s string) bool {
	n := len(s)
	//外层设置子串长度范围
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			//子串遍历
			for j := i; j < n; j++ {
				if s[j] == s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "abcabcabcabc"
	ans := repeatedSubstringPattern1(s)
	fmt.Println(ans)
}
