package main

// https://leetcode.cn/problems/distinct-subsequences-ii/
// 940. 不同的子序列 II

// 动态规划
func distinctSubseqII(s string) int {
	n := len(s)
	f := make([][26]int, n)
	f[0][s[0]-'a'] = 1
	for i := 1; i < n; i++ {
		total := int64(0)
		for _, v := range f[i-1] {
			total += int64(v)
		}
		f[i] = f[i-1]
		f[i][s[i]-'a'] = 1 + int(total%(1e9+7))
	}
	total := int64(0)
	for _, v := range f[n-1] {
		total += int64(v)
	}
	return int(total % (1e9 + 7))
}
func main() {

}
