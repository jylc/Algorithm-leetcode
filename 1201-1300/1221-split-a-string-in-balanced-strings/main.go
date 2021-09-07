package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
//1221. 分割平衡字符串
func balancedStringSplit(s string) int {
	stack := make([]rune, 0, len(s))
	cnt := 0
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			cnt++
		} else {
			top := stack[len(stack)-1]
			if top == c {
				stack = append(stack, c)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return cnt
}
func main() {
	s := "RL"
	res := balancedStringSplit(s)
	fmt.Println(res)
}
