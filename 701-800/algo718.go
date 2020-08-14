package _01_800

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	dp := make([][]int, m+1)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,n+1)
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(dp[i][j], res)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
