package main

// https://leetcode.cn/problems/score-of-parentheses/
// 856. 括号的分数
// 时间复杂度：O(n)；空间复杂度：O(n)
func scoreOfParentheses(s string) int {
	stack := []int{0}
	for _, c := range s {
		if c == '(' { // 计算左括号内部子平衡括号字符串 A 的分数，是‘子平衡括号字符串’的分数
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*v, 1)
		}
	}
	return stack[0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
