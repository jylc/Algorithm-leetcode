package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//https://leetcode-cn.com/problems/reorder-data-in-log-files/
//937. 重新排列日志文件
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		str1, str2 := logs[i], logs[j]
		s1, s2 := strings.SplitN(str1, " ", 2)[1], strings.SplitN(str2, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit2 && isDigit1 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && str1 < str2
		}
		return !isDigit1
	})
	return logs
}
func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	strs := reorderLogFiles(logs)
	fmt.Println(strs)
}
