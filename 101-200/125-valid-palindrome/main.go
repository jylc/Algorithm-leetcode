package main

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/valid-palindrome/
//125. 验证回文串

func isPalindrome(s string) bool {
	tmp := strings.ToLower(s)
	arr := make([]int32, 0)
	for _, word := range tmp {
		if (word >= 'a' && word <= 'z') || (word >= '0' && word <= '9') {
			arr = append(arr, word)
		}
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func main() {
	s := "   "
	ans := isPalindrome(s)
	fmt.Println(ans)
}
