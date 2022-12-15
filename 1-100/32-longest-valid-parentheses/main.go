package main

import "fmt"

// https://leetcode.cn/problems/longest-valid-parentheses/
// 32. 最长有效括号

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	dp[0] = 0
	maxValue := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i] + dp[i-2]
				}
			} else if dp[i-1] > 0 { // s[i]=='('
				if (i-dp[i-1]-1) >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if (i - dp[i-1] - 2) >= 0 {
						dp[i] = dp[i] + dp[i-dp[i-1]-2]
					}
				}
			}
		}
		maxValue = max(maxValue, dp[i])
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := ")()())"
	ans := longestValidParentheses(s)
	fmt.Println(ans)
}
