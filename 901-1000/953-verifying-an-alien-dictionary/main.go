package main

import (
	"fmt"
)

//https://leetcode.cn/problems/verifying-an-alien-dictionary/
//953. 验证外星语词典
func isAlienSorted(words []string, order string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j == len(words[i+1]) {
				return false
			}
			if diff := m[words[i][j]] - m[words[i+1][j]]; diff > 0 {
				return false
			} else if diff < 0 {
				break
			}
		}
	}

	return true
}
func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	ans := isAlienSorted(words, order)
	fmt.Println(ans)
}
