package main

import "fmt"

//https://leetcode.cn/problems/generate-parentheses/
//22. 括号生成
func generateParenthesis(n int) (ans []string) {
	var dfs func(arr []byte, l, r int)
	dfs = func(arr []byte, l, r int) {
		//减枝;‘(’的数量要大于')'
		if l < r || l > n || r > n {
			return
		}
		if l == n && r == n {
			ans = append(ans, string(arr))
		}
		dfs(append(arr, '('), l+1, r)
		dfs(append(arr, ')'), l, r+1)
	}
	arr := make([]byte, 0)
	dfs(arr, 0, 0)
	return
}

func main() {
	ans := generateParenthesis(3)
	fmt.Println(ans)
}
