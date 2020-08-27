package main

import "fmt"

//392. 判断子序列
//https://leetcode-cn.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	i, j := 0, 0

	for i < lenS && j < lenT {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == lenS
}

func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	fmt.Println("result = ", result)
}
