package main

import "strings"

// https://leetcode.cn/problems/string-rotation-lcci/
// 面试题 01.09. 字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	copied := strings.Repeat(s2, 2)
	s3 := s2 + copied
	res := strings.Contains(s3, s1)
	return res
}
func main() {

}
