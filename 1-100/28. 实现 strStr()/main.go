package main

import "fmt"

//28. 实现 strStr()
//https://leetcode-cn.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	if needle == "" ||len(haystack)<len(needle){
		return 0
	}
	lenA, lenB := len(haystack), len(needle)
	j := 0
	for k := 0; k < lenA; k++ {
		i := k
		for i < lenA && j < lenB {
			if haystack[i] == needle[j] {
				i++
				j++
			} else {
				j = 0
				break
			}
		}
		if j == lenB {
			return k
		}
	}
	return -1
}
func main() {
	result := strStr("aaaaaaa", "bba")
	fmt.Println("result = ",result)
}
