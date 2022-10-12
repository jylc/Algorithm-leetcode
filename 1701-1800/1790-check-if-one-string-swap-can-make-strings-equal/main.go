package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
// 1790. 仅执行一次字符串交换能否使两个字符串相等

func areAlmostEqual(s1 string, s2 string) bool {
	isSame := strings.Compare(s1, s2)
	if isSame == 0 {
		return true
	}
	count := 0
	a, b := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
		if count > 2 {
			return false
		}
	}
	tmp1 := []rune(s1)
	tmp1[a], tmp1[b] = tmp1[b], tmp1[a]
	tmp2 := string(tmp1)
	fmt.Println(tmp2)
	return strings.Compare(tmp2, s2) == 0
}

func main() {

}
