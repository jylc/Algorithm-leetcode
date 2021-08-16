package main

//https://leetcode-cn.com/problems/longest-palindromic-substring/
//5. 最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		for k := 0; k <= i; k++ {
			dp[i][k] = s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1])
			if dp[i][k] && i-k+1 > len(ans) {
				ans = s[k : i+1]
			}
		}
	}
	return ans
}

func main() {
}
