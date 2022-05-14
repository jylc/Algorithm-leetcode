package main

import "fmt"

//https://leetcode.cn/problems/one-away-lcci/
//面试题 01.05. 一次编辑
func oneEditAway(first string, second string) bool {
	firstN, secondN := len(first), len(second)
	//长度差异大于1需要多次编辑
	length := firstN - secondN
	if length > 1 || length < -1 {
		return false
	}
	a, b := 0, 0
	cnt := 1
	for a < firstN && b < secondN {
		//如果字符不同，短的字符串退回到上一个字符
		if first[a] != second[b] {
			if length == 1 {
				b--
			} else if length == -1 {
				a--
			}
			cnt--
		}
		if cnt < 0 {
			return false
		}
		a++
		b++
	}
	return true
}
func main() {
	first := "teacher"
	second := "bleacher"
	ans := oneEditAway(first, second)
	fmt.Println(ans)
}
