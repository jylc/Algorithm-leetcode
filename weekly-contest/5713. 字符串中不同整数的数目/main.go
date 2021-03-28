package main

import (
	"fmt"
	"strings"
)

//5713. 字符串中不同整数的数目
//https://leetcode-cn.com/contest/weekly-contest-234/problems/number-of-different-integers-in-a-string/
func numDifferentIntegers(word string) int {
	n := len(word)
	var flag bool
	count := 0

	arr := make([]uint8, 0)
	for i := 0; i < n; i++ {
		if word[i] < '0' || word[i] > '9' {
			flag = false
			continue
		} else if flag == false && (word[i] >= '0' && word[i] <= '9') {
			flag = true
			arr = append(arr, ' ')
		}
		arr = append(arr, word[i])
	}
	s := string(arr)
	str := strings.Split(s, " ")
	for i := 0; i < len(str); i++ {
		if len(str[i]) != 1 {
			index := 0
			for j := 0; j < len(str[i]); j++ {
				if str[i][j] != '0' {
					index = j
					break
				}
			}
			str[i] = str[i][index:]
		}
	}

	m := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		if _, v := m[str[i]]; !v {
			m[str[i]] = true
		}
	}
	for i := 0; i < len(m); i++ {
		if _, v := m[str[i]]; str[i] != "" && v {
			count++
		}
	}

	return count
}
func main() {
	fmt.Print("res = ", numDifferentIntegers("0a0"))
}
