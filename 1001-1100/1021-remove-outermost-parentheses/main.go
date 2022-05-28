package main

import "fmt"

//https://leetcode.cn/problems/remove-outermost-parentheses/
//1021. 删除最外层的括号
func removeOuterParentheses(s string) string {
	right := 0
	var ans string
	need := 0
	tmp := make([]uint8, 0)
	for right < len(s) {
		c := s[right]
		if c == '(' {
			need++
		} else if c == ')' {
			need--
		}
		tmp = append(tmp, c)
		if need == 0 {
			ans = ans + string(tmp)[1:len(tmp)-1]
			tmp = tmp[:0]
		}
		right++
	}
	return ans
}
func main() {
	s := "(()())(())(()(()))"
	ans := removeOuterParentheses(s)
	fmt.Println(ans)
}
