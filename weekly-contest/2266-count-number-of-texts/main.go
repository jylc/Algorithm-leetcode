package main

//https://leetcode.cn/problems/count-number-of-texts/
//2266. 统计打字方案数

func countTexts(pressedKeys string) int {
	mod := int(1e9 + 7)
	n := len(pressedKeys)
	letters := []int{0, 0, 3, 3, 3, 3, 3, 4, 3, 4}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		c := pressedKeys[i-1]
		//i-j+1 <= letters[c-'0']：按键次数，对应不同的字母
		//pressedKeys[j-1] == c：相同的按键，比如222,33
		for j := i; j >= 1 && i-j+1 <= letters[c-'0'] && pressedKeys[j-1] == c; j-- {
			dp[i] = (dp[i] + dp[j-1]) % mod
		}
	}
	return dp[n]
}
func main() {

}
