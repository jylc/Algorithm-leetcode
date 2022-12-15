package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/regular-expression-matching/
// 10. 正则表达式匹配

func isMatch(s string, p string) bool {
	var dp func(s string, i int, p string, j int) bool
	var memo = make(map[string]bool)
	dp = func(s string, i int, p string, j int) bool {
		m, n := len(s), len(p)
		if j == n {
			return i == m
		}

		if i == m {
			if ((n - j) % 2) == 1 {
				return false
			}
			for ; j+1 < n; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}

		key := strconv.Itoa(i) + "," + strconv.Itoa(j)
		if _, ok := memo[key]; ok {
			return memo[key]
		}
		res := false
		if s[i] == p[j] || p[j] == '.' {
			if j+1 < n && p[j+1] == '*' {
				res = dp(s, i, p, j+2) || dp(s, i+1, p, j)
			} else {
				res = dp(s, i+1, p, j+1)
			}
		} else {
			if j+1 < n && p[j+1] == '*' {
				res = dp(s, i, p, j+2)
			} else {
				res = false
			}
		}
		memo[key] = res
		return res
	}
	return dp(s, 0, p, 0)
}

func main() {
	s := "aa"
	p := ".*"
	ans := isMatch(s, p)
	fmt.Println(ans)
}
