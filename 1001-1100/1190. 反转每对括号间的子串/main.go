package main

import "fmt"

//1190. 反转每对括号间的子串
//https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/
func reverseParentheses(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]int, 0)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			front := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			end := i
			s = reverse(s, front, end)
		}
	}
	result := make([]rune, 0)
	for _, c := range s {
		if c != '(' && c != ')' {
			result = append(result, c)
		}
	}
	return string(result)
}

func reverse(s string, i, j int) string {
	r := []rune(s)
	for i, j = i+1, j-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	result := reverseParentheses("(abcd)")
	fmt.Print(result)
}
