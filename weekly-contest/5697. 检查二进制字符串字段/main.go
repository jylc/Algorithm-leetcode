package main

import "fmt"

//5697. 检查二进制字符串字段
//https://leetcode-cn.com/contest/weekly-contest-231/problems/check-if-binary-string-has-at-most-one-segment-of-ones/

func checkOnesSegment(s string) bool {
	n := len(s)
	flag := false
	for i := 0; i < n; i++ {
		if flag && s[i] == '1' {
			return false
		}
		if s[i] == '0' {
			flag = true
		}
	}
	return true
}

func main() {
	s := "1110000"
	segment := checkOnesSegment(s)
	fmt.Print("res = ", segment)
}
