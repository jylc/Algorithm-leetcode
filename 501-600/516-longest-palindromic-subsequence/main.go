package main

//https://leetcode-cn.com/problems/longest-palindromic-subsequence/
//516. 最长回文子序列

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else if s[i] != s[j] {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
