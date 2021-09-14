package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/
//524. 通过删除字母匹配到字典里最长单词

func findLongestWord(s string, dictionary []string) (ans string) {
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					if len(t) > len(ans) || len(t) == len(ans) && t < ans {
						ans = t
					}
					break
				}
			}
		}
	}
	return
}

func main() {
	str := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}

	ans := findLongestWord(str, dictionary)
	fmt.Println("ans:", ans)
}
