package main

// https://leetcode.cn/problems/minimum-moves-to-convert-string/
// 2027. 转换字符串的最少操作次数

func minimumMoves(s string) (ans int) {
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'X' {
			ans++
			i += 2
		}
	}
	return
}

func main() {

}
