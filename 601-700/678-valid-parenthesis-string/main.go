package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/valid-parenthesis-string/
//678. 有效的括号字符串
func checkValidString(s string) bool {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for i, v := range s {
		if v == '(' {
			stack1 = append(stack1, i)
		} else if v == '*' {
			stack2 = append(stack2, i)
		} else if v == ')' {
			if len(stack1) != 0 {
				stack1 = stack1[:len(stack1)-1]
			} else if len(stack2) != 0 {
				stack2 = stack2[:len(stack2)-1]
			} else {
				return false
			}
		}
	}
	i := len(stack1) - 1
	for j := len(stack2) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if stack1[i] > stack2[j] {
			return false
		}
	}
	return i < 0
}
func main() {
	s := "(*)"
	res := checkValidString(s)
	fmt.Println(res)
}
